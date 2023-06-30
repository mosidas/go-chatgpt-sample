package azure

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Message2 struct {
	Index   int    `json:"index"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Choice2 struct {
	Index        int        `json:"index"`
	FinishReason string     `json:"finish_reason"`
	Messages     []Message2 `json:"message"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type ChatResponse2 struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int       `json:"created"`
	Model   string    `json:"model"`
	Choices []Choice2 `json:"choices"`
}

type ChatRequest struct {
	Messages         []Message         `json:"messages,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	Stop             []string          `json:"stop,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
	DataSources      []DataSource      `json:"dataSources,omitempty"`
}

type DataSource struct {
	Type       string     `json:"type,omitempty"`
	Parameters Parameters `json:"parameters,omitempty"`
}

type Parameters struct {
	Endpoint              string        `json:"endpoint,omitempty"`
	Key                   string        `json:"key,omitempty"`
	IndexName             string        `json:"indexName,omitempty"`
	FieldsMapping         FieldsMapping `json:"fieldsMapping,omitempty"`
	InScope               bool          `json:"inScope,omitempty"`
	TopNDocuments         int           `json:"topNDocuments,omitempty"`
	QueryType             string        `json:"queryType,omitempty"`
	SemanticConfiguration string        `json:"semanticConfiguration,omitempty"`
	RoleInformation       string        `json:"roleInformation,omitempty"`
}

type FieldsMapping struct {
	ContentField  []string `json:"contentField,omitempty"`
	TitleField    string   `json:"titleField,omitempty"`
	UrlField      string   `json:"urlField,omitempty"`
	FilepathField string   `json:"filepathField,omitempty"`
}

type DataSources struct {
	DataSources []DataSource `json:"dataSources,omitempty"`
}
