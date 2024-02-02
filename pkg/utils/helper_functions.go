package utils

import (
	"fmt"
	"time"

	_ "github.com/gofiber/fiber/v2"
	"github.com/oyesaurav/go-todo/app/models"
	"github.com/oyesaurav/go-todo/platform/database"
	"github.com/robfig/cron/v3"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func UpdateTaskPriority(task models.Task, db *database.Queries) models.Task {
	// Calculate the time difference between now and the due date.
	// If the due date is today.
	if time.Now().Before(task.DueDate) && task.DueDate.Before(time.Now().Add(24*time.Hour)) {
		task.Priority = 0
		fmt.Printf("Task ID %d priority updated. New priority: %d\n", task.ID, task.Priority)
	} else if time.Now().Before(task.DueDate) && task.DueDate.Before(time.Now().Add(48*time.Hour)) {
		// If the due date is between tomorrow and day after tomorrow.
		task.Priority = 1
		fmt.Printf("Task ID %d priority updated. New priority: %d\n", task.ID, task.Priority)
	} else if time.Now().Before(task.DueDate) && task.DueDate.Before(time.Now().Add(96*time.Hour)) {
		// If the due date is between tomorrow and day after tomorrow.
		task.Priority = 2
		fmt.Printf("Task ID %d priority updated. New priority: %d\n", task.ID, task.Priority)
	} else {
		// For other cases, you can add additional priority rules.
		// For simplicity, we don't change priority in other cases.
		fmt.Printf("Task ID %d priority not updated. Priority remains: %d\n", task.ID, task.Priority)
	}

	task.UpdatedAt = time.Now()

	if time.Now().After(task.DueDate) {
		// If the task is not completed, then call the users.
		CallUsers(db)
	}
	return task
}

func CallUsers(db *database.Queries) {
	client := twilio.NewRestClient()

	params := &api.CreateCallParams{}
	params.SetUrl("http://demo.twilio.com/docs/voice.xml")
	params.SetFrom("+16592183722")

	users, err := db.GetallUsers()
	if err != nil {
		fmt.Println("Error fetching users from the database:", err)
		return
	}

	for _, user := range users {
		params.SetTo(user.PhoneNumber)
		resp, err := client.Api.CreateCall(params)
		if err != nil {
			fmt.Println("Error initiating voice call:", err)
			continue // Move on to the next user on error
		}

		for i := 0; i < 10; i++ {
			// Sleep for 10 seconds before checking call status again
			time.Sleep(10 * time.Second)

			params := &api.FetchCallParams{}

			updatedCall, err := client.Api.FetchCall(*resp.Sid, params)
			if err != nil {
				fmt.Println("Error fetching call details:", err)
				continue
			}

			if *updatedCall.Status == "in-progress" || *updatedCall.Status == "queued" || *updatedCall.Status == "ringing" {
				fmt.Printf("Call to user %s is still in progress\n", user.PhoneNumber)
			} else if *updatedCall.Status == "completed" {
				fmt.Printf("Call to user %s is completed\n", user.PhoneNumber)
				// Handle completed call as needed
				return
			} else {
				fmt.Printf("Call to user %s failed with status: %s\n", user.PhoneNumber, *updatedCall.Status)
				// Move on to the next user if the call failed
				break
			}
		}
	}
}

func CronScheduler() {

	db, err := database.OpenDBConnection()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	fmt.Print("DB connected")
	// Create a new cron scheduler.
	c := cron.New()

	// Schedule the UpdateTaskPriority function to run every minute.
	_, err1 := c.AddFunc("* * * * *", func() {
		// Fetch tasks from the database.
		tasks, err := db.GetallTasks()
		if err != nil {
			fmt.Println("Error fetching tasks from the database:", err)
			return
		}

		// Iterate over tasks and update priorities.
		for _, task := range tasks {
			UpdateTaskPriority(task, db)
		}
	})
	if err1 != nil {
		fmt.Println("Error scheduling the first cron job:", err)
		return
	}

	// Start the cron scheduler.
	c.Start()

	// Run indefinitely, you may use a signal to gracefully stop the cron job.
	// select {}
}
