/*
idem que pour Groq


ce cookbook est le plus adapté pour faire de la prise de décision

https://docs.typesafe.ai/cookbooks/classification_using_confidence


*/

package env

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const jevURL = "https://api.typesafe.ai/v1/systemone"

const jevModel = "jev-latest"

const jevQuestion = "classification"

type JevRequest struct {
	Model     string                 `json:"model"`
	State     string                 `json:"state"`
	Questions map[string]JevQuestion `json:"questions"`
}

type JevQuestion struct {
	Type     string            `json:"type"`
	Criteria map[string]string `json:"criteria"`
}

type jevResponse struct {
	Answers map[string]struct {
		Choice     string  `json:"choice"`
		Confidence float64 `json:"confidence"`
	} `json:"answers"`
}

type JevClient struct {
	*Client
}

func NewJevClient(apiKey string) *JevClient {
	c := NewClient()
	c.Header.Set("Authorization", "Bearer "+apiKey)
	c.Header.Set("Content-Type", "application/json")
	return &JevClient{Client: c}
}

func (j *JevClient) Classify(state string, criteria map[string]string) (string, float64, error) {
	body, err := json.Marshal(JevRequest{
		Model: jevModel,
		State: state,
		Questions: map[string]JevQuestion{
			jevQuestion: {Type: "choice", Criteria: criteria},
		},
	})
	if err != nil {
		return "", 0, err
	}

	resp, err := j.Do(context.Background(), http.MethodPost, jevURL, body)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(resp.Body)
		return "", 0, fmt.Errorf("jev: unexpected HTTP status %s: %s", resp.Status, msg)
	}

	var out jevResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", 0, err
	}
	answer, ok := out.Answers[jevQuestion]
	if !ok {
		return "", 0, fmt.Errorf("jev: missing answer")
	}
	return answer.Choice, answer.Confidence, nil
}
