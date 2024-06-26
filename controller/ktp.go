package controller

import (
	"leafly/config"
	"leafly/helper/ghupload"
	"leafly/helper/ktp"
	"leafly/model"

	"github.com/gofiber/fiber/v2"
)

func KTPDetect(ctx *fiber.Ctx) error {
	var h model.Secret
	var body model.FaceInfo
	err := ctx.ReqHeaderParser(&h)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusEarlyHints).JSON(body)
	}
	if h.Secret != config.Secret {
		body.Error = "Secret salah"
		return ctx.Status(fiber.StatusForbidden).JSON(body)
	}
	var msg ktp.KTPProps
	err = ctx.BodyParser(&msg)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(body)
	}
	_, err = ktp.DetectandCropKTP(&msg)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusConflict).JSON(body)
	}
	if config.GHCreds.GitHubAccessToken == "" {
		body.Error = "access token tidak ada: " + config.GHCreds.GitHubAccessToken
		return ctx.Status(fiber.StatusExpectationFailed).JSON(body)
	}
	// Call GithubUpload with the file header
	content, response, filehash, err := ghupload.GithubUploadJPG(config.GHCreds, msg.Base64Str, "mymyid", "ktp", msg.IDUser, false)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusFailedDependency).JSON(body)
	}
	body.Commit = *content.Commit.SHA
	body.Remaining = response.Rate.Remaining
	body.FileHash = filehash
	body.PhoneNumber = msg.IDUser

	return ctx.Status(fiber.StatusOK).JSON(body)
}
