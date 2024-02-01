package queries

import (
	"fmt"
	"strconv"

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

func (q *TaskQueries) CreateSubTask(t *models.SubTask) error {
	// Define query string.
	query := `INSERT INTO subtask VALUES ($1, $2, $3, $4, $5, $6)`

	// Send query to database.
	result, err := q.Exec(query, t.ID, t.TaskID, t.Status, t.CreatedAt, t.UpdatedAt, t.DeletedAt)
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

func (q *TaskQueries) GetSubTasks() ([]models.SubTask, error) {
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

func (q *TaskQueries) UpdateTask(t *models.Task) error {
    // Define query string.
    query := "UPDATE task SET"
    
    // Check which fields are provided and add them to the query.
    var params []interface{}
    paramCount := 1  // Counter for placeholder position

    if !t.DueDate.IsZero() {
        query += " duedate = $" + strconv.Itoa(paramCount) + ","
        params = append(params, t.DueDate)
        paramCount++
    }

    if t.Status != "" {
        query += " status = $" + strconv.Itoa(paramCount) + ","
        params = append(params, t.Status)
        paramCount++
    }

    // Add more conditions for other fields as needed.

	query += " updatedat = $" + strconv.Itoa(paramCount) + ","
	params = append(params, t.UpdatedAt)
	paramCount++
	
    // Remove the trailing comma if there are parameters.
    if len(params) > 0 {
        query = query[:len(query)-1]
    }

    // Add the WHERE clause.
    query += " WHERE id = $" + strconv.Itoa(paramCount)
	params = append(params, t.ID)

    // Send query to database.
    _, err := q.Exec(query, params...)
    if err != nil {
        return err
    }

    return nil
}

func (q *TaskQueries) UpdateSubTask(t *models.SubTask) error {
    // Define query string.
    query := "UPDATE subtask SET"
    
    // Check which fields are provided and add them to the query.
    var params []interface{}
    paramCount := 1  // Counter for placeholder position

	query += " status = $" + strconv.Itoa(paramCount) + ","
	params = append(params, t.Status)
	paramCount++

    // Add more conditions for other fields as needed.

	query += " updatedat = $" + strconv.Itoa(paramCount) + ","
	params = append(params, t.UpdatedAt)
	paramCount++
	
    // Remove the trailing comma if there are parameters.
    if len(params) > 0 {
        query = query[:len(query)-1]
    }

    // Add the WHERE clause.
    query += " WHERE id = $" + strconv.Itoa(paramCount)
	params = append(params, t.ID)

    // Send query to database.
    _, err := q.Exec(query, params...)
    if err != nil {
        return err
    }

    return nil
}
