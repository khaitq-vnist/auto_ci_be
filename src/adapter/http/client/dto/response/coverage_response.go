package response

type CoverageResponse struct {
	Component Component `json:"component"`
}

type Component struct {
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Qualifier   string    `json:"qualifier"`
	Measures    []Measure `json:"measures"`
}

type Measure struct {
	Metric    string `json:"metric"`
	Value     string `json:"value"` // Use float64 if you want to parse the value as a number
	BestValue bool   `json:"bestValue"`
}
