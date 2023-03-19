package stdnum

import (
	"strconv"
	"strings"
	"unicode"
)

func isOnlyDigits(value string) bool {
	for _, v := range value {
		if !unicode.IsDigit(v) {
			return false
		}
	}

	return true
}

func toInts(value string) ([]int, error) {
	ints := make([]int, 0, len(value))

	for _, v := range strings.Split(value, "") {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}

	return ints, nil
}
