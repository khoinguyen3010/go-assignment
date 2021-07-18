package models

import (
	"time"
)

type BaseObject struct {
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	CreatedBy string    `json:"created_by" form:"created_by"`
	UpdatedBy string    `json:"updated_by" form:"updated_by"`
}

type BaseSuccessResponse struct {
	Message string `json:"message"`
}

type BaseFailedResponse struct {
	Error string `json:"error"`
}
