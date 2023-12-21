package mytypes

type MyInt int
// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (i MyString) Len() int {
	return len(i)
}