// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(gradeStudentSwitch(22))
	fmt.Println(checkVowel("O"))
	fmt.Println(checkTicketPrice(34))
}

// 1. Write a program that takes a student's score as input and converts it into a letter grade according to the following scale:

// A: 90-100
// B: 80-89
// C: 70-79
// D: 60-69
// F: 0-59

func gradeStudent(score int) string {
	var grade string
	if score >= 90 && score <= 100 {
		grade = "A"
	} else if score >= 80 && score <= 89 {
		grade = "B"
	} else if score >= 70 && score <= 79 {
		grade = "C"
	} else if score >= 60 && score <= 69 {
		grade = "D"
	} else if score >= 0 && score <= 59 {
		grade = "F"
	}

	return grade
}

func gradeStudentSwitch(score int) string {
	var grade string
	switch {
	case score >= 90 && score <= 100:
		grade = "A"
	case score >= 80 && score <= 89:
		grade = "B"
	case score >= 70 && score <= 79:
		grade = "C"
	case score >= 60 && score <= 69:
		grade = "D"
	case score >= 0 && score <= 59:
		grade = "F"
	}

	return grade
}

// 2. Write a program that takes a year as input and determines whether it is a leap year or not.

// func checkLeapYear(year int) Bool{
// }

//3. 3. Write a program that takes a single character as input and determines whether it is a vowel or a consonant.
// Use if statements and boolean values to implement this vowel checker.

func checkVowel(letter string) string {
	letter = strings.ToLower(letter)
	vowels := [5]string{"a", "e", "i", "o", "u"}
	var letterType string
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == letter {
			letterType = "vowel"
			break
		} else {
			letterType = "consonant"
		}
	}
	return letterType
}

// 4. Write a Go program that takes a person's age as input and determines the ticket price for a movie theater based on the following rules:
// Children (age <= 12): N5
// Adults (age > 12 and age <= 64): N10
// Seniors (age >= 65): N7

func checkTicketPrice(age int) int {
	var price int
	if age <= 12 {
		price = 5
	} else if age > 12 && age <= 64 {
		price = 10
	} else {
		price = 7
	}

	return price
}
