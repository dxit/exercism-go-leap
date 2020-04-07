package leap

// IsLeapYear returns true if the input year is a leap one
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
