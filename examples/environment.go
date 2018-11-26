package examples

import (
	"github.com/ewilde/go-runscope"
)

func createSharedEnvironment() {
	environment := &runscope.Environment{
		Name: "tf_environment",
		InitialVariables: map[string]string{
			"VarA": "ValB",
			"VarB": "ValB",
		},
		Integrations: []*runscope.EnvironmentIntegration{
			{
				ID:              "27e48b0d-ba8e-4fe0-bcaa-dd9de08dc47d",
				IntegrationType: "pagerduty",
			},
			{
				ID:              "574f4560-0f50-41da-a2f7-bdce419ad378",
				IntegrationType: "slack",
			},
		},
	}

	environment, err := client.CreateSharedEnvironment(environment, createBucket())
	if err != nil {
		runscope.ErrorF(1, "error creating environment: %s", err)
	}
}
