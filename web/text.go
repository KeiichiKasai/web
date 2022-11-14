package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var database = map[string]string{
	"A": "123456",
}
var username = make(chan string, 1000)
var password = make(chan string, 1000)

func main() {
	file1, err := os.Open("用户名.txt")

	// handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file1.Close()

	fileScanner1 := bufio.NewScanner(file1)

	// read line by line
	for fileScanner1.Scan() {
		username <- fileScanner1.Text()
	}
	// handle first encountered error while reading
	if err := fileScanner1.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	//_________________________________________________
	file2, err := os.Open("密码.txt")

	// handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file2.Close()

	fileScanner2 := bufio.NewScanner(file2)

	// read line by line
	for fileScanner2.Scan() {
		password <- fileScanner2.Text()
	}
	// handle first encountered error while reading
	if err := fileScanner2.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	for {
		A := <-username
		B := <-password
		database[A] = B
		if len(username) == 0 {
			break
		}
	}
	fmt.Println(database)
}
