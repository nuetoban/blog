package main

import (
	"bytes"
	"net/http"
)

type Section struct {
	Header  string
	Content *bytes.Buffer
}

func NewSection(header string) *Section {
	s := &Section{}
	s.Content = bytes.NewBufferString("")
	s.Header = header
	return s
}

func (s *Section) AddText(text string) *Section {
	s.Content.WriteString(text)
	s.Content.WriteString("<br>")
	return s
}

type Page struct {
	Header   string
	Sections []Section
}

func NewPage(header string) *Page {
	p := &Page{}
	p.Header = header
	return p
}

func (p *Page) AddSection(s *Section) *Page {
	p.Sections = append(p.Sections, *s)
	return p
}

func (p *Page) Render() *bytes.Buffer {
	b := bytes.NewBufferString("<html>\n")

	b.WriteString("<head>\n")
	b.WriteString("<title>")
	b.WriteString(p.Header)
	b.WriteString("</title>\n")
	b.WriteString(`<meta charset="utf-8">`)
	b.WriteString(`<link rel="preconnect" href="https://fonts.gstatic.com">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link href="https://fonts.googleapis.com/css2?family=Roboto+Mono&display=swap" rel="stylesheet">
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-+0n0xVW2eSR5OomGNYDnhzAbDsOXxcvSN1TPprVMTNDbiYZCxYbOOl7+AMvyTG2x" crossorigin="anonymous">
	`)
	b.WriteString(`<style>
	* {
		background-color: #eee8d5;
		font-family: 'Roboto Mono', monospace;
		color: #586e75;
	}

	body {
		background-color: #eee8d5;
		padding-top: 3em;
		padding-bottom: 3em;
	}

	ul {
		position: relative;
		list-style: none;
		margin-left: 0;
		padding-left: 2em;
	}

	ul li:before {
		content: "+";
		position: absolute;
		left: 1em;
	}

	a {
		color: #d33682;
	}

	a:hover {
		color: #268bd2;
	}

	h1 {
		color: #859900;
		font-size: 16;
	}

	h2 {
		color: #859900;
		font-size: 16;
	}
	</style>`)
	b.WriteString("\n</head>\n")

	b.WriteString("<body>\n")
	b.WriteString(`<div class="container">`)
	// b.WriteString(`<div class="col-2">`)
	// b.WriteString(`</div>`)
	b.WriteString(`<div class="col-sm-7">`)
	b.WriteString("\n<h1># ")
	b.WriteString(p.Header)
	b.WriteString("</h1>\n")

	for _, s := range p.Sections {
		b.WriteString("\n<p>\n")

		// Header
		if s.Header != "" {
			b.WriteString("<h2>## ")
			b.WriteString(s.Header)
			b.WriteString("</h2>\n")
		}

		// Content
		b.WriteString(s.Content.String())

		b.WriteString("\n</p>\n")
		if s.Header == "" {
			b.WriteString("<br>\n")
		}
	}

	b.WriteString(`
	------------------------
	<br>
	2021	
	<br>
	You can find the source code at ` + Link("github.com/nuetoban/blog", "https://github.com/nuetoban/blog") + `
	`)

	b.WriteString(`</div>`)
	b.WriteString(`</div>`)
	b.WriteString("\n</body>")
	b.WriteString("</html>\n")

	return b
}

func (p *Page) Handler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(p.Render().Bytes())
	}
}
