package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Who wants to solve some math problems?")
	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	fmt.Println("Welcome " + name + "! Let's get started.")

	
}
