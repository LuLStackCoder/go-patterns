package tex_doc

type visitor interface {
	ParseTex(tex TexDoc) string
}

type TexDoc interface {
	Accept(v visitor) string
	GetText() string
}

type texDoc struct {
	text string
}

func (t *texDoc) Accept(v visitor) string {
	return v.ParseTex(t)
}

func (t *texDoc) GetText() string {
	return t.text
}

func NewTexDoc(text string) TexDoc {
	return &texDoc{
		text: text,
	}
}
