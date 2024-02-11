package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Alvphil/LN_to_speech/internal/websites"
)

type siteUrl struct {
	domain   string
	callback func(url string) (string, error)
}

func getCommands() map[string]siteUrl {
	return map[string]siteUrl{
		"wanderingInn": {
			domain:   "wanderinginn.com",
			callback: websites.GetChapterWI,
		},
	}
}

func cleanDomain(str string) string {
	lowered := strings.ToLower(str)
	var re = regexp.MustCompile(`(?m)https?:\/\/([^\/]+)`)

	match := re.FindStringSubmatch(lowered)
	if match != nil && len(match) > 1 {
		// match[0] is the entire match, match[1] is the first submatch
		fmt.Println(match[1])
		return match[1]
	}
	return ""
}
