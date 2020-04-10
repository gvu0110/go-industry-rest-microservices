package utils

import "sort"

func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

func Sort(elements []int) {
	if len(elements) < 50000 {
		BubbleSort(elements)
		return
	}
	sort.Ints(elements)
}
