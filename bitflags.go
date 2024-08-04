// Package bitflags provides a generic way to manage sets of flags.
package bitflags

import (
	"fmt"
)

// FlagType is a generic constraint that includes unsigned integer types.
type FlagType interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// FlagSet represents a set of flags.
type FlagSet[T FlagType] struct {
	flags   T
	flagMap map[string]T
}

// NewFlagSetFromMap creates a new FlagSet from a map of flag names to their values.
func NewFlagSetFromMap[T FlagType](flagMap map[string]T) *FlagSet[T] {
	return &FlagSet[T]{flagMap: flagMap}
}

// NewFlagSetFromSlice creates a new FlagSet from a slice of flag names.
// The flags are assigned values of consecutive powers of two.
func NewFlagSetFromSlice(flags []string) *FlagSet[uint64] {
	flagMap := make(map[string]uint64)
	var bit uint64 = 1
	for _, name := range flags {
		flagMap[name] = bit
		bit <<= 1
	}
	return &FlagSet[uint64]{flagMap: flagMap}
}

// GetValue returns the current value of the flags.
func (fs *FlagSet[T]) GetValue() T {
	return fs.flags
}

// SetByName sets the flag with the given name.
func (fs *FlagSet[T]) SetByName(name string) error {
	flag, exists := fs.flagMap[name]
	if !exists {
		return fmt.Errorf("unknown flag: %s", name)
	}
	fs.flags |= flag
	return nil
}

// ClearByName clears the flag with the given name.
func (fs *FlagSet[T]) ClearByName(name string) error {
	flag, exists := fs.flagMap[name]
	if !exists {
		return fmt.Errorf("unknown flag: %s", name)
	}
	fs.flags &^= flag
	return nil
}

// HasByName checks if the flag with the given name is set.
func (fs *FlagSet[T]) HasByName(name string) (bool, error) {
	flag, exists := fs.flagMap[name]
	if !exists {
		return false, fmt.Errorf("unknown flag: %s", name)
	}
	return fs.flags&flag != 0, nil
}

func (fs *FlagSet[T]) HasByValue(value T) bool {
	return fs.flags&value != 0
}

// ToggleByName toggles the flag with the given name.
func (fs *FlagSet[T]) ToggleByName(name string) error {
	flag, exists := fs.flagMap[name]
	if !exists {
		return fmt.Errorf("unknown flag: %s", name)
	}
	fs.flags ^= flag
	return nil
}

// HasAnyByName checks if any of the flags with the given names are set.
func (fs *FlagSet[T]) HasAnyByName(names ...string) (bool, error) {
	for _, name := range names {
		flag, exists := fs.flagMap[name]
		if !exists {
			return false, fmt.Errorf("unknown flag: %s", name)
		}
		if fs.flags&flag != 0 {
			return true, nil
		}
	}
	return false, nil
}

// HasAllByName checks if all of the flags with the given names are set.
func (fs *FlagSet[T]) HasAllByName(names ...string) (bool, error) {
	for _, name := range names {
		flag, exists := fs.flagMap[name]
		if !exists {
			return false, fmt.Errorf("unknown flag: %s", name)
		}
		if fs.flags&flag == 0 {
			return false, nil
		}
	}
	return true, nil
}

// SetByValue sets flags according to the given value.
func (fs *FlagSet[T]) SetByValue(value T) error {
	var bit T = 1
	for ; bit != 0; bit <<= 1 {
		if value&bit != 0 {
			flagFound := false
			for _, flag := range fs.flagMap {
				if flag == bit {
					fs.flags |= bit
					flagFound = true
					break
				}
			}
			if !flagFound {
				return fmt.Errorf("invalid flag bit: %08b", bit)
			}
		}
	}
	return nil
}

// String returns the string representation of the current flag value.
func (fs *FlagSet[T]) String() string {
	return fmt.Sprintf("%08b", fs.flags)
}

// GetActiveFlags returns a map of the currently active flags.
func (fs *FlagSet[T]) GetActiveFlags() map[string]T {
	result := make(map[string]T)
	for name, flag := range fs.flagMap {
		if fs.flags&flag != 0 {
			result[name] = flag
		}
	}
	return result
}
