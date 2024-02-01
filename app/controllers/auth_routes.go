package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/oyesaurav/go-todo/pkg/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "pseudo-random"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("G_CLIENT_ID"),
		ClientSecret: os.Getenv("G_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("G_REDIRECT"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}
}

func Login(c *fiber.Ctx) error {
	URL := googleOauthConfig.AuthCodeURL(oauthStateString)
	return c.Redirect(URL)
}

func Callback(c *fiber.Ctx) error {

	if state := c.Query("state"); state != oauthStateString {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	code := c.Query("code")

  token, err := googleOauthConfig.Exchange(c.Context(), code)
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  // convert token to user data
  profile, err := utils.ConvertToken(token.AccessToken)
  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  return c.JSON(profile)
}
