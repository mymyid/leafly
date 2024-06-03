package config

import (
	"leafly/helper"

	"github.com/gofiber/fiber/v2"
)

var IPPort, Net = helper.GetAddress()

var Iteung = fiber.Config{
	Prefork:       false,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "DoMyID",
	AppName:       "Domyikado",
	Network:       Net,
}
