package sort

/**
算法描述：
将数组分为两个区间，已排序区间和未排序区间
从第一个元素开始，依次从未排序区间找最小值对应的下标
然后交换数据

优化：
1.
*/

func selectionSort(input []int) {
	if len(input) < 2 {
		return
	}

	for i := 0; i < len(input); i++ {
		minIdx := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIdx] {
				minIdx = j
			}
		}
		input[minIdx], input[i] = input[i], input[minIdx]
	}
}
