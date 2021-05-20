// This application is designed to be my personal blog.
// It will be as simple as possible.
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/",
		NewPage("Viktor, Go Developer").
			AddSection(
				NewSection("").
					AddText("Experienced DevOps-engineer and Golang developer. I work with technologies listed below and write my own projects."),
			).
			AddSection(
				NewSection("Skills").
					AddText(
						BulletList(
							"Go",
							"Docker",
							"Kubernetes",
							"DevOps",
							"Python",
							"Ansible",
							"Linux (Ubuntu)",
						),
					),
			).
			AddSection(
				NewSection("Projects").
					AddText(
						BulletList(
							Link("Crocodile Game Bot", "https://t.me/Crocodile_Game_Bot")+" — allows to play word-guess game in Telegram chats",
							Link("Tovarishch Bot", "https://t.me/TovarishchBot")+" — counts total words and swearings in Telegram chats",
						),
					),
			).
			AddSection(
				NewSection("Contacts").
					AddText(
						BulletList(
							Link("GitHub", "https://github.com/nuetoban"),
							Link("Telegram", "https://t.me/blackstalin"),
						),
					),
			).
			Handler(),
	)

	http.ListenAndServe(":8090", nil)
}
