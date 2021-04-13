package dao

// swagger:parameters User
type User struct {
	// keyin username
	//
	// required: true
	// in: query
	Username string `json:"username"`
	// keyin password
	//
	// required: true
	// in: query
	Password string `json:"password"`
}
