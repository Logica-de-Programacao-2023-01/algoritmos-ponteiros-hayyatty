package main

import (
	"fmt"
	"time"
)

func inverter(ptr *bool) {
	if *ptr == false {
		*ptr = true
	} else if *ptr == true {
		*ptr = false
	}

}

func main() {
	var (
		x   bool
		ptr *bool = &x
	)
	fmt.Println("A informação é true ou false :")
	fmt.Scan(&x)
	inverter(ptr)
	for i := 0; i < 6; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(" ")
	fmt.Println("Baseado nas nossas analises essa info é :")
	fmt.Println(x)

}
