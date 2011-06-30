package count

//Count1 convert string to int array.
func Count1(s string) int {
	return len([]int(s))
}

//Count2 count characters of string.
func Count2(s string) int {
	c := 0
	for _, _ = range s {
		c++
	}
	return c
}
