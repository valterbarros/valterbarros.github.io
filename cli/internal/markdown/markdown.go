package markdown

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func MarkToHTML(mk []byte) string {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	markdownData := mk

	var buf bytes.Buffer

	// Converte o Markdown para HTML
	if err := md.Convert(markdownData, &buf); err != nil {
		panic(err)
	}

	return buf.String()
}

func RemoveMetadata(data string) string {
	reg := regexp.MustCompile(`(?mi)---[\s\S]+---`)

	return strings.TrimLeft(reg.ReplaceAllString(data, ""), " ")
}
