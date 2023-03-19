package stdnum

const (
	COUNTRY_PL = "PL"

	IDENTIFIER_PESEL = "PESEL"
	IDENTIFIER_REGON = "REGON"
	IDENTIFIER_NIP   = "NIP"
)

var peselWeights = []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3, 7}

func PolandValidatePesel(value string) *ValidationResult {
	if !isOnlyDigits(value) {
		return &ValidationResult{Valid: false, Error: ErrOnlyDigitsAllowed}
	}

	if len(value) != 11 {
		return &ValidationResult{Valid: false, Error: ErrIncorrectInputLength}
	}

	val, err := toInts(value)
	if err != nil {
		return &ValidationResult{Valid: false}
	}

	sum := 0

	for i, v := range val[:10] {
		sum += v * peselWeights[i]
	}

	modulo := sum % 10

	if val[10] == 0 && modulo == 0 {
		return &ValidationResult{Valid: true}
	}

	if 10-modulo == val[10] {
		return &ValidationResult{Valid: true}
	}

	return &ValidationResult{Valid: false, Error: ErrIncorrect}
}

var regon9DigitWeights = []int{8, 9, 2, 3, 4, 5, 6, 7}
var regon14DigitWeights = []int{2, 4, 8, 5, 0, 9, 7, 3, 6, 1, 2, 4, 8}

func PolandValidateRegon(value string) *ValidationResult {
	if !isOnlyDigits(value) {
		return &ValidationResult{Valid: false, Error: ErrOnlyDigitsAllowed}
	}

	if len(value) != 9 && len(value) != 14 {
		return &ValidationResult{Valid: false, Error: ErrIncorrectInputLength}
	}

	val, err := toInts(value)
	if err != nil {
		return &ValidationResult{Valid: false, Error: ErrToDigitsConversion}
	}

	if len(value) == 14 {
		if baseResult := regon9Digits(val); !baseResult.Valid {
			return baseResult
		}

		return regon14Digits(val)
	}

	return regon9Digits(val)
}

func regon9Digits(val []int) *ValidationResult {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += val[i] * regon9DigitWeights[i]
	}

	checkDigit := sum % 11 % 10
	if checkDigit == val[8] {
		return &ValidationResult{Valid: true}
	}

	return &ValidationResult{Valid: false, Error: ErrIncorrect}
}

func regon14Digits(val []int) *ValidationResult {
	sum := 0
	for i := 0; i < 13; i++ {
		sum += val[i] * regon14DigitWeights[i]
	}

	checkDigit := sum % 11
	if checkDigit != val[13] {
		return &ValidationResult{Valid: false, Error: ErrIncorrect}
	}

	return &ValidationResult{Valid: true}
}

var nipWeights = []int{6, 5, 7, 2, 3, 4, 5, 6, 7}

func PolandValidateNip(value string) *ValidationResult {
	if !isOnlyDigits(value) {
		return &ValidationResult{Valid: false, Error: ErrOnlyDigitsAllowed}
	}

	if len(value) != 10 {
		return &ValidationResult{Valid: false, Error: ErrIncorrectInputLength}
	}

	val, err := toInts(value)
	if err != nil {
		return &ValidationResult{Valid: false, Error: ErrToDigitsConversion}
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += val[i] * nipWeights[i]
	}

	checkDigit := sum % 11 % 10
	if checkDigit != val[9] {
		return &ValidationResult{Valid: false, Error: ErrIncorrect}
	}

	return &ValidationResult{Valid: true}
}
