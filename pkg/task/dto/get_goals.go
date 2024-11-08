package dto

type GetGoalsRequest struct {
	Keyword   string `query:"keyword"`
	PageToken string `query:"pageToken"`
}

type GetGoalsResponse struct {
	Goals         []Goal `json:"goals"`
	NextPageToken string `json:"nextPageToken"`
}
