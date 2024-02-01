package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	_"github.com/oyesaurav/go-todo/pkg/utils"
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