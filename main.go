package main

import (
	"bufio"
	"fmt"
	"os"
)

var tasks []string

func main(){
	item := -1 
	for item != 0 {
		item = menu()
		switch(item){
		case 1 : 
			showList()
			break
		case 2:
			addItem()
		case 3: 
			removeItem()
		case 0:
			break
		default:
			fmt.Println("Enter a valid option")
		}
	}
}

func menu() int {
	var input int
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
	fmt.Scan(&input)
	return input
}

func showList(){
	fmt.Println()
	fmt.Println("----------------------")
	fmt.Println("     To-Do List")
	fmt.Println("----------------------")
	number := 0
	for _, element := range tasks{
		number = number + 1
		fmt.Printf("%d  %s", number, element)
	} 
	fmt.Println()
	fmt.Println("----------------------")
}

func addItem(){
	fmt.Println()
	fmt.Println("Add Item")
	fmt.Println("----------------------")
	fmt.Print("Enter an item: ")
	reader := bufio.NewReader(os.Stdin)
    userInput, _ := reader.ReadString('\n')
	tasks = append(tasks, userInput)
}

func removeItem(){

}