package domain

type Assessment struct {
	Username string           `json:"username"`
	Marks    map[string][]int `json:"marks"`
}
