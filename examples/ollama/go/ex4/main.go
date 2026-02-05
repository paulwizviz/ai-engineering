package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main() {

	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	isTrue := false

	// User starts conversation
	req := &api.ChatRequest{
		Model: "llama3.1:8b",
		Messages: []api.Message{
			{
				Role:    "user",
				Content: "Hello",
			},
		},
		Stream: &isTrue,
	}

	// LLM response
	if err := client.Chat(context.TODO(), req, func(cr api.ChatResponse) error {
		fmt.Printf("Role: %v Content: %v\n", cr.Message.Role, cr.Message.Content)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// User prompt
	req = &api.ChatRequest{
		Model: "llama3.1:8b",
		Messages: []api.Message{
			{
				Role:    "user",
				Content: "Tell me about Italy in one sentence.",
			},
		},
		Stream: &isTrue,
	}

	// LLM response
	err = client.Chat(context.TODO(), req, func(cr api.ChatResponse) error {
		fmt.Printf("Role: %v Content: %v\n", cr.Message.Role, cr.Message.Content)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
