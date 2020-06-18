package common

func SliceInsert(s []int, index int, value int) []int {
	rear := append([]int{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}
