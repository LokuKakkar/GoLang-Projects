package main

// run with go run .
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello World")
	x := rand.Intn(10)
	fmt.Println(x)

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// array
	xi := []int{2, 3, 5, 2, 4132}
	for i, v := range xi {
		fmt.Println("index: ", i, "value: ", v)
	}

	// map
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	for k, v := range m {
		fmt.Println("key: ", k, "value: ", v)
		fmt.Println(m["apple"])
	}

	main2()
	main3()
	main4()

}
