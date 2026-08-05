package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string){
	fmt.Println("Goodbye", name)
}

func main(){
	sayHello("Bob")
	fmt.Println("Hows the weather?")
	sayGoodbye("Bob")
}
