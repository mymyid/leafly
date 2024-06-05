package url

import (
	"leafly/controller"
	"leafly/helper/wrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {

	//page.Get("/ws", websocket.New(chatroot.RunSocket))
	page.Get("/webrtc", websocket.New(wrtc.RunWebRTCSocket)) // New route for WebRTC signaling
	page.Post("/face/detect", controller.FaceDetect)

}
