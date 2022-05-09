package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	for _, name := range input {
		result += name
	}
	result = result / float32(len(input))
	return
}
