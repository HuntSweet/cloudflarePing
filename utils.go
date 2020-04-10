package main

//排序
func quickSortBySucnums(r []result) []result {
	if len(r) < 2{
		return r
	}
	i := 0
	j := len(r) - 1
	key := r[0]

	for {
		for r[j].sucNums > key.sucNums{
			j--
		}
		for r[i].sucNums < key.sucNums{
			i ++
		}

		if i >= j{
			break
		}

		r[j],r[i] = r[i],r[j]
		j--
		i++
	}

	quickSortBySucnums(r[:i])
	quickSortBySucnums(r[j+1:])

	return r
}



//按延迟排序，最低的排在前面
func quickSortByLantency(r []result)  {
	if len(r) < 2{
		return
	}

	i := 0
	j := len(r)-1
	key := r[0].latency

	for{
		for r[i].latency > key{
			i++
		}
		for r[j].latency < key{
			j--
		}

		if i >= j{
			break
		}

		r[i],r[j] = r[j],r[i]
		i++
		j--
	}

	quickSortByLantency(r[:i])
	quickSortByLantency(r[j+1:])

}



//按成功率进行分组，成功率相同组再进行延迟排序
func getPartion(r []result) ([]int) {
	i := 0
	j := 1
	res := []int{}
	for {
		if j > len(r) - 1{
			break
		}

		if r[i].sucRate != r[j].sucRate {
			res = append(res,j)
			i=j
		}
		j++
	}

	return res
}

func avarage(l []int) int {
	if len(l) == 0{
		return -1
	}
	re := 0
	for _,v := range l{
		re += v
	}

	return re / len(l)
}

