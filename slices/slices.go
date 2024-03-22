package slices

import "strings"

// Map iterates and transforms the origin array (a) into a new array that consists of the
// transformed using the transformer (t).
func Map[T any, R any](a []T, t func(v T) R) []R {
	final := make([]R, len(a))
	for i, v := range a {
		v := v
		final[i] = t(v)
	}
	return final
}

// TriDimensional creates a three-dimensional array, similar to an array.
func TriDimensional[T any](v ...T) [][]T {
	var final [][]T
	for i, val := range v {
		if i%2 != 0 {
			continue
		}
		final = append(final, []T{val, v[i+1]})
	}
	return final
}

// AllMatches checks whether all the values in the array fits the predicate condition.
func AllMatches[T any](a []T, predicate func(b T) bool) []T {
	var l = []T{}
	for _, v := range a {
		v := v
		if predicate(v) {
			l = append(l, v)
		}
	}
	return l
}

// Filter iterates through the array and finds the first item that matches the predicate condition.
func Filter[T any](a []T, predicate func(b T) bool) *T {
	for _, v := range a {
		v := v
		if predicate(v) {
			return &v
		}
	}
	return nil
}

// AnyMatch iterates through the array up until the point where there is a single match.
func AnyMatch[T any](a []T, predicate func(b T) bool) bool {
	match := Filter(a, predicate)
	return match != nil
}

// AnyMatchComparable iterates through the array up until the point where an item matches the value (c).
func AnyMatchComparable[T comparable](a []T, c T) bool {
	match := Filter(a, func(b T) bool {
		return b == c
	})
	return match != nil
}

// FilterSlice is a short-hand of Filter that filters using an equality comparison.
func FilterSlice[T comparable](a []T, b T) *T {
	return Filter(a, func(c T) bool {
		return b == c
	})
}

// AnyMatchStringCaseInsensitive scans through the slice and finds the first equal case-insensitive match.
func AnyMatchStringCaseInsensitive(a []string, match string) bool {
	return AnyMatch(a, func(b string) bool {
		return strings.EqualFold(b, match)
	})
}
