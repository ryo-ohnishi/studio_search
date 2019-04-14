package rooms

import (
	"time"
)

type Room struct {
	id         int
	studio_id  int
	name       string
	size       string
	url        string
	remark     string
	created_at time.Time
	updated_at time.Time
}
