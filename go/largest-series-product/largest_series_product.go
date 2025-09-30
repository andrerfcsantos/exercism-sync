package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	return LargestSeriesProduct_EndsOfSpan(digits, span)
}

func LargestSeriesProduct_MultiplyAllSpan(digits string, span int) (int64, error) {
	if span < 0 || span > len(digits) {
		return 0, errors.New("span must be greater than 0 and less than the size of the string")
	}

	if span == 0 {
		return 1, nil
	}

	var largest int64
	for i := 0; i <= len(digits)-span; i++ {

		product := int64(1)
		for j := 0; j < span; j++ {
			if !isDigit(digits[i+j]) {
				return 0, errors.New("string must contain only digits")
			}
			product *= toInt64(digits[i+j])
		}
		if product > largest {
			largest = product
		}
	}
	return largest, nil
}

func LargestSeriesProduct_EndsOfSpan(digits string, span int) (int64, error) {
	if span < 0 || span > len(digits) {
		return 0, errors.New("span must be greater than 0 and less than the size of the string")
	}

	if span == 0 {
		return 1, nil
	}

	product := int64(1)

	// compute the product of the first span of digits
	for j := 0; j < span; j++ {
		if !isDigit(digits[j]) {
			return 0, errors.New("string must contain only digits")
		}
		product *= int64(digits[j] - '0')
	}

	largest := product

	// process the rest of the spans
	for i := 1; i <= len(digits)-span; i++ {
		if !isDigit(digits[i-1]) || !isDigit(digits[i+span-1]) {
			return 0, errors.New("string must contain only digits")
		}

		if product != 0 {
			// fast path: compute the product of the current span by diving its product
			// by the last digit of the previous span and multiplying it by the last digit
			// of the current span
			first, last := toInt64(digits[i-1]), toInt64(digits[i+span-1])
			product /= first
			product *= last
		} else {
			// if product of the previous span 0, we can't use it to compute the product
			// of the current span using the fast path - we have to fallback to compute
			// the product by multiplying all the digits in the span
			product = 1
			for j := 0; j < span; j++ {
				product *= toInt64(digits[i+j])
			}
		}

		if product > largest {
			largest = product
		}
	}

	return largest, nil
}

func toInt64(char byte) int64 {
	return int64(char - '0')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
