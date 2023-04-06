package breaker

import (
	"fmt"
	"ginessential/pkg/mathx"
	"ginessential/pkg/proc"
	"ginessential/pkg/stat"
	"ginessential/pkg/stringx"
	"ginessential/pkg/timex"
	"strings"
	"sync"
)

const (
	numHistoryReasons = 5
	timeFormat        = "15:04:05"
)

type (
	// A Breaker represents a circuit breaker.
	Breaker interface {
		// Name returns the name of the Breaker.
		Name() string

		// Allow checks if the request is allowed.
		// If allowed, a promise will be returned, the caller needs to call promise.Accept()
		// on success, or call promise.Reject() on failure.
		// If not allow, ErrServiceUnavailable will be returned.
		Allow() (Promise, error)

		// Do runs the given request if the Breaker accepts it.
		// Do returns an error instantly if the Breaker rejects the request.
		// If a panic occurs in the request, the Breaker handles it as an error
		// and causes the same panic again.
		Do(req func() error) error

		// DoWithAcceptable runs the given request if the Breaker accepts it.
		// DoWithAcceptable returns an error instantly if the Breaker rejects the request.
		// If a panic occurs in the request, the Breaker handles it as an error
		// and causes the same panic again.
		// acceptable checks if it's a successful call, even if the err is not nil.
		DoWithAcceptable(req func() error, acceptable Acceptable) error

		// DoWithFallback runs the given request if the Breaker accepts it.
		// DoWithFallback runs the fallback if the Breaker rejects the request.
		// If a panic occurs in the request, the Breaker handles it as an error
		// and causes the same panic again.
		DoWithFallback(req func() error, fallback func(err error) error) error

		// DoWithFallbackAcceptable runs the given request if the Breaker accepts it.
		// DoWithFallbackAcceptable runs the fallback if the Breaker rejects the request.
		// If a panic occurs in the request, the Breaker handles it as an error
		// and causes the same panic again.
		// acceptable checks if it's a successful call, even if the err is not nil.
		DoWithFallbackAcceptable(req func() error, fallback func(err error) error, acceptable Acceptable) error
	}

	// Option defines the method to customize a Breaker.
	Option func(breaker *circuitBreaker)

	// Promise interface defines the callbacks that returned by Breaker.Allow.
	Promise interface {
		// Accept tells the Breaker that the call is successful.
		Accept()
		// Reject tells the Breaker that the call is failed.
		Reject(reason string)
	}

	internalPromise interface {
		Accept()
		Reject()
	}

	circuitBreaker struct {
		name string
		throttle
	}

	internalThrottle interface {
		allow() (internalPromise, error)
		doReq(req func() error, fallback func(err error) error, acceptable Acceptable) error
	}

	throttle interface {
		allow() (Promise, error)
		doReq(req func() error, fallback func(err error) error, acceptable Acceptable) error
	}
)

// NewBreaker returns a Breaker object.
// opts can be used to customize the Breaker.
func NewBreaker(opts ...Option) Breaker {
	var b circuitBreaker
	for _, opt := range opts {
		opt(&b)
	}
	if len(b.name) == 0 {
		b.name = stringx.Rand()
	}
	b.throttle = newLoggedThrottle(b.name, newGoogleBreaker())
	return &b
}

// Allow checks if the request is allowed.
// If allowed, a promise will be returned, the caller needs to call promise.Accept()
// on success, or call promise.Reject() on failure.
// If not allow, ErrServiceUnavailable will be returned.
func (cb *circuitBreaker) Allow() (Promise, error) {
	return cb.throttle.allow()
}

// Do runs the given request if the Breaker accepts it.
// Do returns an error instantly if the Breaker rejects the request.
// If a panic occurs in the request, the Breaker handles it as an error
// and causes the same panic again.
func (cb *circuitBreaker) Do(req func() error) error {
	return cb.throttle.doReq(req, nil, defaultAcceptable)
}

