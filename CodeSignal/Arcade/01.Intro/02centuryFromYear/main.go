package centuryfromyear

func centuryFromYear(year int) int {
	carry := year % 100
	if carry > 0 {
		return year/100 + 1
	} else {
		return year / 100
	}
}
