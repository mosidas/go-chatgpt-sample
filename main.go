package main

import (
	"bufio"

	client "chatgpt-sample/azure"
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the API key from environment variables
	// apiKey := os.Getenv("OPENAI_API_KEY")
	azureApiKey := os.Getenv("AZURE_OPENAI_API_KEY")
	apiBase := os.Getenv("AZURE_OPENAI_API_BASE")
	model := os.Getenv("AZURE_OPENAI_MODEL")
	// model := "gpt-3.5-turbo-0613"

	// Initialize a new ChatGPT struct
	// chat := client.NewChatClient(apiKey, model, SystemPrompt())
	chat := client.NewChatClientAzure(azureApiKey, apiBase, "2023-03-15-preview", model, SystemPrompt())

	// Create a new reader to read from stdin
	reader := bufio.NewReader(os.Stdin)

	// Loop forever
	for {
		// Read a line from stdin
		fmt.Print("Me: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// If message is empty, break
		if len(line) <= 2 {
			fmt.Println("sys:", "Good bye.")
			break
		}

		// Send the line to the OpenAI API
		answer, err := chat.Chat(line, 1)
		if err != nil {
			log.Fatal(err)
		}

		// Print the response
		fmt.Println("AI:", answer)

		// if total tokens over 4000, break
		if chat.GetTotalTokens() > 4000 {
			fmt.Println("sys:", "Total token over 4000. Good bye.")
			break
		}

	}
}
