package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	result = make([]int64, 0, len(input))
	for i := len(input) - 1; ; i-- {
		result = append(result, input[i])
		if i == 0 {
			break
		}
	}
	return
}
