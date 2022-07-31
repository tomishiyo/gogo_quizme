package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var score int
var time_is_up_c chan bool

// Process the quiz file to be used
// Defaults to problems.csv
func set_quiz_file() string {
	// flag.Args() gives remaining arguments after flag parsing.
	remaining_args := flag.Args()
	if len(remaining_args) == 0 {
		return "quizzes/problems.csv"

	} else {
		return remaining_args[0]
	}
}

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

// Runs the quiz: loop over all questions prompting the user for an answer
func run_quiz(questions [][]string, timer_time int) (score int) {
	var answer, question_text string
	var user_ok bool

	fmt.Println("")
	fmt.Println(">>>>>>>>Welcome to GOGO_QUIZME!<<<<<<<<<<<<")
	fmt.Println("")
	fmt.Println("Rule N.1 : All numerical answers must be numbers!")
	fmt.Println("For example, 1+1 must be answered as 2 not as Two")
	fmt.Println("")
	fmt.Println("Rule N.2: Ordinal numbers must be writen in abbreviated form!")
	fmt.Println("For example, Neil Armstrong was the 1st man in the moon (as opossed to 'first')")
	fmt.Println("")

	fmt.Println("Timer set to " + fmt.Sprint(timer_time) + " seconds !")
	fmt.Println("Press any key to start the quizz...")

	fmt.Scanln(&user_ok)
	fmt.Println("Have fun!")

	// Starting timer
	timer := time.NewTimer(time.Duration(timer_time * 1e9)) // Time in ns

	for i, question := range questions {
		question_text = "Question " + fmt.Sprint(i+1) + ": " + question[0]
		fmt.Println(question_text)
		// Print without the \n allows for the formatted input
		fmt.Print(">>")
		answerCh := make(chan string)
		go func() {
			var user_answer string
			fmt.Scanf("%s\n", &user_answer)
			// strings.ToLower remove caps problems
			user_answer = strings.ToLower(user_answer)
			user_answer = strings.TrimSpace(user_answer)
			answerCh <- user_answer
		}()

		select {
		// If timer, stop the program
		case <-timer.C:
			fmt.Println("")
			fmt.Println("\nTime is up!")
			fmt.Println("")
			return score
			// Else compute the score
		case user_answer := <-answerCh:
			answer = strings.ToLower(question[1])
			if answer == user_answer {
				score += 1
			}
		}

	}

	return score
}

func print_score(score, number_of_questions int) {
	fmt.Println("And DONE ...")
	fmt.Printf("In total, %v questions were answered!\n", number_of_questions)
	fmt.Println("")
	fmt.Println("#######################################")
	fmt.Printf("              You scored %v/%v!          \n", score, number_of_questions)
	fmt.Println("#######################################")
}

func main() {
	// Quiz time (settled by flag or defaulted to 30 seconds)
	timer_time := flag.Int("time", 30, "Quiz timer")
	flag.Parse()
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

	score := run_quiz(questions, *timer_time)
	print_score(score, number_of_questions)

}
