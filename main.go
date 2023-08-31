package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var tasks []string

func main(){
	item := -1 
	for item != 0 {

	}
}

func menu() int {
	fmt.Println()
	fmt.Println("----------------------")
	fmt.Println("     Main Menu")
	fmt.Println("----------------------")
	fmt.Println("0. Exit the program")
	fmt.Println("1. Display to-do list")
	fmt.Println("2. Add item to list")
	fmt.Println("3. Remove item from list")
	fmt.Println()
	fmt.Println("Enter choice: ")
	scanner := bufio.NewScanner(os.Stdin)
	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return input
}