package main

import "flag"

// scan given a path crawls it and its subfolders
// searching for Git repositories
func scan(path string) {
	print("scan")
}

// stats generates a nice graph of your Git contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for GIT repositories")
	flag.StringVar(&email, "email", "qwer@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)

	/*
	✗ go run main.go
	stats%                                                                                                                                                               ➜  git-stat git:(main) ✗ go run main.go -add sd
	✗ go run main.go -add sd
	scan%
	*/
}
