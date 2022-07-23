package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Setting file name and attempting to open the file
	file_name := set_quiz_file()
	dat, err := os.ReadFile(file_name)
	check_error(err)

	// Parsing the file to CSV
	r := csv.NewReader(strings.NewReader(string(dat)))

	/* Questions becomes a slice of string slices. questions[i] runs over all
	questions/answers pairs; questions[i][0] is the question and questions[i][1]
	is the answer.
	*/
	questions, err := r.ReadAll()
	check_error(err)

	fmt.Println(questions[0][0])

}
