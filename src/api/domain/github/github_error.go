package github

/*
{
	"message": "Repository creation failed.",
	"errors": [
		{
			"resource": "Repository",
			"code": "custom",
			"field": "name",
			"message": "name already exists on this account"
		}
	],
    "documentation_url": "https://developer.github.com/v3"
}
*/

type GitHubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"documentation_url"`
	Errors           []GitHubError `json:"errors"`
}

type GitHubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
