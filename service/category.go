package service

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ParentID  string `json:"parent_id"`
	Sequence  string `json:"sequence"`
	StatusID  int    `json:"status_id"`
	CreatedAt int64  `json:"created_at"`
}

type CategoryResult struct {
	Categories []Category `json:"categories"`
	Total      int64      `json:"total"`
	Page       int64      `json:"page"`
	Limit      int64      `json:"limit"`
}
