package utils

// Check checks errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
