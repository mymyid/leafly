package controller

import (
	"leafly/helper/face"

	"github.com/gofiber/fiber/v2"
)

func FaceDetect(ctx *fiber.Ctx) error {

	var msg face.FaceDetect
	err := ctx.BodyParser(&msg)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	hasil, err := face.DetectandCropFace(msg.Base64Str)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(hasil)
}
