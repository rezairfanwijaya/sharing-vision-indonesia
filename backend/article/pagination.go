package article

type PaginationArticle struct {
	Limit     int         `json:"limit"`
	TotalData int         `json:"total_data"`
	Articles  interface{} `json:"articles"`
}
