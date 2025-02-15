package prime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PrimeNumber() {
	intro()
	
	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("Good Bye!")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}

}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a whole number", false
	}
	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime!")
	fmt.Println("____________")
	fmt.Println("Enter a whole number. Enter q to quit!")
	prompt()
}

func prompt() {
	fmt.Print("->")
}

func isPrime(n int) (bool, string) {

	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number by definition!", n)
	}

	if n < 0 {
		return false, "Negative numbers cannot be prime number, by definition!"
	}

	for i := 2; i*i < n; i++ {

		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
