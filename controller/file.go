package controller

import (
	"fmt"
	"leafly/config"
	"leafly/helper/face"
	"leafly/helper/ghupload"
	"leafly/model"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ArsipGambarLMSDesa(ctx *fiber.Ctx) error {
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
	if config.GHCreds.GitHubAccessToken == "" {
		body.Error = "access token tidak ada: " + config.GHCreds.GitHubAccessToken
		return ctx.Status(fiber.StatusExpectationFailed).JSON(body)
	}
	// Mendapatkan waktu saat ini
	now := time.Now()

	// Mendapatkan tahun, bulan, dan tanggal sebagai string
	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%02d", int(now.Month()))
	day := fmt.Sprintf("%02d", now.Day())
	path := year + "/" + month + "/" + day
	// Call GithubUpload with the file header
	content, response, _, err := ghupload.GithubUploadJPG(config.GHCreds, msg.Base64Str, "domyid", "lmsdesa", path, false)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusFailedDependency).JSON(body)
	}
	body.Commit = *content.Commit.SHA
	body.Remaining = response.Rate.Remaining
	body.FileHash = *content.Content.Path
	body.PhoneNumber = msg.IDUser
	return ctx.Status(fiber.StatusOK).JSON(body)
}

func ArsipFileLMSDesa(ctx *fiber.Ctx) error {
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
	if config.GHCreds.GitHubAccessToken == "" {
		body.Error = "access token tidak ada: " + config.GHCreds.GitHubAccessToken
		return ctx.Status(fiber.StatusExpectationFailed).JSON(body)
	}
	// Mendapatkan waktu saat ini
	now := time.Now()

	// Mendapatkan tahun, bulan, dan tanggal sebagai string
	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%02d", int(now.Month()))
	day := fmt.Sprintf("%02d", now.Day())
	path := year + "/" + month + "/" + day
	// Call GithubUpload with the file header
	content, response, _, err := ghupload.GithubUploadFile(config.GHCreds, msg.Base64Str, "domyid", "lmsdesa", path, filepath.Ext(msg.IDFile), false)
	if err != nil {
		body.Error = err.Error()
		return ctx.Status(fiber.StatusFailedDependency).JSON(body)
	}
	body.Commit = *content.Commit.SHA
	body.Remaining = response.Rate.Remaining
	body.FileHash = *content.Content.Path
	body.PhoneNumber = msg.IDUser
	return ctx.Status(fiber.StatusOK).JSON(body)
}
