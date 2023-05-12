package main

import (
	"fmt"
	"time"
)

func swap(ptr2 *float64, ptr1 *float64) {
	var x float64
	x = *ptr1
	*ptr1 = *ptr2
	*ptr2 = x
}

func main() {
	var (
		x, y float64
		ptr1 *float64 = &x
		ptr2 *float64 = &y
	)
	fmt.Println("Digite o valor para X :")
	fmt.Scan(&x)
	fmt.Println("Digite o valor para Y :")
	fmt.Scan(&y)
	fmt.Println("X =", x)
	fmt.Println("Y =", y)
	fmt.Println("Os valores estao sendo trocados")
	for i := 0; i < 6; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(" ")
	if x == y {
		fmt.Println("As variáveis já possuem o mesmo valor")
		fmt.Println("ERROR")
	} else {
		swap(ptr1, ptr2)
		fmt.Println("Os valores agora são :")
		fmt.Println("X =", x)
		fmt.Println("Y =", y)
	}
}
