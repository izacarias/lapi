package responses

type ProblemDetails struct {
	// Type of the problem, typically a URI that identifies the problem type
	Detail string `json:"detail,omitempty"`
	// Instance string `json:"instance,omitempty"`
	Status int `json:"status"`
	// Title string `json:"title,omitempty"`
	// Type string `json:"type,omitempty"`
}
