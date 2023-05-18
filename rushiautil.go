package rushiautil

import (
	"context"
	"time"
)

// String returns a pointer string for SQL nullable.
func String(v string) *string {
	return &v
}

// Bool returns a pointer boolean for SQL nullable.
func Bool(v bool) *bool {
	return &v
}

// Int returns a pointer int for SQL nullable.
func Int(v int) *int {
	return &v
}

// SliceInt returns a pointer slice of ints for SQL nullable.
func SliceInt(v []int) *[]int {
	return &v
}

// SliceString returns a pointer string for SQL nullable.
func SliceString(v []string) *[]string {
	return &v
}

// Float64 returns a pointer float64 for SQL nullable.
func Float64(v float64) *float64 {
	return &v
}

// Time returns a pointer time.Time for SQL nullable.
func Time(v time.Time) *time.Time {
	return &v
}

// ResolveContext
func ResolveContext[T any, M any](c context.Context, data []T, fetch func(T) string, keyer func(T, M) bool, swap func(T, M), solver func(context.Context, []string) ([]M, error)) error {
	//
	if len(data) == 0 {
		return nil
	}
	//
	ids := make([]string, len(data))
	for i, v := range data {
		id := fetch(v)
		if id == "" {
			continue
		}
		ids[i] = id
	}
	//
	solved, err := solver(c, ids)
	if err != nil {
		return err
	}
	//
	for _, v := range data {
		for _, j := range solved {
			if keyer(v, j) {
				swap(v, j)
				break
			}
		}
	}
	return nil
}

// Resolve
func Resolve[T any, M any](data []T, fetch func(T) string, keyer func(T, M) bool, swap func(T, M), solver func([]string) ([]M, error)) error {
	newSovler := func(_ context.Context, ids []string) ([]M, error) {
		return solver(ids)
	}
	return ResolveContext(context.Background(), data, fetch, keyer, swap, newSovler)
}
