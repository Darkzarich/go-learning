package mathslice

type Element int

type Slice []Element

func SumSlice(slice Slice) (res Element) {
	for _, v := range slice {
		res += v
	}
	return res
}

func MapSlice(slice Slice, f func(Element) Element) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}

func AverageSlice(slice Slice) float64 {
	sum := SumSlice(slice)
	return float64(sum) / float64(len(slice))
}