// DoWithAcceptable runs the given request if the Breaker accepts it.
// DoWithAcceptable returns an error instantly if the Breaker rejects the request.
// If a panic occurs in the request, the Breaker handles it as an error
// and causes the same panic again.
// acceptable checks if it's a successful call, even if the err is not nil.
func (cb *circuitBreaker) DoWithAcceptable(req func() error, acceptable Acceptable) error {
	return cb.throttle.doReq(req, nil, acceptable)
}

// DoWithFallback runs the given request if the Breaker accepts it.
// DoWithFallback runs the fallback if the Breaker rejects the request.
// If a panic occurs in the request, the Breaker handles it as an error
// and causes the same panic again.
func (cb *circuitBreaker) DoWithFallback(req func() error, fallback func(err error) error) error {
	return cb.throttle.doReq(req, fallback, defaultAcceptable)
}

// DoWithFallbackAcceptable runs the given request if the Breaker accepts it.
// DoWithFallbackAcceptable runs the fallback if the Breaker rejects the request.
// If a panic occurs in the request, the Breaker handles it as an error
// and causes the same panic again.
// acceptable checks if it's a successful call, even if the err is not nil.
func (cb *circuitBreaker) DoWithFallbackAcceptable(req func() error,
	fallback func(err error) error, acceptable Acceptable) error {
	return cb.throttle.doReq(req, fallback, acceptable)
}

// Name returns the name of the Breaker.
func (cb *circuitBreaker) Name() string {
	return cb.name
}

// WithName returns a function to set the name of a Breaker.
func WithName(name string) Option {
	return func(b *circuitBreaker) {
		b.name = name
	}
}

func defaultAcceptable(err error) bool {
	return err == nil
}

type loggedThrottle struct {
	name string
	internalThrottle
	errWin *errorWindow
}

func newLoggedThrottle(name string, t internalThrottle) loggedThrottle {
	return loggedThrottle{
		name:             name,
		internalThrottle: t,
		errWin:           new(errorWindow),
	}
}

func (lt loggedThrottle) allow() (Promise, error) {
	promise, err := lt.internalThrottle.allow()
	return promiseWithReason{
		promise: promise,
		errWin:  lt.errWin,
	}, lt.logError(err)
}

func (lt loggedThrottle) doReq(req func() error,
	fallback func(err error) error, acceptable Acceptable) error {
	return lt.internalThrottle.doReq(req, fallback, acceptable)
}

func (lt loggedThrottle) logError(err error) error {
	if err == ErrServiceUnavailable {
		// if circuit open, not possible to have empty error window
		stat.Report(fmt.Sprintf(
			"proc(%s/%d), callee: %s, breaker is open and requests dropped\nlast errors:\n%s",
			proc.ProcessName(), proc.Pid(), lt.name, lt.errWin))
	}

	return err
}

type errorWindow struct {
	reasons [numHistoryReasons]string
	index   int
	count   int
	lock    sync.Mutex
}

func (ew *errorWindow) add(reason string) {
	ew.lock.Lock()
	ew.reasons[ew.index] = fmt.Sprintf("%s %s", timex.Time().Format(timeFormat), reason)
	ew.index = (ew.index + 1) % numHistoryReasons
	ew.count = mathx.MinInt(ew.count+1, numHistoryReasons)
	ew.lock.Unlock()
}

func (ew *errorWindow) String() string {
	var reasons []string

	ew.lock.Lock()
	// reverse order
	for i := ew.index - 1; i >= ew.index-ew.count; i-- {
		reasons = append(reasons, ew.reasons[(i+numHistoryReasons)%numHistoryReasons])
	}
	ew.lock.Unlock()

	return strings.Join(reasons, "\n")
}

type promiseWithReason struct {
	promise internalPromise
	errWin  *errorWindow
}

// Accept tells the Breaker that the call is successful.
func (p promiseWithReason) Accept() {
	p.promise.Accept()
}

// Reject tells the Breaker that the call is failed.
func (p promiseWithReason) Reject(reason string) {
	p.errWin.add(reason)
	p.promise.Reject()
}
