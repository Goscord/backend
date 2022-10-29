package webhook

import (
	"net/http"
	"os"

	"github.com/Goscord/DocGen/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v47/github"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func POSTHandler(c *fiber.Ctx) error {
	req := new(http.Request)

	err := fasthttpadaptor.ConvertRequest(c.Context(), req, false)
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(&utils.HTTPResponse{Message: "Internal Server Error"})
	}

	payload, err := github.ValidatePayload(req, []byte(os.Getenv("GITHUB_SECRET")))
	if err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.JSON(&utils.HTTPResponse{Message: "SHA1 and SHA256 signatures are not valid"})
	}

	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(&utils.HTTPResponse{Message: "Cannot parse GitHub webhook"})
	}

	if _, ok := event.(github.PushEvent); ok {
		// ToDo: handle
	}

	return c.JSON(&utils.HTTPResponse{Message: "Webhook event successfully parsed"})
}
