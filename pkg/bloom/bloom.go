package bloom

import (
	"gin-essential/pkg/errors"
	"gin-essential/pkg/hashx"
	"sort"
	"strconv"

	"github.com/go-redis/redis"
)

const (
	maps      = 14
	setScript = `
for _, offset in ipairs(ARGV) do
	redis.call("setbit", KEYS[1], offset, 1)
end
`
	testScript = `
for _, offset in ipairs(ARGV) do
	if tonumber(redis.call("getbit", KEYS[1], offset)) == 0 then
		return false
	end
end
return true
`
)

// ErrTooLargeOffset indicates the offset is too large in bitset.
var ErrTooLargeOffset = errors.New("too large offset")

type (
	// Filter 过滤器
	Filter struct {
		bits   uint
		bitSet bitSetProvider
	}

	bitSetProvider interface {
		check([]uint) (bool, error)
		add([]uint) error
	}
)

// New 实例化
func New(store interface{}, key string, bits uint) *Filter {
	return &Filter{
		bits:   bits,
		bitSet: nil,
	}
}

// Add adds data into f.
func (f *Filter) Add(data []byte) error {
	localtions := f.getLocaltions(data)
	return f.bitSet.add(localtions)
}

// Exsits checks if data is in f.
func (f *Filter) Exsits(data []byte) (bool, error) {
	localtions := f.getLocaltions(data)
	isSet, err := f.bitSet.check(localtions)
	if err != nil {
		return false, err
	}
	if !isSet {
		return false, nil
	}

	return true, nil
}

func (f *Filter) getLocaltions(data []byte) []uint {
	localtions := make([]uint, maps)
	for i := uint(0); i < maps; i++ {
		hashValue := hashx.Hash(append(data, byte(i)))
		localtions[i] = uint(hashValue % uint64(f.bits))
	}
	return localtions
}

type store interface {
	Eval(script string, key string, args []string) (interface{}, error)
}

type redisBitSet struct {
	store store
	key   string
	bits  uint
}

type cache struct {
	m map[string][]string
}

var memcache cache

func (r cache) Eval(script string, key string, args []string) (interface{}, error) {
	sort.Strings(args)

	storedArgs, ok := r.m[key]
	if !ok {
		return 0, redis.Nil
	}

	for k, storedArg := range storedArgs {
		if storedArg != args[k] {
			return 0, nil
		}
	}

	return 1, nil
}

func newRedisBitSet(key string, bits uint) redisBitSet {
	return redisBitSet{
		store: memcache,
		key:   key,
		bits:  bits,
	}
}

func (r *redisBitSet) buildOffsetArgs(offset []uint) ([]string, error) {
	var args []string

	for _, offset := range offset {
		if offset >= r.bits {
			return nil, ErrTooLargeOffset
		}
		args = append(args, strconv.FormatUint(uint64(offset), 10))
	}

	return args, nil
}

func (r *redisBitSet) check(offsets []uint) (bool, error) {
	args, err := r.buildOffsetArgs(offsets)
	if err != nil {
		return false, err
	}

	resp, err := r.store.Eval(testScript, r.key, args)
	if err != nil {
		return false, err
	}

	exists, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return exists == 1, nil
}

func (r *redisBitSet) set(offsets []uint) error {
	args, err := r.buildOffsetArgs(offsets)
	if err != nil {
		return err
	}
	_, err = r.store.Eval(setScript, r.key, args)
	if err == redis.Nil {
		return nil
	}

	return nil
}
