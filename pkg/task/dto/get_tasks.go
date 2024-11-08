package dto

type GetTasksRequest struct {
	GoalID    string `query:"goalId"`
	Keyword   string `query:"keyword"`
	Status    string `query:"status"`
	PageToken string `query:"pageToken"`
}

type GetTasksResponse struct {
	Tasks         []Task `json:"tasks"`
	NextPageToken string `json:"nextPageToken"`
}
