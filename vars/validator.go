package variable

import (
	"net/http"
)

func EmptyFieldValidator(data string, cookie_name string, cookieJar *[]*http.Cookie) bool {
	if data == "" {
		cookie := cookieGenerator(cookie_name, ErrEmptyField.Error())
		*cookieJar = append(*cookieJar, cookie)
		return false
	}
	return true
}

func MustMatchAndNotEmptyValidator(data1 string, data2 string, cookie_name1 string, cookie_name2 string, cookieJar *[]*http.Cookie) bool {
	if EmptyFieldValidator(data1, cookie_name1, cookieJar) && EmptyFieldValidator(data2, cookie_name2, cookieJar) {
		if data1 != data2 {
			cookie1 := cookieGenerator(cookie_name1, ErrPasswordNotMatch.Error())
			cookie2 := cookieGenerator(cookie_name2, ErrPasswordNotMatch.Error())
			*cookieJar = append(*cookieJar, cookie1)
			*cookieJar = append(*cookieJar, cookie2)
			return false
		}
	}
	return true
}

func cookieGenerator(name string, value string) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}
}

// func PasswordValidator(pass string, cookie_name string, cookieJar *[]*http.Cookie) bool {
// 	if EmptyFieldValidator(pass, cookie_name, cookieJar) {
// 		re := regexp.MustCompile("(?m)^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[\!-\/\:-\@\[-\`\{-\~])(?:.{8})$`")
// 		if !re.MatchString(pass) {
// 			cookie := cookieGenerator(cookie_name, ErrPasswordNotQualified.Error())
// 			*cookieJar = append(*cookieJar, cookie)
// 			return false
// 		}
// 	}
// 	return true
// }

// func UsernameValidator(username string, cookie_name string, cookieJar *[]*http.Cookie) bool {
// 	if EmptyFieldValidator(username, cookie_name, cookieJar) {
// 		re, _ := regexp.Compile(`/^(?.*[a-z]{3})[a-z0-9]+$/i`)
// 		if !re.MatchString(username) {
// 			*cookieJar = append(*cookieJar, cookieGenerator(cookie_name, ErrPasswordNotQualified.Error()))
// 			return false
// 		}
// 	}
// }
