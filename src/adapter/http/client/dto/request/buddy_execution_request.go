package request

import "encoding/json"

type BuddyRunExecutionRequest struct {
	ToRevision ToRevision `json:"to_revision"`
}
type ToRevision struct {
	Revision string `json:"revision" default:"HEAD"`
}

func (r *BuddyRunExecutionRequest) ToJSON() []byte {
	value, _ := json.Marshal(r)
	return value
}
