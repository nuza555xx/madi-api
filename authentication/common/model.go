package common

import "time"

type BaseSchemas struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}