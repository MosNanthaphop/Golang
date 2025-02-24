package main

import (
	"fmt"
)

type Student struct {
	name  (string)
	score (float64)
}

func calculateGrade(score float64) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	}
	return "F"
}

func (s Student) getGrade() string {
	return calculateGrade(s.score)
}

func main() {
	Students := make([]Student, 4)
	Students[0] = Student{"test", 60}
	Students[1] = Student{"Somchai", 70}
	Students[2] = Student{"Jai", 90}
	Students[3] = Student{"Metee", 62}

	for _, s := range Students {
		fmt.Printf("ชื่อ %s, คะแนน %.2f, เกรด %s\n", s.name, s.score, s.getGrade())
	}
}
