package sort

/**
算法描述：
每次从第一个位置进行冒泡，每次冒泡就是取相邻的两个数进行比较，判断是否需要交换

优化：
1. 若中间出现一次都不需要进行交换，则排序已完成
*/

func bubbleSort(input []int) {
	if len(input) < 2 {
		return
	}

	for i := 0; i < len(input); i++ {
		flag := false
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
