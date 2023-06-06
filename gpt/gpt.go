package gpt

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
	"jay-lane.com/hackathon/structs"
)

var Client chatgpt.Client

func GetClient() (chatgpt.Client, error) {
	key := os.Getenv("OPENAI_KEY")
	if key == "" {
		return Client, errors.New("invalid openai key")
	}

	Client, err := chatgpt.NewClient(key)

	if err != nil {
		return *Client, err
	}

	return *Client, nil
}

func SendDownloadableToChatGPT(d structs.Downloadable) (err error) {
	c, err := GetClient()

	if err != nil {
		return err
	}

	ctx := context.Background()

	res, err := c.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: "Hey, Explain GoLang to me in 2 sentences.",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	a, _ := json.MarshalIndent(res, "", "  ")

	log.Println(string(a))

	return nil
}
