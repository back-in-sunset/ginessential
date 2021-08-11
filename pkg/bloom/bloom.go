package bloom

import (
	"errors"
	"gin-essential/pkg/hash"
	"strconv"
)

const (
	maps = 14
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
	return nil
}

// Exsits checks if data is in f.
func (f *Filter) Exsits(data []byte) (bool, error) {
	return false, nil
}

func (f *Filter) getLocaltions(data []byte) []uint {
	localtions := make([]uint, maps)
	for i := uint(0); i < maps; i++ {
		hashValue := hash.Hash(append(data, byte(i)))
		localtions[i] = uint(hashValue % uint64(f.bits))
	}
	return localtions
}

type redisBitSet struct {
	store interface{}
	key   string
	bits  uint
}

func newRedisBitSet(key string, bits uint) redisBitSet {
	return redisBitSet{
		store: nil,
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

// func (r *redisBitSet) check(offsets []uint) (bool, error) {
// 	args, err := r.buildOffsetArgs(offsets)
// 	if err != nil {
// 		return false, err
// 	}

// 	return false, nil
// }
