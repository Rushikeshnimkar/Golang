package swap

import "fmt"

// func Swap(x, y string) (string, string) {
// 	return y, x
// }

func Lodu(x, y int) (int, int) {
	temp := x
	x = y
	y = temp
	return x, y

}
func San() {
	fmt.Println(Lodu(12, 23))
}
