package api

type Page struct {
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}

type PageRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
