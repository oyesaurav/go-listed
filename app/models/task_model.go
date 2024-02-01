package models

import (
	_"database/sql/driver"
	_"encoding/json"
	_"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	DueDate     time.Time `db:"duedate" json:"due_date"`
	Status      string    `db:"status" json:"status"`
	CreatedAt   time.Time `db:"createdat" json:"created_at"`
	UpdatedAt   time.Time `db:"updatedat" json:"updated_at"`
	DeletedAt   time.Time `db:"deletedat" json:"deleted_at"`
	// store a ref to all subtasks
	// SubTasks    []SubTask `db:"sub_tasks" json:"sub_tasks"`
}

type User struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required"`
	PhoneNumber int       `db:"phonenumber" json:"phone_number" validate:"required"`
	Priority    int       `db:"priority" json:"priority"`
}

type SubTask struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required"`
	TaskID     uuid.UUID `db:"taskid" json:"task_id" validate:"required"`
	Status     int       `db:"status" json:"status"`
	CreatedAt  time.Time `db:"createdat" json:"created_at"`
	UpdatedAt  time.Time `db:"updatedat" json:"updated_at"`
	DeletedAt  time.Time `db:"deletedat" json:"deleted_at"`
}
