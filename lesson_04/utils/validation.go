package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
			case "gt":
				errors[e.Field()] = e.Field() + " must greater  than minimum"
			case "lt":
				errors[e.Field()] = e.Field() + " must less than maximum"
			case "gte":
				errors[e.Field()] = e.Field() + " must equal or greater than minimum"
			case "lte":
				errors[e.Field()] = e.Field() + " must equal less than maximum"
			case "uuid":
				errors[e.Field()] = e.Field() + " must be valid uuid"
			case "slug":
				errors[e.Field()] = e.Field() + " must be characters, hyphens and dots"
			case "min":
				errors[e.Field()] = fmt.Sprintf("%s must longer than %s", e.Field(), e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("%s must shortter than %s", e.Field(), e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[e.Field()] = fmt.Sprintf("%s must be one of %s", e.Field(), allowedValues)
			case "required":
				errors[e.Field()] = e.Field() + " must required"
			case "search":
				errors[e.Field()] = e.Field() + " must be lower case, uppercase and spaces"
			case "email":
				errors[e.Field()] = e.Field() + " must follow email format"
			case "datetime":
				errors[e.Field()] = e.Field() + "must follow datetime format"
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

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("Failed to get validator engine")
	}
	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	if err := v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	}); err != nil {
		return err
	}

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})
	return nil
}
