package config

import "fmt"

const (
	AppName = "Daily Aromatic"
)

func FromViews(path string) string {
	return fmt.Sprintf("./views%s", path)
}
