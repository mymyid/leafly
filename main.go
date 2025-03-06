package main

import (
	"log"

	"leafly/config"
	"leafly/controller"
	"leafly/url"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mhale/smtpd"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//email inbond
	log.Println("SMTP MX Server berjalan di port 25...")
	err := smtpd.ListenAndServe(":25", controller.HandleMail, "MX Server", "")
	if err != nil {
		log.Fatal(err)
	}
	//go chatroot.RunHub()

	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(config.IPPort))
}
