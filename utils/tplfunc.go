package utils

import (
	"github.com/gomarkdown/markdown"
	"html/template"
	"regexp"
	"strings"
)

// Substr returns the substr from start to length.
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// HTML2str returns escaping text convert from html.
func HTML2str(html string) string {

	re := regexp.MustCompile(`\<[\S\s]+?\>`)
	html = re.ReplaceAllStringFunc(html, strings.ToLower)

	//remove STYLE
	re = regexp.MustCompile(`\<style[\S\s]+?\</style\>`)
	html = re.ReplaceAllString(html, "")

	//remove SCRIPT
	re = regexp.MustCompile(`\<script[\S\s]+?\</script\>`)
	html = re.ReplaceAllString(html, "")

	re = regexp.MustCompile(`\<[\S\s]+?\>`)
	html = re.ReplaceAllString(html, "\n")

	re = regexp.MustCompile(`\s{2,}`)
	html = re.ReplaceAllString(html, "\n")

	return strings.TrimSpace(html)
}

// Str2html Convert string to template.HTML type.
func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

// FormatAttribute 格式化级标题
func FormatAttribute(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}
