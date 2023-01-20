package entity

import (
	"net/http"

	utility "github.com/salamanderman234/daily-aromatic/utilities"
	variable "github.com/salamanderman234/daily-aromatic/vars"
)

func emptyFieldValidator(data string, cookieName string, cookieJar *[]*http.Cookie) bool {
	if data == "" {
		cookie := utility.CookieFactory(cookieName, variable.ErrEmptyField.Error())
		*cookieJar = append(*cookieJar, cookie)
		return false
	}
	return true
}

func mustInRangeValidator(value float64, floor float64, ceil float64, errCookieName string, cookieJar *[]*http.Cookie) bool {
	if value <= ceil && value >= floor {
		return true
	}
	cookie := utility.CookieFactory(errCookieName, variable.ErrValueNotInRange.Error())
	*cookieJar = append(*cookieJar, cookie)
	return false
}

func mustEqualOrGreaterValidator(value float64, bound float64, errCookieName string, cookieJar *[]*http.Cookie) bool {
	if value >= bound {
		return true
	}
	cookie := utility.CookieFactory(errCookieName, variable.ErrValueNotInRange.Error())
	*cookieJar = append(*cookieJar, cookie)
	return false
}

func mustMatchAndNotEmptyValidator(data1 string, data2 string, cookie_name1 string, cookie_name2 string, cookieJar *[]*http.Cookie) bool {
	if emptyFieldValidator(data1, cookie_name1, cookieJar) && emptyFieldValidator(data2, cookie_name2, cookieJar) {
		if data1 != data2 {
			cookie1 := utility.CookieFactory(cookie_name1, variable.ErrPasswordNotMatch.Error())
			cookie2 := utility.CookieFactory(cookie_name2, variable.ErrPasswordNotMatch.Error())
			*cookieJar = append(*cookieJar, cookie1)
			*cookieJar = append(*cookieJar, cookie2)
			return false
		}
	}
	return true
}
