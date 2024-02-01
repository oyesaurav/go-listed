package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	title string `db:"title" json:"title"`
	description string `db:"description" json:"description"`
	due_date time.Time `db:"due_date" json:"due_date"`

}

type User struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	phone_number int `db:"phone_number" json:"phone_number" validate:"required"`
	priority  int `db:"priority" json:"priority"`
}

type SubTask struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	task_id  uuid.UUID  `db:"task_id" json:"task_id" validate:"required"`
	status  int `db:"status" json:"status"`
	ceated_at time.Time `db:"created_at" json:"created_at"`
	updated_at time.Time `db:"updated_at" json:"updated_at"`
	deleted_at time.Time `db:"deleted_at" json:"deleted_at"`
}

type BookAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

// Value make the BookAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b *BookAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan make the BookAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *BookAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}