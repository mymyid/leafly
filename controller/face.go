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
	var body model.FaceInfo
	err := ctx.ReqHeaderParser(&h)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusEarlyHints).JSON(body)
	}
	if h.Secret != config.Secret {
		body.Error = "Secret salah "
		return ctx.Status(fiber.StatusForbidden).JSON(body)
	}
	var msg face.FaceDetect
	err = ctx.BodyParser(&msg)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(body)
	}
	_, err = face.DetectandCropFace(&msg)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusConflict).JSON(body)
	}
	if msg.Nfaces == 0 {
		body.FileHash = msg.Base64Str
		body.Error = "Mukanya ga kelihatan kak, coba pas pas in deh"
		return ctx.Status(fiber.StatusGone).JSON(body)
	}
	if msg.Nfaces > 1 {
		body.FileHash = msg.Base64Str
		body.Error = "Minta tolong tutupin dengan tangan atau benda atau barang. Pada bagian yang dikotakin selain daripada wajah kakak ya... Terus coba lagi ya kak... Makasih..."
		return ctx.Status(fiber.StatusMultipleChoices).JSON(body)
	}
	if config.GHCreds.GitHubAccessToken == "" {
		body.Error = "access token tidak ada: " + config.GHCreds.GitHubAccessToken
		return ctx.Status(fiber.StatusExpectationFailed).JSON(body)
	}
	// Call GithubUpload with the file header
	content, response, filehash, err := ghupload.GithubUploadJPG(config.GHCreds, msg.Base64Str, "mymyid", "face", msg.IDUser, false)
	if err != nil {
		body.FileHash = msg.Base64Str
		body.Error = err.Error()
		return ctx.Status(fiber.StatusFailedDependency).JSON(body)
	}
	body.Commit = *content.Commit.SHA
	body.Remaining = response.Rate.Remaining
	body.FileHash = filehash
	body.PhoneNumber = msg.IDUser

	return ctx.Status(fiber.StatusOK).JSON(body)
}
