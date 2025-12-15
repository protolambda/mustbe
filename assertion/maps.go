package assertion

import (
	"context"
	"fmt"
)

// InMap asserts that the Key is present in Map
type InMap[K comparable, V any] struct {
	Key K
	Map map[K]V
}

var _ Assertion = InMap[int, int]{}

func (i InMap[K, V]) String() string {
	return fmt.Sprintf("isInMap(key: %v, map: %d entries)", i.Key, len(i.Map))
}

func (i InMap[K, V]) Check(ctx context.Context) error {
	if _, ok := i.Map[i.Key]; ok {
		return nil
	}
	return fmt.Errorf("key %v not found in map", i.Key)
}
