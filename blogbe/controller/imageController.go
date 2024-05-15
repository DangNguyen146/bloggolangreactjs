package controller

import (
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Upload(c *fiber.Ctx) error {
	form, erro := c.MultipartForm()
	if erro != nil {
		return erro
	}
	files := form.File["image"]
	fileName := ""
	for _, file := range files {
		fileName = randLetter(10) + file.Filename
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return err
		}
	}
	return c.JSON(fiber.Map{"url": os.Getenv("URL_APP") + "/api/uploads" + fileName})
}
