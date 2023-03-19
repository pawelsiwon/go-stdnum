package stdnum

import "errors"

var (
	ErrIncorrect            = errors.New("checksum does not match")
	ErrIncorrectInputLength = errors.New("incorrect input length")
	ErrOnlyDigitsAllowed    = errors.New("only digits allowed")
	ErrToDigitsConversion   = errors.New("error when extracting digits")
)
