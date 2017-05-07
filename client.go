package runscope

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
	"io/ioutil"
	"log"
	"strings"
)

type Client struct {
	ApiUrl      string
	AccessToken string
	Http        *http.Client
}

type Team struct {
	Name string
	Id   string
}

type response struct {
	Meta  metaResponse           `json:"meta"`
	Data  map[string]interface{} `json:"data"`
	Error errorResponse          `json:"error"`
}

type collectionResponse struct {
	Meta  metaResponse  `json:"meta"`
	Data  []interface{} `json:"data"`
	Error errorResponse `json:"error"`
}

type errorResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error"`
}

type metaResponse struct {
	Status string `json:"status"`
}

func NewClient(apiUrl string, accessToken string) *Client {
	client := Client{
		ApiUrl:      apiUrl,
		AccessToken: accessToken,
		Http:        cleanhttp.DefaultClient(),
	}

	return &client
}

func (client *Client) createResource(
	resource interface{}, resourceType string, resourceName string, endpoint string) (*response, error) {
	log.Printf("[DEBUG] creating %s %s", resourceType, resourceName)

	bytes, err := json.Marshal(resource)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] 	request: POST %s %s", endpoint, string(bytes))

	req, err := client.newRequest("POST", endpoint, bytes)
	if err != nil {
		return nil, err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] 	response: %d %s", resp.StatusCode, bodyString)

	if resp.StatusCode >= 300 {
		errorResp := new(errorResponse)
		if err = json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return nil, fmt.Errorf("Error creating %s: %s", resourceType, resourceName)
		} else {
			return nil, fmt.Errorf("Error creating %s: %s, status: %d reason: %q", resourceType,
				resourceName, errorResp.Status, errorResp.ErrorMessage)
		}
	} else {
		response := new(response)
		json.Unmarshal(bodyBytes, &response)
		return response, nil
	}
}

func (client *Client) readResource(resourceType string, resourceName string, endpoint string) (*response, error) {
	log.Printf("[DEBUG] reading %s %s", resourceType, resourceName)
	response := new(response)

	req, err := client.newRequest("GET", endpoint, nil)
	if err != nil {
		return response, err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] 	%d %s", resp.StatusCode, bodyString)

	if resp.StatusCode >= 300 {
		errorResp := new(errorResponse)
		if err = json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return response, fmt.Errorf("Status: %s Error reading %s: %s",
				resp.Status, resourceType, resourceName)
		} else {
			return response, fmt.Errorf("Status: %s Error reading %s: %s, reason: %q",
				resp.Status, resourceType, resourceName, errorResp.ErrorMessage)
		}
	} else {
		json.Unmarshal(bodyBytes, &response)
		return response, nil
	}
}

func (client *Client) updateResource(resource interface{}, resourceType string, resourceName string, endpoint string) (interface{}, error) {
	log.Printf("[DEBUG] updating %s %s", resourceType, resourceName)
	response := response{}
	bytes, err := json.Marshal(resource)
	if err != nil {
		return "", err
	}

	log.Printf("[DEBUG] 	request: PUT %s %s", endpoint, string(bytes))
	req, err := client.newRequest("PUT", endpoint, bytes)
	if err != nil {
		return response, err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] 	response: %d %s", resp.StatusCode, bodyString)

	if resp.StatusCode >= 300 {
		errorResp := new(errorResponse)
		if err = json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return response, fmt.Errorf("Status: %s Error reading %s: %s",
				resp.Status, resourceType, resourceName)
		} else {
			return response, fmt.Errorf("Status: %s Error reading %s: %s, reason: %q",
				resp.Status, resourceType, resourceName, errorResp.ErrorMessage)
		}
	} else {
		json.Unmarshal(bodyBytes, &response)
		return response, nil
	}
}

func (client *Client) deleteResource(resourceType string, resourceName string, endpoint string) (error) {
	log.Printf("[DEBUG] deleting %s %s", resourceType, resourceName)
	req, err := client.newRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] 	request: DELETE %s", endpoint)
	resp, err := client.Http.Do(req)
	log.Printf("[DEBUG] 	response: %d", resp.StatusCode)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Printf("[DEBUG] %s", bodyString)

		errorResp := new(errorResponse)
		if err = json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return fmt.Errorf("Status: %s Error deleting %s: %s",
				resp.Status, resourceType, resourceName)
		} else {
			return fmt.Errorf("Status: %s Error deleting %s: %s, reason: %q",
				resp.Status, resourceType, resourceName, errorResp.ErrorMessage)
		}
	}

	return nil
}

func (client *Client) newFormUrlEncodedRequest(method string, endpoint string, data url.Values) (*http.Request, error) {

	var urlStr string
	urlStr = client.ApiUrl + endpoint
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Error during parsing request URL: %s", err)
	}

	req, err := http.NewRequest(method, url.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %s", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.AccessToken))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (client *Client) newRequest(method string, endpoint string, body []byte) (*http.Request, error) {

	var urlStr string
	urlStr = client.ApiUrl + endpoint
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Error during parsing request URL: %s", err)
	}

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %s", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.AccessToken))
	req.Header.Add("Accept", "application/json")

	if method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}
