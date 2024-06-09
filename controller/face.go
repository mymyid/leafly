package controller

import (
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
		return ctx.Status(fiber.StatusEarlyHints).JSON(fiber.Map{"error": err.Error()})
	}
	if h.Secret != config.Secret {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Secret Salah"})
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
	if msg.Nfaces == 0 {
		return ctx.Status(fiber.StatusFailedDependency).JSON(fiber.Map{"error": "Tidak ditemukan muka"})
	}
	if msg.Nfaces > 1 {
		return ctx.Status(fiber.StatusFailedDependency).JSON(fiber.Map{"error": "Harus selfie tidak boleh ramean"})
	}
	if config.GHCreds.GitHubAccessToken == "" {
		return ctx.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{"error": "access token tidak ada: " + config.GHCreds.GitHubAccessToken})
	}
	// Call GithubUpload with the file header
	content, response, filehash, err := ghupload.GithubUploadJPG(config.GHCreds, msg.Base64Str, "mymyid", "face", msg.IDUser, false)
	if err != nil {
		return ctx.Status(fiber.StatusFailedDependency).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"commit": content.Commit.SHA, "remaining": response.Rate.Remaining, "filehash": filehash, "iduser": msg.IDUser})
}
