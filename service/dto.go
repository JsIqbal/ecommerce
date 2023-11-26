package service

type ResponseData struct {
	Timestamp   int64       `json:"timestamp"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}
