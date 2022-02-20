package response

import "main/app/internal/model"

type LoginRes struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
