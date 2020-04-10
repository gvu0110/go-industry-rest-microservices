package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// Initialization
	elements := []int{9, 8, 7, 6, 5}

	// Execution
	BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func createElements(n int) []int {
	elements := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		elements[i] = j
		i++
	}
	return elements
}

func TestCreateElements(t *testing.T) {
	elements := createElements(5)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 4, elements[0])
	assert.EqualValues(t, 3, elements[1])
	assert.EqualValues(t, 2, elements[2])
	assert.EqualValues(t, 1, elements[3])
	assert.EqualValues(t, 0, elements[4])
}

func BenchmarkBubbleSort10Elements(b *testing.B) {
	elements := createElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
} // 9.27 ns/op

func BenchmarkGoSort10Elements(b *testing.B) {
	elements := createElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
} // 115 ns/op

func BenchmarkBubbleSort1000Elements(b *testing.B) {
	elements := createElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
} // 627 ns/op

func BenchmarkGoSort1000Elements(b *testing.B) {
	elements := createElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
} // 48,998 ns/op

func BenchmarkBubbleSort10000Elements(b *testing.B) {
	elements := createElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
} // 6673 ns/op

func BenchmarkGoSort10000Elements(b *testing.B) {
	elements := createElements(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
} // 630,203 ns/op

func BenchmarkBubbleSort50000Elements(b *testing.B) {
	elements := createElements(50000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
} // 4,002,598,679 ns/op

func BenchmarkGoSort50000Elements(b *testing.B) {
	elements := createElements(50000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
} // 3,567,262 ns/op

func BenchmarkBubbleSort1000000000Elements(b *testing.B) {
	elements := createElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
} // 9,933,754,013 ns/op

func BenchmarkGoSort100000Elements(b *testing.B) {
	elements := createElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
} // 8,279,014 ns/op
