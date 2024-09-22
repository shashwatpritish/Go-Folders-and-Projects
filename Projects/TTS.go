package main

import (
	"github.com/hegedustibor/htgo-tts"
)

func main() {
	// Create a new speech instance
	speech := htgotts.Speech{Folder: "audio", Language: "en"}

	text := "Hello I am learning GoLang."

	speech.Speak(text)

}
