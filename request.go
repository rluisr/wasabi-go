package wasabi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	JSONRPC string   `json:"jsonrpc"`
	ID      string   `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type Response struct {
	Result json.RawMessage `json:"result"`
	Error  *clientErr      `json:"error"`
}

type clientErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

func (c *Client) request(body interface{}) (*Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read body")
	}

	// Todo temp
	fmt.Println(string(respBody))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not 200 got %d. body: %d", resp.StatusCode, respBody)
	}

	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, fmt.Errorf("response error. code: %d, message: %s", response.Error.Code, response.Error.Message)
	}

	return &response, nil
}
