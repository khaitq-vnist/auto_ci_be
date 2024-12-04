package response

type ThirdPartyListExecutionResponse struct {
	Page              int64                `json:"page"`
	PageSize          int64                `json:"page_size"`
	TotalPageCount    int64                `json:"total_page_count"`
	ElementCount      int64                `json:"element_count"`
	TotalElementCount int64                `json:"total_element_count"`
	Executions        []*ExecutionResponse `json:"executions"`
}
type ExecutionResponse struct {
	ID               int64                    `json:"id"`
	StartDate        int64                    `json:"start_date"`
	FinishDate       int64                    `json:"finish_date"`
	Status           string                   `json:"status"`
	TriggeredOn      string                   `json:"triggered_on"`
	Branch           *ExecutionBranchResponse `json:"branch"`
	ActionExecutions []*ActionExecution       `json:"action_executions,omitempty"`
}
type ExecutionBranchResponse struct {
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

type ActionExecution struct {
	Status   string  `json:"status"`
	Progress float32 `json:"progress"`
	Action   Action  `json:"action"`
}
type Action struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	TriggerTime         string `json:"trigger_time"`
	StartDate           int64  `json:"start_date"`
	FinishDate          int64  `json:"finish_date"`
	LastExecutionStatus string `json:"last_execution_status"`
}
