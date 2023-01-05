package utility

import (
	"fmt"
	"net/http"
)

func SameErrorCookieGen(names []string, prefix string, value string) []*http.Cookie {
	cookieList := []*http.Cookie{}
	for _, name := range names {
		cookieList = append(cookieList, &http.Cookie{
			Name:  fmt.Sprintf("%s_%s", name, prefix),
			Value: value,
			Path:  "/",
		})
	}
	return cookieList
}
