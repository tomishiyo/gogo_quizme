package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Process the quiz file to be used
// Defaults to problems.csv
func set_quiz_file() string {
	if len(os.Args) == 1 {
		return "quizzes/problems.csv"

	} else if len(os.Args) == 2 {

		return os.Args[1]
	} else {
		error_msg := "Usage: " + os.Args[0] + " <quizz_name>"
		panic(error_msg)
	}

}

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

// Runs the quiz: loop over all questions prompting the user for an answer
func run_quiz(questions [][]string) (score int) {
	var user_answer, answer, question_text string

	fmt.Println("")
	fmt.Println(">>>>>>>>Welcome to GOGO_QUIZME!<<<<<<<<<<<<")
	fmt.Println("")
	fmt.Println("Rule N.1 : All numerical answers must be numbers!")
	fmt.Println("For example, 1+1 must be answered as 2 not as Two")
	fmt.Println("")
	fmt.Println("Rule N.2: Ordinal numbers must be writen in abbreviated form!")
	fmt.Println("For example, Neil Armstrong was the 1st man in the moon (as opossed to 'first')")
	fmt.Println("")
	fmt.Println("Have fun!")

	for i, question := range questions {
		question_text = "Question " + fmt.Sprint(i+1) + ": " + question[0]
		fmt.Println(question_text)
		// Print without the \n allows for the formatted input
		fmt.Print(">>")
		// The Scanln does not capture the \n input
		fmt.Scanln(&user_answer)
		// strings.ToLower remove caps problems
		user_answer = strings.ToLower(user_answer)

		answer = strings.ToLower(question[1])
		if answer == user_answer {
			score += 1
		}
	}

	return score
}

func print_score(score, number_of_questions int) {
	fmt.Println("And DONE ...")
	fmt.Printf("In total, %v questions were answered!\n", number_of_questions)
	fmt.Println("")
	fmt.Println("#######################################")
	fmt.Printf("              You scored %v!          \n", score)
	fmt.Println("#######################################")
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
	number_of_questions := len(questions)

	score := run_quiz(questions)
	print_score(score, number_of_questions)

}
