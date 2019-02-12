package luhn

import (
	"strconv"
)

//CheckLuhn implementation of the Luhn algorithm
func CheckLuhn(numberStr string) (bool, error) {
	var sum uint64
	if len(numberStr)%2 != 0 {
		num, err := strconv.ParseUint(numberStr[:1], 10, 64)
		if err != nil {
			return false, err
		}
		sum = num
		numberStr = numberStr[1:]
	}
	for i, n := range numberStr {
		if (i+1)%2 != 0 {
			num, err := strconv.ParseUint(string(n), 10, 64)
			if err != nil {
				return false, err
			}
			num = num * 2
			if num > 9 {
				num = num - 9
			}
			sum = sum + num
			continue
		}
		num, err := strconv.ParseUint(string(n), 10, 64)
		if err != nil {
			return false, err
		}
		sum = sum + num
	}
	if sum%10 != 0 {
		return false, nil
	}
	return true, nil
}
