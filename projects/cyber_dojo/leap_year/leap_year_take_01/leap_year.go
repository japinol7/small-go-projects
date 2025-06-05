package main

// IsLeapYear determines if a given year is a leap year
func IsLeapYear(year int) bool {
	if year < 1 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%4 == 0 {
		return !(year%100 == 0)
	}
	return false
}
