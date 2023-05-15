package handlers_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/minacio00/adopet/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateAbrigo(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name       string
		route      string
		expectCode int
		body       *strings.Reader
	}{
		{
			name:  "get status 200",
			route: "/abrigo",
			body: strings.NewReader(`{
				"nome": "theDogShelter",
				"cidade": "anapolis",
				"uf": "GO"
			}`),
			expectCode: 200,
		},
		{
			name:       "get status 400",
			route:      "/abrigo",
			expectCode: 400,
			body:       nil,
		},
	}

	app := fiber.New()
	app.Post("/abrigo", handlers.CreateAbrigo)

	for _, t := range tests {
		req := httptest.NewRequest("POST", "/abrigo", t.body)
		resp, _ := app.Test(req, 1)
		assert.Equalf()

	}

}
