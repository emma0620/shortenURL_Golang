package utils

import "net/url"

// IsValidURL 驗證是否為合法的 URL
func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}
