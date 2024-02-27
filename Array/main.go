package main

func main() {
	num := []int{1, 2, 3}

	num2 := make([]int, 3)

	for i := 0; i < len(num); i++ {
		print(i, " ")
	}
	for _, i := range num2 {
		println(i)
	}

	println(len(num2), len(num))
}
