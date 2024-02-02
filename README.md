# GO listed

Go server to set tasks and subtasks with auto priority update and call reminders

## API Endpoints

### 1. Create Task

- **Endpoint:** `POST /api/v1/tasks`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2024-02-04T12:00:00Z"
    }
    ```

### 2. Create Sub Task

- **Endpoint:** `POST /api/v1/subtasks`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "task_id": "a1f8a062-76cb-4e78-a6e1-7e165c57fe07"
    }
    ```

### 3. Get All User Tasks

- **Endpoint:** `GET /api/v1/tasks?page=1&limit=10`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "priority" : Number,  (optional)
      "due_date" : time     (optional)
    }
    ```

### 4. Get All User Sub Tasks

- **Endpoint:** `GET /api/v1/subtasks?page=1&limit=10`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "task_id" : uuid,  (optional)
    }
    ```

### 5. Update Task

- **Endpoint:** `PUT /api/v1/tasks`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "id": "",
      "due_date": "2024-02-15",
      "status": "DONE"
    }
    ```

### 6. Update Sub Task

- **Endpoint:** `PUT /api/v1/subtasks`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "id": "",
      "status": 1
    }
    ```

### 7. Delete Task (Soft Deletion)

- **Endpoint:** `DELETE /api/v1/tasks`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "id": "",
    }
    ```

### 8. Delete Sub Task (Soft Deletion)

- **Endpoint:** `DELETE /api/v1/subtasks/`
- **Input:**
  - Headers: `Authorization: Bearer <jwt-auth-token>`
  - Body:
    ```json
    {
      "id": "",
    }
    ```

## Cron Jobs

### 1. Cron Logic for Changing Priority of Task based on Due Date

- This cron job runs periodically to update the priority of tasks based on their due dates.
- Priority Rules:
  - Priority 0: Due date is today.
  - Priority 1: Due date is between tomorrow and the day after tomorrow.
  - Priority 2: Due date is between 3-4 days
  - Priority 3: Otherwise

### 2. Cron Logic for Voice Calling using Twilio

- This cron job runs periodically to initiate voice calls using Twilio for tasks that pass their due dates.
- Calling Priority:
  - Calls are made based on the priority of the user.
  - Users with priority 0 are called first, followed by priority 1, and then priority 2.
  - A user is called only if the previous user does not attend the call.
- Priority is fetched from the user table.

## Usage / Get Started

### Go Initialization

1. Make the .env file from the .env.example
2. In the root directory:

```sh
go mod tidy
go run main.go
```

The server would start in port 5000.
