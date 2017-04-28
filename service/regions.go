package service

import (
	"golang.org/x/text/language"
)

// Find returns Region and Confidence for Language
func Find(lang string) (language.Region, language.Confidence) {
	en := language.Make(lang)
	return en.Region()
}
