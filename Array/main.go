package main

func main() {
	num := []int{1, 2, 3}
	num = append(num, 3)
	for _, i := range num {
		println(i)
	}
	println(len(num))
}
