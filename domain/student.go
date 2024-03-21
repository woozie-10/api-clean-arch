package domain

type Student struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Group     string `json:"group"`
	Course    string `json:"course"`
	Specialty string `json:"specialty"`
}