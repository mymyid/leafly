package config

import (
	"leafly/helper"
	"leafly/helper/ghupload"
	"os"

	"github.com/gofiber/fiber/v2"
)

var IPPort, Net = helper.GetAddress()

var Secret = os.Getenv("SECRET")

var Iteung = fiber.Config{
	Prefork:       false,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "DoMyID",
	AppName:       "Domyikado",
	Network:       Net,
}

var GHCreds = ghupload.GHCreds{
	GitHubAccessToken: os.Getenv("GH_ACCESS_TOKEN"),
	GitHubAuthorName:  "Rolly Maulana Awangga",
	GitHubAuthorEmail: "awangga@gmail.com",
}
