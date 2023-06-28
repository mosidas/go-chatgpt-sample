package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ChatClient struct {
	apiKey      string
	model       string
	messages    []Message
	systemPrmpt string
	totalTokens int
}

// initChatGPT initializes a new ChatGPT struct
func NewChatClient(apiKey string, model string, systemPrompt string) *ChatClient {
	return &ChatClient{
		apiKey:      apiKey,
		model:       model,
		messages:    make([]Message, 0),
		systemPrmpt: systemPrompt,
	}
}

func (c *ChatClient) setSystemPrompt() {
	// if initilal send, send system prompt
	if len(c.messages) == 0 && c.systemPrmpt != "" {
		c.messages = append(c.messages, Message{Role: "system", Content: c.systemPrmpt})
	}
}

func (c *ChatClient) setChatRequest(url string, body ChatRequest) (*http.Request, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.apiKey)

	return request, nil
}

func (c *ChatClient) sendChatRequest(request *http.Request) (*ChatResponse, error) {
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response body
	var response ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Chat sends a prompt to the OpenAI API and returns the response
func (c *ChatClient) Chat(message string, temperature float64) (string, error) {
	c.setSystemPrompt()

	c.messages = append(c.messages, Message{Role: "user", Content: message})
	body := ChatRequest{
		Model:       c.model,
		Messages:    c.messages,
		Temperature: temperature,
	}

	request, err := c.setChatRequest("https://api.openai.com/v1/chat/completions", body)
	if err != nil {
		return "", err
	}

	response, err := c.sendChatRequest(request)
	if err != nil {
		return "", err
	}

	c.totalTokens = response.Usages.TotalTokens

	// Append the response to the messages
	c.messages = append(c.messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages.Content,
	})

	return c.GetNewMessage(), err
}

// GetMessages returns the messages stored in the ChatGPT struct
func (c *ChatClient) GetMessages() []Message {
	return c.messages
}

// GetNewMesasge
func (c *ChatClient) GetNewMessage() string {
	return c.messages[len(c.messages)-1].Content
}

func (c *ChatClient) GetTotalTokens() int {
	return c.totalTokens
}
