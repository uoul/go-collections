# Go Collections Utility Package

This Go package provides a set of utility functions for working with slices and maps, offering functional programming patterns for data transformation and filtering.

## Slice Functions

The `slices` package includes the following functions:

- `Map[T, V any](col []T, fn func(T) V) []V`: Transforms each element of a slice using the provided function.
- `Filter[T any](col []T, pred func(T) bool) []T`: Returns a new slice containing only elements that satisfy the predicate.
- `Contains[T any](col []T, pred func(T) bool) bool`: Checks if any element in the slice satisfies the predicate.
- `Merge[T1, T2, V any](col1 []T1, col2 []T2, merge func(T1, T2) V) ([]V, error)`: Combines two slices of equal length using a merge function; returns an error if lengths differ.
- `GroupBy[T any, K comparable](col []T, groupBy func(item T) K) map[K][]T`: Groups slice elements by a key generated from each element.

## Map Functions

The `maps` package includes the following functions:

- `Map[T comparable, U, V comparable, W any](col map[T]U, fn func(T, U) (V, W)) map[V]W`: Transforms each key-value pair in a map using the provided function.
- `Filter[K comparable, V any](col map[K]V, pred func(K, V) bool) map[K]V`: Returns a new map containing only key-value pairs that satisfy the predicate.
- `Contains[K comparable, V any](col map[K]V, pred func(K, V) bool) bool`: Checks if any key-value pair in the map satisfies the predicate.

## Usage Examples

```go
// Example of using Map with slices
numbers := []int{1, 2, 3, 4}
doubled := slices.Map(numbers, func(n int) int { return n * 2 })
// Result: []int{2, 4, 6, 8}

// Example of using Filter with maps
ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
adults := maps.Filter(ages, func(name string, age int) bool { return age >= 30 })
// Result: map[string]int{"Bob": 30, "Charlie": 35}

```

## Installation

```bash
go get github.com/uoul/go-collections
```

## Import

```go
import (
    "github.com/uoul/go-collections/slices"
    "github.com/uoul/go-collections/maps"
)
```

