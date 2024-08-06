package slices

import "sort"

func Each[T any, R any](values []T, fn func(item T, indx int) (R, error)) ([]R, error) {
	result := []R{}

	for i, v := range values {
		v := v
		r, err := fn(v, i)
		if err != nil {
			return result, err
		}
		result = append(result, r)
	}
	return result, nil
}

// saves elements to array if they pass specific function
func Filter[T any](values []T, fn func(item T, indx int) (bool, error)) ([]T, error) {
	result := []T{}

	for i, v := range values {
		v := v
		ok, err := fn(v, i)
		if err != nil {
			return result, err
		}

		if ok {
			result = append(result, v)
		}

	}

	return result, nil
}

func Flat[T any](values [][]T) (result []T) {
	for _, v := range values {
		v := v
		result = append(result, v...)
	}
	return
}

func Reduce[T any, R any](values []T, init R, fn func(previous R, item T, indx int) (R, error)) (R, error) {
	last := init // a value to append to

	for i, v := range values {
		// v := v
		p, err := fn(last, v, i)
		if err != nil {
			return p, err
		}
		last = p
	}

	return last, nil // return modified value
}

func Chunk[T any](values []T, size int) [][]T {
	var slices [][]T
	dataLength := len(values)
	for i := 0; i < dataLength; i += size {
		end := i + size
		if end > dataLength {
			end = dataLength
		}
		slices = append(slices, values[i:end])
	}
	return slices
}

func Sort[T any](values []T, fn func(value T, indx int) (string, error)) ([]string, map[string]T, error) {
	result := map[string]T{}
	keys := []string{}

	for i, v := range values {
		v := v
		key, err := fn(v, i)
		if err != nil {
			return keys, result, err
		}

		keys = append(keys, key)
		result[key] = v
	}

	sort.Strings(keys)
	return keys, result, nil
}
