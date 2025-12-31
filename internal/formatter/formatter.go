package formatter

import (
	"strings"
)

type Answer struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FormData struct {
	Answers []Answer `json:"answers"`
}

func FormatForm(form *FormData) string {
	if len(form.Answers) == 0 {
		return ""
	}

	var lines []string
	for _, a := range form.Answers {
		escaped := escapeMarkdown(a.Value)
		lines = append(lines, "*"+a.Key+":* "+escaped)
	}
	return strings.Join(lines, "\n")
}

func escapeMarkdown(s string) string {
	replacer := strings.NewReplacer(
		"*", "\\*",
		"_", "\\_",
	)
	return replacer.Replace(s)
}
