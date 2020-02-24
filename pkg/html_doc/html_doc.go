package html_doc

type visitor interface {
	ParseHtml(html HtmlDoc) string
}

type HtmlDoc interface {
	Accept(v visitor) string
	GetText() string
}

type htmlDoc struct {
	text string
}

func (h *htmlDoc) Accept(v visitor) string {
	return v.ParseHtml(h)
}

func (h *htmlDoc) GetText() string {
	return h.text
}

func NewHtmlDoc(text string) HtmlDoc {
	return &htmlDoc{
		text: text,
	}
}
