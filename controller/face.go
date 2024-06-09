package controller

import (
	"leafly/helper/face"
	"leafly/model"

	"github.com/gofiber/fiber/v2"
)

func FaceDetect(ctx *fiber.Ctx) error {
	var h model.Secret
	err := ctx.ReqHeaderParser(&h)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}
	var msg face.FaceDetect
	err = ctx.BodyParser(&msg)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	_, err = face.DetectandCropFace(&msg)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	// Call GithubUpload with the file header
	//content, response, err := ghupload.GithubUpload(config.GHCreds, buf.GetBytes(), "mymyid", "face", msg.IDUser+"/"+msg.IDFile+".jpg", true)
	//if err != nil {
	//	return ctx.Status(fiber.StatusFailedDependency).JSON(fiber.Map{"error": err.Error()})
	//}

	//ret := fmt.Sprintf("Upload successful: %v, response: %v\n", content, response)
	return ctx.Status(fiber.StatusOK).JSON(msg)
}
