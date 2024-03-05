package dateutils

import "errors"

func GetMaxDays(month, year int) (int, error) {
	if month < 1 {
		return 0, errors.New("month should be >= 1")
	}
	if month > 12 {
		return 0, errors.New("month should be <=12")
	}
	if year < 1900 {
		return 0, errors.New("year should be >=1900")
	}

	switch month {
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			return 29, nil
		} else {
			return 28, nil
		}
	case 4, 6, 9, 11:
		return 30, nil
	default:
		return 31, nil
	}
}