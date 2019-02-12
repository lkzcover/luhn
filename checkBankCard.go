package luhn

import (
	"errors"
)

//CheckBankCardNumber function receives the Bank card number and checks it for validity
func CheckBankCardNumber(number string) (bool, error) {
	if len(number) != 16 {
		return false, errors.New("invalid len card number")
	}

	return CheckLuhn(number)
}
