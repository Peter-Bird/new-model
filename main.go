package main

import "fmt"

func main() {
	// Define a deeply nested workflow
	nestedWorkflow := Workflow{
		ID:          "nested_workflow_001",
		Name:        "Data Validation Workflow",
		Description: "Validates and enriches data before further processing",
		Steps: []Step{
			{
				ID:          "nested_step_001",
				Name:        "Validate Format",
				Description: "Ensure data is in the correct format",
				Service:     &Service{ID: "service_validate", Name: "Validation Service", BaseURL: "https://api.example.com", Endpoint: "/validate", Method: "POST"},
				NextSteps:   []string{"nested_step_002"},
			},
			{
				ID:          "nested_step_002",
				Name:        "Enrich Data",
				Description: "Enrich data with additional attributes",
				Service:     &Service{ID: "service_enrich", Name: "Enrichment Service", BaseURL: "https://api.example.com", Endpoint: "/enrich", Method: "POST"},
				NextSteps:   nil, // End of the nested workflow
			},
		},
	}

	// Define the main workflow
	mainWorkflow := Workflow{
		ID:          "main_workflow_001",
		Name:        "Main Data Processing Workflow",
		Description: "Processes data through multiple steps and workflows",
		Steps: []Step{
			{
				ID:          "step_001",
				Name:        "Receive Data",
				Description: "Receives data for processing",
				Service:     &Service{ID: "service_receive", Name: "Data Receiver", BaseURL: "https://api.example.com", Endpoint: "/receive", Method: "POST"},
				NextSteps:   []string{"step_002"},
			},
			{
				ID:          "step_002",
				Name:        "Validate and Enrich Data",
				Description: "Performs validation and enrichment via a nested workflow",
				SubWorkflow: &nestedWorkflow, // Nested workflow
				NextSteps:   []string{"step_003"},
			},
			{
				ID:          "step_003",
				Name:        "Finalize Data",
				Description: "Finalizes and stores the processed data",
				Service:     &Service{ID: "service_finalize", Name: "Finalization Service", BaseURL: "https://api.example.com", Endpoint: "/finalize", Method: "POST"},
				NextSteps:   nil, // End of the main workflow
			},
		},
	}

	// Output or process the workflow
	jsonStr, err := WorkflowToJSON(mainWorkflow)
	if err != nil {
		return
	}
	s, err := PrettyPrintJSON(jsonStr)
	if err != nil {
		return
	}

	fmt.Println(s)
}
