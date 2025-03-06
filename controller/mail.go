package controller

import (
	"log"
	"net"
)

func HandleMail(remoteAddr net.Addr, from string, to []string, data []byte) error {
	log.Printf("Email masuk dari: %s ke %v\n", from, to)
	// Bisa disimpan ke database atau file
	return nil
}
