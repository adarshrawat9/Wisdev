package utils

import "net/url"

func IsValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return false
	}

	return parsedURL.Scheme != "" &&
		parsedURL.Host != ""
}