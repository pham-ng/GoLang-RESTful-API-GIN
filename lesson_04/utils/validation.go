package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// ValidationRequired trả lỗi nếu value rỗng.
func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

// ValidationStringLength trả lỗi nếu số ký tự của value nằm ngoài khoảng [min, max].
// Dùng utf8.RuneCountInString để đếm đúng số ký tự (len() đếm byte,
// nên chữ có dấu như "é" sẽ bị tính là 2).
func ValidationStringLength(fieldName, value string, min, max int) error {
	l := utf8.RuneCountInString(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be between %d and %d characters", fieldName, min, max)
	}
	return nil
}

// ValidationRegex trả lỗi nếu value không khớp với mẫu re.
func ValidationRegex(fieldName, value string, re *regexp.Regexp, errorMessage string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s %s", fieldName, errorMessage)
	}
	return nil
}

func ValidationPositiveInt(fieldName, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number", fieldName)
	}
	if v <= 0 {
		return 0, fmt.Errorf("%s must be positive", fieldName)
	}
	return v, nil

}

func ValidationUuid(fieldName, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", fieldName)
	}
	return uid, nil
}

func keys(m map[string]bool) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}

func ValidationInList(fieldName, value string, allowed map[string]bool) error {
	if !allowed[value] {
		return fmt.Errorf("%s must be one of: %v", fieldName, keys(allowed))
	}
	return nil
}

func HandleValidationError(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			switch e.Tag() {
			case "required":
				errors[e.Field()] = e.Field() + " is required"
			case "gt":
				errors[e.Field()] = e.Field() + " must be greater than " + e.Param()
			case "gte":
				errors[e.Field()] = e.Field() + " must be greater than or equal to " + e.Param()
			case "lt":
				errors[e.Field()] = e.Field() + " must be less than " + e.Param()
			case "lte":
				errors[e.Field()] = e.Field() + " must be less than or equal to " + e.Param()
			case "min":
				errors[e.Field()] = e.Field() + " must be at least " + e.Param()
			case "max":
				errors[e.Field()] = e.Field() + " must be at most " + e.Param()
			case "len":
				errors[e.Field()] = e.Field() + " must have length " + e.Param()
			case "oneof":
				errors[e.Field()] = e.Field() + " must be one of: " + e.Param()
			case "uuid":
				errors[e.Field()] = e.Field() + " must be valid uuid"
			case "email":
				errors[e.Field()] = e.Field() + " must be a valid email"
			case "url":
				errors[e.Field()] = e.Field() + " must be a valid URL"
			case "numeric":
				errors[e.Field()] = e.Field() + " must be a number"
			case "alpha":
				errors[e.Field()] = e.Field() + " must contain only letters"
			case "alphanum":
				errors[e.Field()] = e.Field() + " must contain only letters and numbers"
			default:
				errors[e.Field()] = e.Field() + " is invalid"
			}
		}
		return gin.H{
			"error": errors,
		}

	}
	return gin.H{
		"error": "Request invalid" + err.Error(),
	}
}
