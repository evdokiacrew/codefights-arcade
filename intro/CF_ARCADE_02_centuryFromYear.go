/*
intro 02/60 centuryFromYear

Given a year, return the century it is in.
The first century spans from the year 1 up to and including the year 100,
the second - from the year 101 up to and including the year 200, etc.

*/

func centuryFromYear(year int) int {
	if year%100 != 0 {
		return year/100 + 1
	} else {
		return year / 100
	}
}
