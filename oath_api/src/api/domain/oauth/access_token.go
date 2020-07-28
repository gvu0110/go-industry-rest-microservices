package oauth

import "time"

// AccessToken struct
type AccessToken struct {
	Token   string `json:"token"`
	UserID  int64  `json:"user_id"`
	Expires int64  `json:"expires"`
}

// IsExpired function check if a token is expired or not
func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).UTC().Before(time.Now().UTC())
}
