package controller

import (
	"fmt"
	"leafly/config"
	"leafly/helper/face"
	"leafly/helper/ghupload"
	"leafly/model"

	"github.com/gofiber/fiber/v2"
)

func FaceDetect(ctx *fiber.Ctx) error {
	var h model.Secret
	err := ctx.ReqHeaderParser(&h)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var msg face.FaceDetect
	err = ctx.BodyParser(&msg)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	buf, err := face.DetectandCropFace(&msg)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Create a multipart.FileHeader from the encoded byte slice
	fileHeader, err := ghupload.CreateFileHeader(buf.GetBytes())
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Call GithubUpload with the file header
	content, response, err := ghupload.GithubUpload(config.GHCreds, fileHeader, "mymyid", "face", msg.IDUser+msg.IDFile+".jpg", true)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ret := fmt.Sprintf("Upload successful: %v, response: %v\n", content, response)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": ret})
}
