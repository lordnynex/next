package template

import (
	"bytes"
	"html/template"
	"io"
)

type HTML struct {
	*template.Template

	TemplateRoot string
	TemplateExt  string
}

func NewHTML(templateRoot, templateExt string, useCache bool) *HTML {
	html := &HTML{
		TemplateRoot: templateRoot,
		TemplateExt:  templateExt,
	}
	if useCache {
		html.Template = html.loadTemplate()
	}
	return html
}

func (h *HTML) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	template := h.Template
	if template == nil {
		template = h.loadTemplate()
	}
	return template.ExecuteTemplate(wr, name, data)
}

func (h *HTML) ExecuteTemplateToBytes(name string, data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	if err := h.ExecuteTemplate(&buffer, name, data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (h *HTML) ExecuteTemplateToString(name string, data interface{}) (string, error) {
	var buffer bytes.Buffer
	if err := h.ExecuteTemplate(&buffer, name, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (h *HTML) loadTemplate() *template.Template {
	return template.Must(
		ParseWalk(template.New(""), h.TemplateRoot, h.TemplateExt),
	)
}
