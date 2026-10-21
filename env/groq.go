/*
ce code est le client Groq qui hérite de Client, on standardise les headers et le endpoint pour l'API Groq, et on fournit une méthode pour faire des requetes.

https://console.groq.com/docs/api-reference#chat 

btw, peut être que ça peut être cool de faire un layer LLM général du style LiteLLM plutôt, pour le moment on fait comme en TD

*/

package env

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const groqURL = "https://api.groq.com/openai/v1/chat/completions"

const groqModel = "openai/gpt-oss-20b"

type GroqRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type GroqClient struct {
	*Client
}


func NewGroqClient(apiKey string) *GroqClient {
	c := NewClient()
	c.Header.Set("Authorization", "Bearer "+apiKey)
	c.Header.Set("Content-Type", "application/json")
	return &GroqClient{Client: c}
}


func (g *GroqClient) ChatCompletion(prompt string) (string, error) {
	body, err := json.Marshal(GroqRequest{
		Model:    groqModel,
		Messages: []Message{{Role: "user", Content: prompt}},
	})
	if err != nil {
		return "", err
	}

	resp, err := g.Do(context.Background(), http.MethodPost, groqURL, body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("groq: unexpected HTTP status %s", resp.Status)
	}

	var out groqResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if len(out.Choices) == 0 {
		return "", fmt.Errorf("groq: empty response")
	}
	return out.Choices[0].Message.Content, nil
}
