package count

//LenViaIntArr convert string to int array.
func LenViaIntArr(s string) int {
	return len([]int(s))
}

//LenViaForRange characters of string.
func LenViaForRange(s string) int {
	c := 0
	for _, _ = range s {
		c++
	}
	return c
}

//LenViaUTF8 counts the non-0b10xxxxxx bytes of the string.
func LenViaUTF8(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if (s[i] & 0xC0) != 0x80 {
			c++
		}
	}
	return c
}
