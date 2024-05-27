package main

import (
	"bufio"
	"fmt"
	"myproject/entity"
	"myproject/service"
	"os"
	"strings"
)

var users map[string]entity.User

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("1. Login || 2. Sign Up || 3. Exit")
		fmt.Println()
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			service.Login(users)
		case "2":
			service.SignUp(users)
		case "3":
			fmt.Println("Exiting...")
		}
	}
}
