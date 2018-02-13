package patterns

import (
	"github.com/k0kubun/pp"
)

type (
	// 数据不变， 通过不同的方法得到相同的结果
	// 估计只是测试不同解法的时候有点儿用
	Strategy struct{}

	sorter struct {
		sort func([]int) []int
	}
)

func (Strategy) Do() {

	desc.SetDesc("Strategy(策略)", "数据固定，方法变化", "sort方法", "具体的sort的实现", "其实就是一个delegate", "callback, delegate都会定义一个函数指针，以便对应不同需求")
	desc.print()

	data := []int{1, 23, 5, 6, 754, 43}

	s := new(sorter)

	pp.Println("bubbleSort")
	s.sort = bubbleSort
	pp.Println(s.sort(data))

	pp.Println("insertionSort")
	s.sort = insertionSort
	pp.Println(s.sort(data))
	pp.Println("selectionSort")
	s.sort = selectionSort
	pp.Println(s.sort(data))
	pp.Println("shellSort")
	s.sort = shellSort
	pp.Println(s.sort(data))
	pp.Println("mergeSort")
	s.sort = mergeSort
	pp.Println(s.sort(data))

}

func insertionSort(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)
	size := len(output)
	for i := 1; i < size; i++ {
		pp.Println(output)
		for j := i; j > 0 && output[j] < output[j-1]; j-- {
			pp.Printf("to swap i: %s, [%s]->%s, [%s]->%s \n", i, j, output[j], j-1, output[j-1])
			output[j], output[j-1] = output[j-1], output[j]
		}
	}
	return output
}

func bubbleSort(input []int) []int {
	swapped := true
	output := make([]int, len(input))
	copy(output, input)
	size := len(output)
	for swapped {
		swapped = false
		for i := 1; i < size; i++ {
			if output[i-1] > output[i] {
				output[i], output[i-1] = output[i-1], output[i]
				swapped = true
			}
		}
	}
	return output
}

func mergeSort(input []int) []int {
	size := len(input)
	output := make([]int, size)
	copy(output, input)

	if size <= 1 {
		return output
	}

	left := make([]int, 0)
	right := make([]int, 0)
	m := size / 2

	for i, x := range output {
		switch {
		case i < m:
			left = append(left, x)
		case i >= m:
			right = append(right, x)
		}
	}

	left = mergeSort(left)
	right = mergeSort(right)

	output = merge(left, right)
	return output
}

func merge(left, right []int) []int {
	results := make([]int, 0)
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				results = append(results, left[0])
				left = left[1:]
			} else {
				results = append(results, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:]
		}
	}

	return results
}

func selectionSort(input []int) []int {
	size := len(input)
	output := make([]int, size)
	copy(output, input)

	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j <= size-1; j++ {
			if output[j] < output[min] {
				min = j
			}
		}
		output[i], output[min] = output[min], output[i]
	}

	return output
}

func shellSort(input []int) []int {
	size := len(input)
	output := make([]int, size)
	copy(output, input)

	h := 1
	for h < size {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < size; i++ {
			for j := i; j >= h && output[j] < output[j-h]; j = j - h {
				output[j], output[j-h] = output[j-h], output[j]
			}
		}

		h = h / 3
	}

	return output
}
