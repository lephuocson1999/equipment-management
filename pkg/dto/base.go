package dto

type FindRequest struct {
	Queries  []Query
	Join     string
	Sort     string
	GroupBy  string
	Having   string
	Offset   int
	Limit    int
	Preloads []string
}

type Query struct {
	Query interface{}
	Args  interface{}
}
