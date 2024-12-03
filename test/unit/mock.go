package test

type MockUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type MockUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
