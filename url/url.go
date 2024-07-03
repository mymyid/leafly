package url

import (
	"leafly/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {

	//page.Get("/ws", websocket.New(chatroot.RunSocket))
	//page.Get("/webrtc", websocket.New(wrtc.RunWebRTCSocket)) // New route for WebRTC signaling
	page.Post("/face/detect", controller.FaceDetect)
	page.Post("/ktp/detect", controller.KTPDetect)

	page.Post("/arsip/gambar/lmsdesa", controller.ArsipGambarLMSDesa)
	page.Post("/arsip/file/lmsdesa", controller.ArsipFileLMSDesa)
	page.Post("/arsip/file/logmeeting", controller.LogFileMeeting)

}
