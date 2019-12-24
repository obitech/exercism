// Package leap provides a function to calculate if a given year is a leap year.
package leap

// IsLeapYear returns true if a year is a leap year.
func IsLeapYear(y int) bool {
	return (y % 4 == 0) && (y % 100 != 0 || y % 400 == 0)
}
