package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Temperature float64   `json:"temperature"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream"`
}

type IntentResponse struct {
	Intent                    string `json:"intent"`
	RequiresAuthoritativeData bool   `json:"requires_authoritative_data"`
	Ambiguity                 string `json:"ambiguity"`
}

func main() {
	reqBody := ChatRequest{
		Model:       "llama3.1:8b",
		Temperature: 0.0,
		Stream:      false,
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are an intent classification engine for an AI agent. You do not answer user questions.",
			},
			{
				Role:    "system",
				Content: "Choose exactly one intent from the allowed list. If none apply, choose \"unsupported\".",
			},
			{
				Role: "system",
				Content: `Intent taxonomy:

historical_lookup:
- Meaning: Retrieve specific past lottery draw results.
- Includes: results for a specific date.
- Excludes: statistics, trends, predictions.

statistical_aggregation:
- Meaning: Compute counts, frequencies, maxima, minima, or distributions across historical lottery data.
- Includes: most frequent number, highest occurrence.
- Excludes: future predictions, single draw lookups.

filtered_analysis:
- Meaning: Statistical aggregation with explicit filtering conditions.
- Includes: most common number on even draw dates.
- Excludes: unfiltered aggregation, predictions.

metadata_query:
- Meaning: Questions about dataset coverage or structure.
- Includes: number of draws stored, date range.
- Excludes: draw results, statistics.

prediction_request:
- Meaning: Requests to forecast or guess future lottery outcomes.
- Always unsupported.

general_explanation:
- Meaning: Conceptual explanations not requiring data access.
- Includes: how lotteries work.

unsupported:
- Meaning: Requests outside system capability.
`,
			},
			{
				Role: "system",
				Content: `Return ONLY valid JSON in this schema:
{
  "intent": "string",
  "requires_authoritative_data": boolean,
  "ambiguity": "low | medium | high"
}`,
			},
			{
				Role:    "user",
				Content: "What number appears most often in UK Lotto draws?",
			},
		},
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(
		"http://localhost:11434/api/chat",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var response struct {
		Message Message `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	var intent IntentResponse
	if err := json.NewDecoder(bytes.NewReader([]byte(response.Message.Content))).Decode(&intent); err != nil {
		log.Fatal(err)
	}
	fmt.Println(intent)
}
