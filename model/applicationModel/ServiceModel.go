package applicationModel

type DefaultOperator struct {
	DataType string   `json:"data_type"`
	Operator []string `json:"operator"`
}

var GetListMovieValidOperator map[string]DefaultOperator

func InitiateDefaultOperator() {
	GetListMovieValidOperator = make(map[string]DefaultOperator)
	GetListMovieValidOperator["title"] = DefaultOperator{DataType: "char", Operator: []string{"like", "eq"}}
	GetListMovieValidOperator["description"] = DefaultOperator{DataType: "char", Operator: []string{"like", "eq"}}
	GetListMovieValidOperator["artists"] = DefaultOperator{DataType: "char", Operator: []string{"like", "eq"}}
	GetListMovieValidOperator["genres"] = DefaultOperator{DataType: "char", Operator: []string{"like", "eq"}}
}
