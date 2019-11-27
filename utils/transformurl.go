package utils

import "strings"

func TransformUrl(url string) string {
	rs := strings.NewReplacer("background-image", "", ": url(", "", ");", "")
	return rs.Replace(string(url))
}
