package queries

import (
	_"fmt"
	_ "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/oyesaurav/go-todo/app/models"
)

// BookQueries struct for queries from Book model.
type SubTaskQueries struct {
	*sqlx.DB
}

// GetBooks method for getting all books.
func (q *SubTaskQueries) GetSubTasks() ([]models.SubTask, error) {
	// Define books variable.
	tasks := []models.SubTask{}

	// Define query string.
	query := `SELECT * FROM subtask`

	// Send query to database.
	err := q.Select(&tasks, query)
	if err != nil {
		// Return empty object and error.
		return tasks, err
	}

	// Return query result.
	return tasks, nil
}



