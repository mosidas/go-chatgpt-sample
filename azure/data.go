package azure

type Message struct {
	Index   int    `json:"index,omitempty"`
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type Choice struct {
	Index        int       `json:"index,omitempty"`
	FinishReason string    `json:"finish_reason,omitempty"`
	Message      Message   `json:"message,omitempty"`
	Messages     []Message `json:"messages,omitempty"`
}

type Usage struct {
	CompletionTokens int `json:"completion_tokens,omitempty"`
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

type ChatResponse struct {
	ID      string   `json:"id,omitempty"`
	Object  string   `json:"object,omitempty"`
	Created int      `json:"created,omitempty"`
	Model   string   `json:"model,omitempty"`
	Choices []Choice `json:"choices,omitempty"`
	Usage   Usage    `json:"usage,omitempty"`
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
