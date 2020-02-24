package visitor

import (
	"regexp"

	"github.com/LuLStackCoder/go-patterns/pkg/html_doc"
	"github.com/LuLStackCoder/go-patterns/pkg/models"
	"github.com/LuLStackCoder/go-patterns/pkg/tex_doc"
)

type Visitor interface {
	ParseHtml(html html_doc.HtmlDoc) string
	ParseTex(tex tex_doc.TexDoc) string
}

type visitor struct{}

func (v *visitor) ParseHtml(html html_doc.HtmlDoc) (result string) {
	var htmlString = html.GetText()
	r := regexp.MustCompile(models.HtmlTags)
	result = r.ReplaceAllString(htmlString, "")
	return
}

func (v *visitor) ParseTex(tex tex_doc.TexDoc) (result string) {
	var htmlString = tex.GetText()
	r := regexp.MustCompile(models.TexMarkUp)
	result = r.ReplaceAllString(htmlString, "")
	return
}

func NewVisitor() Visitor {
	return &visitor{}
}
