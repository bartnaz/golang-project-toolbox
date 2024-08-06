package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEach(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	applyFunction := func(p int, i int) (int, error) {
		return p + i, nil
	}

	result, err := Each(values, applyFunction)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 3, 5, 7, 9}, result)
}

func TestFilter(t *testing.T) {
	type PassFail struct {
		Pass bool
	}

	a := []PassFail{{Pass: true}, {Pass: false}, {Pass: true}, {Pass: false}}

	r1, _ := Filter(a, func(v PassFail, _ int) (bool, error) {
		return v.Pass, nil
	})

	assert.Len(t, r1, 2)
}

func TestFlat(t *testing.T) {
	a := [][]int{
		{1,
			2,
			3},
		{4,
			5,
			6},
	}

	r := Flat(a)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, r)
}

func TestReduce(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	init := 0
	applyFunction := func(p int, v int, _ int) (int, error) {
		return p + v, nil
	}

	result, err := Reduce(values, init, applyFunction)
	assert.Nil(t, err)
	assert.Equal(t, 15, result)

	values1 := []string{"test", "secondWord"}
	init1 := []string{}
	applyFunction1 := func(p []string, v string, _ int) ([]string, error) {
		p = append(p, v)
		return p, nil
	}

	result1, err := Reduce(values1, init1, applyFunction1)
	assert.Nil(t, err)
	assert.Equal(t, values1, result1)
}

func TestChunk(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	size := 2

	result := Chunk(values, size)

	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, result)
}

func TestSort(t *testing.T) {
	values := []int{1, 3, 2, 5, 4}
	applyFunction := func(p int, _ int) (string, error) {
		return fmt.Sprintf("%d", p), nil
	}

	keys, result, err := Sort(values, applyFunction)

	assert.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, keys)
	assert.Equal(t, map[string]int{"1": 1, "3": 3, "2": 2, "5": 5, "4": 4}, result)
}
