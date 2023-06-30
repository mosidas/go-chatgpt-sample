package azure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatClientAzure struct {
	apiKey      string
	apiBase     string
	apiType     string
	apiVersion  string
	engine      string
	messages    []Message
	systemPrmpt string
	totalTokens int
	chatUrl     string
}

// initChatGPT initializes a new ChatGPT struct
func NewChatClientAzure(apiKey string, apiBase string, apiVersion string, engine string, systemPrompt string) *ChatClientAzure {
	return &ChatClientAzure{
		apiKey:      apiKey,
		apiBase:     apiBase,
		apiType:     "azure",
		apiVersion:  apiVersion,
		engine:      engine,
		messages:    make([]Message, 0),
		systemPrmpt: systemPrompt,
		chatUrl:     "",
	}
}

func (c *ChatClientAzure) setSystemPrompt() {
	if len(c.messages) == 0 && c.systemPrmpt != "" {
		c.messages = append(c.messages, Message{Role: "system", Content: c.systemPrmpt})
	}
}

func (c *ChatClientAzure) setRequest(message string, temperature float64, stream bool) (*http.Request, error) {
	c.chatUrl = c.apiBase + "openai/deployments/" + c.engine + "/chat/completions?api-version=" + c.apiVersion
	c.messages = append(c.messages, Message{Role: "user", Content: message})

	data := ChatRequest{
		Messages:    c.messages,
		Temperature: temperature,
		Stream:      stream,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.chatUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-key", c.apiKey)

	return request, nil
}

// Chat sends a prompt to the OpenAI API and returns the response
func (c *ChatClientAzure) Chat(message string, temperature float64) (string, error) {
	c.setSystemPrompt()

	request, err := c.setRequest(message, temperature, false)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return "", err
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
		return "", err
	}

	// Parse the response body
	var response ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	c.totalTokens = response.Usage.TotalTokens

	// Append the response to the messages
	c.messages = append(c.messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Message.Content,
	})

	return c.GetNewMessage(), err
}

func (c *ChatClientAzure) ChatWithData(message string, temperature float64) (string, error) {
	c.setSystemPrompt()

	request, err := c.setRequestOnYourData(message, temperature, false)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return "", err
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
		return "", err
	}

	fmt.Println(string(body))

	// Parse the response body
	var response ChatResponse2
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	c.totalTokens = 0

	// print response
	fmt.Println(response)

	// Append the response to the messages
	c.messages = append(c.messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages[1].Content,
	})

	return c.GetNewMessage(), err
}

func (c *ChatClientAzure) setRequestOnYourData(message string, temperature float64, stream bool) (*http.Request, error) {
	c.chatUrl = c.apiBase + "openai/deployments/" + c.engine + "/extensions/chat/completions?api-version=" + c.apiVersion
	c.messages = append(c.messages, Message{Role: "user", Content: message})

	data := ChatRequest{
		Messages:    c.messages,
		Temperature: temperature,
		Stream:      stream,
		DataSources: []DataSource{
			{
				Type: "AzureCognitiveSearch",
				Parameters: Parameters{
					Endpoint:  "https://exmaple.net",
					Key:       "secret-key",
					IndexName: "my-index",
					FieldsMapping: FieldsMapping{
						ContentField:  []string{"my-content"},
						TitleField:    "title",
						UrlField:      "url",
						FilepathField: "file-path",
					},
					InScope:               true,
					TopNDocuments:         10,
					QueryType:             "simple",
					SemanticConfiguration: "",
					RoleInformation:       "hoge",
				},
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.chatUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api-key", c.apiKey)

	return request, nil
}

// GetMessages returns the messages stored in the ChatGPT struct
func (c *ChatClientAzure) GetMessages() []Message {
	return c.messages
}

// GetNewMesasge
func (c *ChatClientAzure) GetNewMessage() string {
	return c.messages[len(c.messages)-1].Content
}

func (c *ChatClientAzure) GetTotalTokens() int {
	return c.totalTokens
}
