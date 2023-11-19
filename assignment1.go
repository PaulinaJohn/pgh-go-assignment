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
	fmt.Println(checkVowel("W"))
}

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
