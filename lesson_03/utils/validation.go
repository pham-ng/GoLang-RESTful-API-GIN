package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"

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
