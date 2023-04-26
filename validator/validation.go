package validator

import "strings"

func OrderStatusValid(status string) bool {
	if strings.ToLower(status) == "placed" || strings.ToLower(status) == "completed" || strings.ToLower(status) == "dispatched" || strings.ToLower(status) == "cancelled" {
		return true
	} else {
		return false
	}
}
func ProductCategoryValid(category string) bool {
	if strings.ToLower(category) == "premium" || strings.ToLower(category) == "budget" || strings.ToLower(category) == "regular" {
		return true
	} else {
		return false
	}
}
