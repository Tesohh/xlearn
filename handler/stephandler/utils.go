package stephandler

import (
	"slices"
)

func localize(localee map[string]string, lang string) string {
	// first try: exact language
	localized, ok := localee[lang]
	if ok {
		return localized
	}

	// second try: org's favorite language
	// TODO: add org favorite language

	// thirt try: en_us
	localized, ok = localee["en_us"]
	if ok {
		return localized
	}

	// fourth try: en_uk
	localized, ok = localee["en_uk"]
	if ok {
		return localized
	}

	// at this point, whats even happening? error out.

	return "no valid language found for this item"
}

func getLangs(localee map[string]string) []string {
	s := make([]string, 0)

	for k := range localee {
		s = append(s, k)
	}

	return s
}

func mergeLangs(langss ...[]string) []string {
	s := make([]string, 0)
	for _, langs := range langss {
		for _, lang := range langs {
			if !slices.Contains(s, lang) {
				s = append(s, lang)
			}
		}
	}

	return s
}
