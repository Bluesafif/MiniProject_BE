package applicationModel

type DefaultOperator struct {
	DataType string   `json:"data_type"`
	Operator []string `json:"operator"`
}

var GetListMovieValidOperator map[string]DefaultOperator

func InitiateDefaultOperator() {
	GetListMovieValidOperator = make(map[string]DefaultOperator)
	GetListMovieValidOperator["code"] = DefaultOperator{DataType: "char", Operator: []string{"like", "eq"}}
	GetListMovieValidOperator["name"] = DefaultOperator{DataType: "char", Operator: []string{"like", "eq"}}
}
