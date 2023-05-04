package simple

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Content     string
}

func ParsePost(str string) Post {
	re := regexp.MustCompile(`\[header\]\s*((?:.|\n)*?)\s*\[/header\]`)
	header := re.FindStringSubmatch(str)[1]

	content := re.ReplaceAllString(str, "")
	content = ParseMarkdown(content)

	re = regexp.MustCompile(`title:\s*(.*)`)
	title := re.FindStringSubmatch(header)[1]

	re = regexp.MustCompile(`description:\s*(.*)`)
	description := re.FindStringSubmatch(header)[1]

	re = regexp.MustCompile(`tags:\s*(.*)`)
	tags := re.FindStringSubmatch(header)[1]
	parsedTags := strings.Split(tags, ",")

	post := Post{
		Title:       title,
		Description: description,
		Tags:        parsedTags,
		Content:     content,
	}

	return post
}

func ParseMarkdown(md string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	data := string(markdown.Render(doc, renderer))

	fmt.Println(data)
	return data
}
