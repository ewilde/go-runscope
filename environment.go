package runscope

import (
	"time"
	"fmt"
)

// See https://www.runscope.com/docs/api/environments
type Environment struct {
	Id                  string            `json:"id,omitempty"`
	Name                string            `json:"name,omitempty"`
	Script              string            `json:"script,omitempty"`
	PreserveCookies     bool              `json:"preserve_cookies,omitempty"`
        TestId              string            `json:"test_id,omitempty"`
        InitialVariables    map[string]string `json:"initial_variables,omitempty"`
        Integrations        []Integration     `json:"integrations,omitempty"`
	Regions             []string          `json:"regions,omitempty"`
	VerifySsl           bool              `json:"verify_ssl,omitempty"`
	ExportedAt          time.Time         `json:"exported_at,omitempty"`
	RetryOnFailure      bool              `json:"retry_on_failure,omitempty"`
	RemoteAgents        []LocalMachine    `json:"remote_agents,omitempty"`
	WebHooks            []string          `json:"webhooks,omitempty"`
	ParentEnvironmentId string            `json:"parent_environment_id,omitempty"`
	EmailSettings       EmailSettings     `json:"emails,omitempty"`
	ClientCertificate   string            `json:"client_certificate,omitempty"`
}

type EmailSettings struct {
	NotifyAll       bool      `json:"notify_all,omitempty"`
	NotifyOn        string    `json:"notify_on,omitempty"`
	NotifyThreshold int       `json:"notify_threshold,omitempty"`
	Recipients      []Contact `json:"recipients,omitempty"`
}

type Integration struct {
	Id               string `json:"id"`
	IntegrationType  string `json:"integration_type"`
	Description      string `json:"description,omitempty"`
}

type LocalMachine struct {
	Name             string `json:"name"`
	Uuid             string `json:"uuid"`
}

func NewEnvironment() *Environment {
	return new(Environment)
}

func (client *Client) CreateSharedEnvironment(environment *Environment, bucket *Bucket) (*Environment, error) {
	return client.createEnvironment(environment, fmt.Sprintf("/buckets/%s/environments", bucket.Key))
}

func (client *Client) CreateTestEnvironment(environment *Environment, test *Test) (*Environment, error) {
	return client.createEnvironment(environment, fmt.Sprintf("/buckets/%s/tests/%s/environments",
		test.Bucket.Key, test.Id))
}

func (client *Client) ReadSharedEnvironment(environment *Environment, bucket *Bucket) (*Environment, error) {
	return client.readEnvironment(environment, fmt.Sprintf("/buckets/%s/environments/%s",
		bucket.Key, environment.Id))
}

func (client *Client) ReadTestEnvironment(environment *Environment, test *Test) (*Environment, error) {
	return client.readEnvironment(environment, fmt.Sprintf("/buckets/%s/tests/%s/environments/%s",
		test.Bucket.Key, test.Id, environment.Id))
}

func (client *Client) UpdateSharedEnvironment(environment *Environment, bucket *Bucket) (response, error) {
	resource, error := client.updateResource(environment, "environment", environment.Id,
		fmt.Sprintf("/buckets/%s/Environments/%s", bucket.Key, environment.Id))
	return resource.(response), error
}

func (client *Client) UpdateTestEnvironment(environment *Environment, test *Test) (response, error) {
	resource, error := client.updateResource(environment, "environment", environment.Id,
		fmt.Sprintf("/buckets/%s/tests/%s/environments/%s", test.Bucket.Key, test.Id, environment.Id))
	return resource.(response), error
}

func (client *Client) DeleteSharedEnvironment(environment *Environment, bucket *Bucket) error {
	return client.deleteResource("Environment", environment.Id,
		fmt.Sprintf("/buckets/%s/environments/%s", bucket.Key, environment.Id))
}

func (client *Client) DeleteTestEnvironment(environment *Environment, test *Test) error {
	return client.deleteResource("Environment", environment.Id,
		fmt.Sprintf("/buckets/%s/environments/%s/tests/%s", test.Bucket.Key, test.Id, environment.Id))
}

func (client *Client) createEnvironment(environment *Environment, endpoint string) (*Environment, error) {
	newResource, error := client.createResource(environment, "Environment", environment.Name, endpoint)
	if error != nil {
		return nil, error
	}

	newEnvironment, error := getEnvironmentFromResponse(newResource.Data)
	if error != nil {
		return nil, error
	}

	return newEnvironment, nil
}

func (client *Client) readEnvironment(environment *Environment, endpoint string) (*Environment, error) {
	resource, error := client.readResource("environment", environment.Id, endpoint)
	if error != nil {
		return nil, error
	}

	readEnvironment, error := getEnvironmentFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readEnvironment, nil
}

func getEnvironmentFromResponse(response interface{}) (*Environment, error) {
	environment := new(Environment)
	err := Decode(environment, response)
	return environment, err
}