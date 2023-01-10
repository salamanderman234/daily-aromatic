package variable

import (
	"net/http"
)

func EmptyFieldValidator(data string, cookie_name string, cookieJar *[]*http.Cookie) bool {
	if data == "" {
		cookie := &http.Cookie{
			Name:  cookie_name,
			Value: ErrEmptyField.Error(),
			Path:  "/",
		}
		*cookieJar = append(*cookieJar, cookie)
		return false
	}
	return true
}

func MustMatchAndNotEmptyValidator(data1 string, data2 string, cookie_name1 string, cookie_name2 string, cookieJar *[]*http.Cookie) bool {
	if EmptyFieldValidator(data1, cookie_name1, cookieJar) && EmptyFieldValidator(data2, cookie_name2, cookieJar) {
		if data1 != data2 {
			cookie1 := &http.Cookie{
				Name:  cookie_name1,
				Value: ErrPasswordNotMatch.Error(),
				Path:  "/",
			}
			cookie2 := &http.Cookie{
				Name:  cookie_name2,
				Value: ErrPasswordNotMatch.Error(),
				Path:  "/",
			}
			*cookieJar = append(*cookieJar, cookie1)
			*cookieJar = append(*cookieJar, cookie2)
			return false
		}
	}
	return false
}
