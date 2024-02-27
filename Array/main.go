package main

func main() {
	num := []int{1, 2, 3}
	for _, i := range num {
		println(i)
	}
	num2 := make([]int, 3)

	for _, i := range num {
		num2 = append(num2, i*i)
	}
	println(num2)
}
