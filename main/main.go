package main

import (
	"fmt"
	git "homework/Uyga_vazifa22/git"
	"os"
)

func main() {
	userName, err := git.GetUserName()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create("username.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(userName)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	userEmail, err := git.GetUserEmail()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err = os.Create("useremail.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(userEmail)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
}
