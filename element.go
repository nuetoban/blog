package main

import (
	"bytes"
	"fmt"
)

func Link(text, href string) string {
	return fmt.Sprintf(`<a href="%s">%s</a>`, href, text)
}

func BulletList(elem ...string) string {
	b := bytes.NewBufferString("<ul>")

	for _, v := range elem {
		b.WriteString(fmt.Sprintf("<li>%s</li>", v))
	}

	b.WriteString("</ul>")
	return b.String()
}
