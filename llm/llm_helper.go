package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ChatCompletionResponse struct {
	ID                *string   `json:"id"`
	Object            *string   `json:"object"`
	Created           *int64    `json:"created"`
	Model             *string   `json:"model"`
	Choices           []*Choice `json:"choices"`
	Usage             *Usage    `json:"usage"`
	SystemFingerprint *string   `json:"system_fingerprint"`
	XGroq             *XGroq    `json:"x_groq"`
}

type Choice struct {
	Index        *int     `json:"index"`
	Message      *Message `json:"message"`
	Logprobs     *string  `json:"logprobs"` // Nullable, assuming string or null
	FinishReason *string  `json:"finish_reason"`
}

type Message struct {
	Role    *string `json:"role"`
	Content *string `json:"content"`
}

type Usage struct {
	QueueTime        *float64 `json:"queue_time"`
	PromptTokens     *int     `json:"prompt_tokens"`
	PromptTime       *float64 `json:"prompt_time"`
	CompletionTokens *int     `json:"completion_tokens"`
	CompletionTime   *float64 `json:"completion_time"`
	TotalTokens      *int     `json:"total_tokens"`
	TotalTime        *float64 `json:"total_time"`
}

type XGroq struct {
	ID *string `json:"id"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

const COMPLETIONS_BASE_URL = "https://api.groq.com/openai/v1/chat/completions"

func GetCompletionsResponse(prompt string) string {
	chatRequest := ChatRequest{
		Model: "llama-3.3-70b-versatile",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	reqB, _ := json.Marshal(chatRequest)
	reqBody := bytes.NewBuffer(reqB)

	req, err := http.NewRequest("POST", COMPLETIONS_BASE_URL, reqBody)
	if err != nil {
		fmt.Println("Error while creating completions API request")
		return ""
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	var completionsResponse *ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&completionsResponse)
	if err != nil {
		fmt.Println("Error reading completions response:", err)
		return ""
	}

	if len(completionsResponse.Choices) > 0 {
		if completionsResponse.Choices[0].Message != nil {
			return *completionsResponse.Choices[0].Message.Content
		}
	}

	return ""
}
