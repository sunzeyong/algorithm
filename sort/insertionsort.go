package sort

/**
算法描述：
将数组分为两个区间，已排序区间和未排序区间
从第二个元素开始，依次对已排序区间从后往前进行对比，如果大于当前数，就往后移一位，直到不大于。
再将当前数字放在空出来的位置

优化：
1. 参考希尔排序
*/

func insertionSort(input []int) {
	if len(input) < 2 {
		return
	}

	for i := 1; i < len(input); i++ {
		cur := input[i]

		j := i - 1
		for ; j >= 0; j-- {
			if input[j] > cur {
				input[j+1] = input[j]
			} else {
				break
			}
		}
		input[j+1] = cur
	}
}
