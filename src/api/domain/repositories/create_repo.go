package repositories

// CreateRepoRequest struct to create a new GitHub repo with the name
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRepoResponse struct
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
