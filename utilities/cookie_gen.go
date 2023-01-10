package utility

import (
	"fmt"
	"net/http"
)

func SameCookieGen(names []string, value string) []*http.Cookie {
	cookieList := []*http.Cookie{}
	for _, name := range names {
		cookieList = append(cookieList, &http.Cookie{
			Name:  fmt.Sprintf("%s", name),
			Value: value,
			Path:  "/",
		})
	}
	return cookieList
}
