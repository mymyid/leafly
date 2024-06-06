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
	err = face.DetectandCropFace(&msg)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func FaceCount(ctx *fiber.Ctx) error {

	var msg face.FaceDetect
	err := ctx.BodyParser(&msg)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	err = face.DetectandCropFace(&msg)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"nfaces": msg.Nfaces})
}
