package configs

import (
	"strconv"
	"time"
	"os"

	"github.com/gofiber/fiber/v2"

)

func FiberConfig() fiber.Config {

	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ServerHeader: "Listed Go",
		AppName: "Go Listed TODO",
        ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}