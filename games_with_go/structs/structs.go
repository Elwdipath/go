package main

import "fmt"

type Position struct {
	x float32
	y float32
}

func main(){
	// p.x = 5
	// p.y = 4
	p := Position{4, 2}
	fmt.Println(p.x, p.y)
}
