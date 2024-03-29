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
func (q *TaskQueries) GetTasks(t *models.Task, pageSize string, page string) ([]models.Task, error) {
	// Define books variable.
	tasks := []models.Task{}
	paramCount := 1
	query := "SELECT * FROM task WHERE deletedat = '0001-01-01 00:00:00'"
	var params []interface{}
	// Add filters
	if t.Priority != -1 {
		query += " AND priority = " + strconv.Itoa(t.Priority)
	}

	if !t.DueDate.IsZero() {
		query += " AND duedate = $" + strconv.Itoa(paramCount)
		params = append(params, t.DueDate)
		paramCount++
	}

	// Add pagination
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	query += " ORDER BY createdat DESC LIMIT $" + strconv.Itoa(paramCount) + " OFFSET $" + strconv.Itoa(paramCount+1)
	params = append(params, pageSizeInt, (pageInt-1)*pageSizeInt)

	// Send query to database.
	err := q.Select(&tasks, query, params...)
	if err != nil {
		// Return empty object and error.
		return tasks, err
	}

	// Return query result.
	return tasks, nil
}

func (q *TaskQueries) GetallTasks() ([]models.Task, error) {
	// Define books variable.
	tasks := []models.Task{}
	query := "SELECT * FROM task WHERE deletedat = '0001-01-01 00:00:00'"

	// Send query to database.
	err := q.Select(&tasks, query)
	if err != nil {
		// Return empty object and error.
		return tasks, err
	}

	// Return query result.
	return tasks, nil
}

func (q *TaskQueries) GetSubTasks(t *models.SubTask, pageSize string, page string, taskid int) ([]models.SubTask, error) {
	// Define books variable.
	tasks := []models.SubTask{}
	paramCount := 1
	query := "SELECT * FROM subtask WHERE deletedat = '0001-01-01 00:00:00'"
	var params []interface{}
	// Add filters
	if taskid != -1 {
		query += " AND taskid = $" + strconv.Itoa(paramCount)
		params = append(params, t.TaskID)
		paramCount++
	}
	// Add pagination
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	query += " ORDER BY createdat DESC LIMIT $" + strconv.Itoa(paramCount) + " OFFSET $" + strconv.Itoa(paramCount+1)
	params = append(params, pageSizeInt, (pageInt-1)*pageSizeInt)

	// Send query to database.
	err := q.Select(&tasks, query, params...)
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

func (q *TaskQueries) UpdateTask(t *models.Task) error {
	// Define query string.
	query := "UPDATE task SET"

	// Check which fields are provided and add them to the query.
	var params []interface{}
	paramCount := 1 // Counter for placeholder position

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
	query += " WHERE id = $" + strconv.Itoa(paramCount) + " AND deletedat = '0001-01-01 00:00:00'"
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
	paramCount := 1 // Counter for placeholder position

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
	query += " WHERE id = $" + strconv.Itoa(paramCount) + " AND deletedat = '0001-01-01 00:00:00'"
	params = append(params, t.ID)

	// Send query to database.
	_, err := q.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (q *TaskQueries) DeleteTask(t *models.Task) error {
	query := "UPDATE task SET deletedat = $1 WHERE id = $2 AND deletedat = '0001-01-01 00:00:00'"

	// Send query to database.
	_, err := q.Exec(query, t.DeletedAt, t.ID)
	if err != nil {
		return err
	}

	return nil
}

func (q *TaskQueries) DeleteSubTask(t *models.SubTask) error {
	query := "UPDATE subtask SET deletedat = $1 WHERE id = $2 AND deletedat = '0001-01-01 00:00:00'"

	// Send query to database.
	_, err := q.Exec(query, t.DeletedAt, t.ID)
	if err != nil {
		return err
	}

	return nil
}

func (q *TaskQueries) GetallUsers() ([]models.User, error) {
	// Define books variable.
	users := []models.User{}
	query := "SELECT * FROM user_table ORDER BY priority ASC"

	// Send query to database.
	err := q.Select(&users, query)
	if err != nil {
		// Return empty object and error.
		return users, err
	}

	// Return query result.
	return users, nil
}