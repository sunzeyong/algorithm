package sort

func quickSort(input []int) {
	separate(input, 0, len(input)-1)
}

func separate(input []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(input, l, r)
	separate(input, l, pivot-1)
	separate(input, pivot+1, r)
}

// [:i)区间是已经处理好的，j位置是每次需要判断的，如果input[j]小于pivot，就将input[j]放到已经处理的区间后面
// 将input[j]放到已经处理好的位置后面使用交换方法即可
// 如果j==i说明j位置是符合要求的，而且不需要进行交换
// 最后将input[i] 和 pivot位置进行交换
func partition(input []int, l, r int) int {
	pivot := input[r]

	i := l
	for j := i; j < r; j++ {
		if input[j] < pivot {
			if i != j {
				input[i], input[j] = input[j], input[i]
			}
			i++
		}
	}
	input[i], input[r] = input[r], input[i]

	return i
}
