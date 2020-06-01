package main

// T: O(1) for 32 times
// M: O(1)
// -- start --

func divide(dividend int, divisor int) int {

	negative := false

	if dividend < 0 {
		dividend = -dividend
		negative = !negative
	}

	if divisor < 0 {
		divisor = -divisor
		negative = !negative
	}

	divisor <<= 31
	quotient := 0

	for i := 0; i < 32; i++ {
		quotient <<= 1

		if dividend >= divisor {
			dividend -= divisor
			quotient++
		}

		divisor >>= 1
	}

	if negative {
		quotient = -quotient
	}

	if quotient > 1<<31-1 {
		return 1<<31 - 1
	}

	if quotient < -1<<31 {
		return -1 << 31
	}

	return quotient
}

// -- end --
