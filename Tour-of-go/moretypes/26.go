// Implemented by Hunachi

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	v := 0
	prev := 0
	return func() int {
		if v == 0 {
			v = 1
			return prev
		}
		ans := v + prev
		prev = v
		v = ans
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}