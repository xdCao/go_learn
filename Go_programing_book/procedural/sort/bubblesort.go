package sort

func BubbleSort(in []int) {
	flag := true
	for i := 0; i < len(in)-1; i++ {
		flag = true
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
