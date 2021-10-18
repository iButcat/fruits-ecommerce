package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"ecommerce/models"
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

func UnmarshalRequestData(r *http.Request, model interface{}) (interface{}, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return data, err
	}
	json.Unmarshal(data, &model)
	return model, nil
}

// returns the string representation of any given string
func ToString(field interface{}) string {
	converted, err := json.Marshal(field)
	if err != nil {
		log.Println(err)
	}
	return string(converted)
}

func CalculateDiscountApples(name string, quantity int, price float64) float64 {
	var result float64
	if name == "apples" {
		var biggerThanSeven bool
		if quantity >= 7 {
			biggerThanSeven = true
		}
		if biggerThanSeven {
			result = (float64(quantity) * price / 100)
		}
	} else {
		return 0
	}
	return result
}

func CalculateDiscountSet(setFruits map[string]int, price float64) float64 {
	var isSet = make([]bool, 2)
	for name, quantity := range setFruits {
		if name == "pears" && quantity >= 4 {
			isSet = append(isSet, true)
		} else if name == "bananas" && quantity >= 2 {
			isSet = append(isSet, true)
		}
	}
	if len(isSet) == 2 && isSet[0] && isSet[1] {
		price = price / 300
	}
	return price
}

func CalculateTotalPriceCart(cart models.Cart) float64 {
	var totalPrice float64
	for _, cartItem := range cart.CartItems {
		totalPrice += cartItem.TotalPrice
	}
	return totalPrice
}
