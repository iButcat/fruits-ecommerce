package utils

import (
	"encoding/json"
	"log"
	"regexp"
)

// basic regex that check if corrects email
func RegexEmailChecker(email string) bool {
	var validator = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	var regexEmail = regexp.MustCompile(validator)
	return regexEmail.MatchString(email)
}

// check that our different fields are not empty
func ValidateIfNotEmpty(fields ...int) bool {
	for _, field := range fields {
		if field == 0 {
			return false
		}
	}
	return true
}

// returns the string representation of any given string
func ToString(field interface{}) string {
	converted, err := json.Marshal(field)
	if err != nil {
		log.Println(err)
	}
	return string(converted)
}

func CalculateDiscountBanana(quantity int, totalPrice float64) float64 {
	var result float64
	var biggerThanSever bool
	if quantity >= 7 {
		biggerThanSever = true
	}
	if biggerThanSever {
		result = totalPrice - (totalPrice * .10)
	}
	return result
}
