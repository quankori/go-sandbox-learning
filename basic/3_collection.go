package basic

import "fmt"

// Sự khác nhau giữa slices và array:
// Array có kích thước cố định (fixed size) và phần tử phải cùng loại dữ liệu, còn **Slices** có kích thước động.
// Array dạng tham trị (value types) khi gán biến mới sẽ tạo một array khác tốn bộ nhớ hơn, còn **Slices** dạng tham trị
func array() {
	fmt.Println("Array: ")
	a := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("array a is: ", a)
	b := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println("array a is: ", b)
}

func slices() {
	fmt.Println("Slices: ")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println("slice primes is: ", s)
	q := []int{2, 3, 5, 7, 11, 13}
	b := q[2:4]
	fmt.Println("slices cut is: ", b)
	// Append can work with nil slice.
	s = append(s, 0) // s = [0]

	// Append add one element to slice.
	s = append(s, 1) // s = [0, 1]

	// Append add multiple elements to slice.
	s = append(s, 2, 3, 4) // s = [0, 1, 2, 3, 4]
	fmt.Println("slices after append: ", s)
}
