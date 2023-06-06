package structs

import "github.com/ayush6624/go-chatgpt"

// Downloadable input sent from frontend to pass on to chat
// gpt to parse and give us nice content to stuff in a pdf
type Downloadable struct {
	ID              int32                 `json:"id"`
	Description     string                `json:"description"`
	Title           string                `json:"title"`
	AuthorName      string                `json:"authorName"`
	Filename        string                `json:"filename"`
	ChatGPTResponse *chatgpt.ChatResponse `json:"data omitempty"`
}
