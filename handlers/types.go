package handlers

type NewTodo struct {
	Title  string `json:"title"`
	Status bool   `json:"boolean"`
	UserId string `json:"user_id"`
}

type FullTodo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"boolean"`
	UserId string `json:"user_id"`
}
