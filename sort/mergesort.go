package sort

/**
算法描述：
通过递归实现
**递推公式**  
merge_sort(p…r) = merge(merge_sort(p…q), merge_sort(q+1…r))  
**终止条件**  
p >= r 不用再继续分解

优化：
1.
*/

func mergeSort(input []int) {
	recursion(input, 0, len(input)-1)
}

func recursion(input []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	recursion(input, l, mid)
	recursion(input, mid+1, r)
	merge(input, l, mid, r)
}

func merge(input []int, l, mid, r int) {
	temp := make([]int, r-l+1)

	i, j := l, mid+1
	k := 0
	for ; i <= mid && j <= r; k++ {
		if input[i] <= input[j] {
			temp[k] = input[i]
			i++
		} else {
			temp[k] = input[j]
			j++
		}
	}

	for ; i <= mid; k++ {
		temp[k] = input[i]
		i++
		k++
	}

	for ; j <= r; k++ {
		temp[k] = input[j]
		j++
	}
	copy(input[l:r+1], temp)
}
