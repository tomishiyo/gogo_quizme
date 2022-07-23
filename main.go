package main

import (
	"fmt"
	"os"

	// Import not working for some mysterious reason. Fix it
	"pkg.go.dev/encoding/csv"
)

// Gets the quiz file to be used
// Defaults to problems.csv
func set_quiz_file() string {
	if len(os.Args) == 1 {
		return "quizzes/problems.csv"

	} else if len(os.Args) == 2 {

		return os.Args[1]
	} else {
		error_msg := "Usage: " + os.Args[0] + " <quizz_name>"
		fmt.Println(error_msg)
		os.Exit(1)
		return ""
	}

}

func main() {
	file_name := set_quiz_file()
	fmt.Println(file_name)

	file := csv.Read(file_name)
}
