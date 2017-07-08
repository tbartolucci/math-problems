package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"strings"
	"time"
	"strconv"
)

func main() {
	fmt.Println("Who wants to solve some math problems?")

	randSource := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(randSource)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! Let's get started.\n", name)
	fmt.Println("Please type 'exit' if you want to stop")

	quit := false
	for !quit {
		quit = askQuestion(generator)
	}

	fmt.Println("Thank you for playing!")
}

func askQuestion(r *rand.Rand) bool {
	x := r.Intn(10)
	y := r.Intn(10)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("What does %d + %d = ", x,y)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if strings.Compare(answer,"exit") == 0 {
		return true
	}

	intAnswer, _ := strconv.Atoi(answer)
	if intAnswer == (x + y) {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Incorrect, the answer is: %d\n", x+y)
	}
	fmt.Println("")
	return false
}
