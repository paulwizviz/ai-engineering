package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	DefaultLocalBaseURL = "http://localhost:11434"
)

type RequestGenerate struct {
	Model     string      `json:"model"`
	Prompt    string      `json:"prompt"`
	Suffix    string      `json:"suffix,omitempty"`
	Images    []string    `json:"images,omitempty"`
	Format    string      `json:"format,omitempty"`
	Options   interface{} `json:"options,omitempty"`
	System    string      `json:"system,omitempty"`
	Raw       bool        `json:"raw,omitempty"`
	KeepAlive string      `json:"keep_alive,omitempty"`
	Stream    bool        `json:"stream"`
}

type ResponseGenerate struct {
	Model              string `json:"model"`
	Created            string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

type Client interface {
	GenerateAPI(request RequestGenerate) ([]ResponseGenerate, error)
}

func NewDefaultClient(timeout int, baseURL string) Client {
	return &defaultClient{
		baseURL: baseURL,
		timeout: timeout,
	}
}

func EncodeBase64(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

type defaultClient struct {
	baseURL string
	timeout int // seconds
}

func (c defaultClient) GenerateAPI(request RequestGenerate) ([]ResponseGenerate, error) {

	endPoint, err := url.JoinPath(c.baseURL, "/api/generate")
	if err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	rBody := bytes.NewReader(reqBody)

	req, err := http.NewRequest(http.MethodPost, endPoint, rBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responses []ResponseGenerate
	if request.Stream {
		responses := parseGenerateStream(response.Body)
		return responses, nil
	}

	decoder := json.NewDecoder(response.Body)
	var resp ResponseGenerate
	err = decoder.Decode(&resp)
	if err != nil {
		return nil, err
	}
	responses = append(responses, resp)
	return responses, nil
}

func parseGenerateStream(body io.ReadCloser) []ResponseGenerate {

	var responses []ResponseGenerate
	reader := bufio.NewReader(body)
loop:
	for {
		ln, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break loop
			}
			continue
		}
		if ln[0] != '{' && ln[len(ln)-1] != '}' {
			continue
		}
		var response ResponseGenerate
		err = json.Unmarshal([]byte(ln), &response)
		if err != nil {
			continue
		}
		responses = append(responses, response)
	}
	return responses
}

func readImage(fname string) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if filepath.Base(pwd) != "ai-engineering" {
		return nil, fmt.Errorf("error calling path")
	}

	imageFile := filepath.Join(pwd, "examples", "ollama", "go", "ex3", fname)
	data, err := os.ReadFile(imageFile)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {

	img, err := readImage("cat.jpeg")
	b64Img := EncodeBase64(img)
	if err != nil {
		log.Fatal(err)
	}
	timeoutDuration := 120 //seconds
	client := NewDefaultClient(timeoutDuration, DefaultLocalBaseURL)
	reqBody := RequestGenerate{
		Model:  "llava",
		Prompt: "Is this a picture of a cat?",
		Images: []string{b64Img},
	}
	resp, err := client.GenerateAPI(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response:", resp)
}
