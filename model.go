package main

type Workflow struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Steps        []Step   `json:"steps"`   // Steps in the workflow
	Metrics      Metrics  `json:"metrics"` // Metrics for the workflow
	StartTrigger Trigger  `json:"start_trigger"`
	EndState     EndState `json:"end_state"`
}

type Step struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Inputs      []Input     `json:"inputs"`
	Outputs     []Output    `json:"outputs"`
	Service     *Service    `json:"service,omitempty"`      // Optional: Service responsible for this step
	SubWorkflow *Workflow   `json:"sub_workflow,omitempty"` // Optional: Nested workflow
	Conditions  []Condition `json:"conditions"`             // Conditions for executing this step
	NextSteps   []string    `json:"next_steps"`             // List of next step IDs
}

type Input struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`     // "json", "query", "header"
	Payload  interface{}       `json:"payload"`  // Actual input data (e.g., JSON body)
	Metadata map[string]string `json:"metadata"` // Optional metadata (e.g., headers)
}

type Output struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`     // "json", "status_code"
	Payload  interface{}       `json:"payload"`  // Expected output data
	Metadata map[string]string `json:"metadata"` // Optional metadata (e.g., headers)
}

type Service struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	BaseURL     string            `json:"base_url"`     // Base URL of the web service
	Endpoint    string            `json:"endpoint"`     // Specific API path
	Method      string            `json:"method"`       // HTTP method (GET, POST, etc.)
	Headers     map[string]string `json:"headers"`      // Custom headers
	QueryParams map[string]string `json:"query_params"` // Query parameters
	Auth        Auth              `json:"auth"`         // Authentication details
}

type Auth struct {
	Type     string `json:"type"`               // "basic", "bearer", "api_key", etc.
	Username string `json:"username,omitempty"` // For basic auth
	Password string `json:"password,omitempty"` // For basic auth
	Token    string `json:"token,omitempty"`    // For bearer or API key
}

type Condition struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Expression  string `json:"expression"` // Logical condition (e.g., response.status == 200)
}

type Trigger struct {
	ID       string `json:"id"`
	Type     string `json:"type"`     // "webhook", "schedule", "event"
	Criteria string `json:"criteria"` // Defines what starts the workflow (e.g., event type)
}

type Metrics struct {
	ID             string   `json:"id"`
	TotalTime      string   `json:"total_time"`     // Total workflow execution time
	StepLatencies  []string `json:"step_latencies"` // Latency per step
	CompletedSteps int      `json:"completed_steps"`
	PendingSteps   int      `json:"pending_steps"`
	FailedSteps    []string `json:"failed_steps"` // List of failed step IDs
}

type EndState struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Outputs     []Output `json:"outputs"`
}
