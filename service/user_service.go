package service

import (
	"bufio"
	"fmt"
	"myproject/entity"
	"os"
	"regexp"
	"strings"
)

func RegisterUser(user entity.User) string {
	return fmt.Sprintf("User %s registered successfully!", user.Username)
}

func ValidateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func SignUp(users map[string]entity.User) string {
	var newUser entity.User
	if users == nil {
		users = make(map[string]entity.User)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		if !ValidateEmail(email) {
			fmt.Println("Invalid Email Format! Please try again...")
			continue
		}

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Enter Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		newUser.Email = email
		newUser.Username = username
		newUser.Password = password

		users[username] = newUser
		return "User registered successfully!"
	}

}

func Login(users map[string]entity.User) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println()

	fmt.Println("Enter Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user, exists := users[username]
	if !exists {
		fmt.Println("\n***** Invalid username! *****")
		fmt.Println("\nWant to regist? Y/N")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "Y" {
			SignUp(users)
		}
		return
	}

	if user.Password != password {
		fmt.Println("Incorrect Password!")
		return
	}

	fmt.Printf("Welcome %s\n", user.Username)
}
