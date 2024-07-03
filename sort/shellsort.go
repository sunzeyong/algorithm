package sort

import (
	"fmt"
)

/**
算法描述：
插入排序：
将数组分为两个区间，已排序区间和未排序区间
从第二个元素开始，依次对已排序区间从后往前进行对比，如果大于当前数，就往后移一位，直到不大于。
再将当前数字放在空出来的位置
希尔排序：
在插入排序的基础上实现，即将input进行分组，分别进行插入排序 最终实现O(n^1.3~2)的时间复杂度

*/

func shellSort(input []int) {
	if len(input) < 2 {
		return
	}

	for step := len(input) / 2; step >= 1; step /= 2 {

		for i := step; i < len(input); i++ {
			cur := input[i]

			j := i - step
			for ; j >= 0; j -= step {
				if input[j] > cur {
					input[j+step] = input[j]
				} else {
					break
				}
			}
			input[j+step] = cur
		}

		fmt.Printf("step: %v, arr: %v \n", step, input)

	}
}
