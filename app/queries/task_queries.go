package queries

import (
	"fmt"
	_ "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/oyesaurav/go-todo/app/models"
)

// BookQueries struct for queries from Book model.
type TaskQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *TaskQueries) GetTasks() ([]models.Task, error) {
	// Define books variable.
	tasks := []models.Task{}

	// Define query string.
	query := `SELECT * FROM task`

	// Send query to database.
	err := q.Select(&tasks, query)
	if err != nil {
		// Return empty object and error.
		return tasks, err
	}

	// Return query result.
	return tasks, nil
}

func (q *TaskQueries) CreateTask(t *models.Task) error {
	// Define query string.
	query := `INSERT INTO task VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	result, err := q.Exec(query, t.ID, t.Title, t.Description, t.DueDate, t.Status, t.CreatedAt, t.UpdatedAt, t.DeletedAt)
	if err != nil {
		return err
	}

	// Check the number of rows affected.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Log the result.
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	// This query returns nothing.
	return nil
}
