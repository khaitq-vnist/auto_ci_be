package response

import (
	"github.com/khaitq-vnist/auto_ci_be/core/entity/dto/response"
	"time"
)

type BuddyListExecutionResponse struct {
	Page              int64                     `json:"page"`
	PageSize          int64                     `json:"page_size"`
	TotalPageCount    int64                     `json:"total_page_count"`
	ElementCount      int64                     `json:"element_count"`
	TotalElementCount int64                     `json:"total_element_count"`
	Executions        []*BuddyExecutionResponse `json:"executions"`
}
type BuddyExecutionResponse struct {
	ID               int64                         `json:"id"`
	StartDate        time.Time                     `json:"start_date"`
	FinishDate       time.Time                     `json:"finish_date"`
	Status           string                        `json:"status"`
	TriggeredOn      string                        `json:"triggered_on"`
	Branch           *BuddyExecutionBranchResponse `json:"branch"`
	ActionExecutions []*BuddyActionExecution       `json:"action_executions"`
}
type BuddyExecutionBranchResponse struct {
	Name    string `json:"name"`
	Default bool   `json:"default"`
}
type BuddyActionExecution struct {
	Status   string      `json:"status"`
	Progress float32     `json:"progress"`
	Action   BuddyAction `json:"action"`
}
type BuddyAction struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Type                string    `json:"type"`
	TriggerTime         string    `json:"trigger_time"`
	StartDate           time.Time `json:"start_date"`
	FinishDate          time.Time `json:"finish_date"`
	LastExecutionStatus string    `json:"last_execution_status"`
}

func ToListExecutionResponse(rsp *BuddyListExecutionResponse) *response.ThirdPartyListExecutionResponse {
	executions := make([]*response.ExecutionResponse, 0)
	for _, execution := range rsp.Executions {
		executions = append(executions, &response.ExecutionResponse{
			ID:          execution.ID,
			StartDate:   execution.StartDate.Unix(),
			FinishDate:  execution.FinishDate.Unix(),
			Status:      execution.Status,
			TriggeredOn: execution.TriggeredOn,
			Branch: &response.ExecutionBranchResponse{
				Name:    execution.Branch.Name,
				Default: execution.Branch.Default,
			},
		})
	}
	return &response.ThirdPartyListExecutionResponse{
		Page:              rsp.Page,
		PageSize:          rsp.PageSize,
		TotalPageCount:    rsp.TotalPageCount,
		TotalElementCount: rsp.TotalElementCount,
		ElementCount:      rsp.ElementCount,
		Executions:        executions,
	}
}
func ToExecutionDetail(rsp *BuddyExecutionResponse) *response.ExecutionResponse {
	actionExecutions := make([]*response.ActionExecution, 0)
	for _, actionExecution := range rsp.ActionExecutions {
		actionExecutions = append(actionExecutions, &response.ActionExecution{
			Status:   actionExecution.Status,
			Progress: actionExecution.Progress,
			Action: response.Action{
				ID:                  actionExecution.Action.ID,
				Name:                actionExecution.Action.Name,
				Type:                actionExecution.Action.Type,
				TriggerTime:         actionExecution.Action.TriggerTime,
				StartDate:           actionExecution.Action.StartDate.Unix(),
				FinishDate:          actionExecution.Action.FinishDate.Unix(),
				LastExecutionStatus: actionExecution.Action.LastExecutionStatus,
			},
		})
	}
	return &response.ExecutionResponse{
		ID:          rsp.ID,
		StartDate:   rsp.StartDate.Unix(),
		FinishDate:  rsp.FinishDate.Unix(),
		Status:      rsp.Status,
		TriggeredOn: rsp.TriggeredOn,
		Branch: &response.ExecutionBranchResponse{
			Name:    rsp.Branch.Name,
			Default: rsp.Branch.Default,
		},
		ActionExecutions: actionExecutions,
	}
}
