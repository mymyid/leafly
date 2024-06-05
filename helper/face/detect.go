package face

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

func DetectandCropFace(base64Str string) (fdetect FaceDetect, err error) {
	// Decode base64 string
	imgData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return
	}

	// Buat file sementara untuk menulis data gambar
	tmpFile, err := os.CreateTemp("", "image-*.jpg")
	if err != nil {
		return
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(imgData)
	if err != nil {
		return
	}
	tmpFile.Close()

	// Baca gambar dari file sementara
	img := gocv.IMRead(tmpFile.Name(), gocv.IMReadColor)
	if img.Empty() {
		err = errors.New("Error reading image")
		return
	}
	defer img.Close()

	// Inisialisasi detektor wajah dengan menggunakan model pre-trained
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("Error loading cascade file")
		return
	}

	// Deteksi wajah di gambar
	rects := classifier.DetectMultiScale(img)
	fdetect.Nfaces = len(rects)

	// Tandai setiap wajah yang terdeteksi dengan persegi panjang
	for _, r := range rects {
		gocv.Rectangle(&img, r, color.RGBA{0, 255, 0, 0}, 3)
	}

	fdetect.Base64Str, err = matToBase64(img)
	if err != nil {
		return
	}

	// Simpan hasil gambar dengan deteksi wajah ke file
	// resultFile := "result.jpg"
	// gocv.IMWrite(resultFile, img)
	// fmt.Println("Result saved to", resultFile)
	return
}

func matToBase64(img gocv.Mat) (string, error) {
	// Encode gambar ke buffer dalam format JPEG
	buf, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return "", fmt.Errorf("error encoding image: %v", err)
	}
	defer buf.Close()

	// Konversi buffer ke byte slice
	imgBytes := buf.GetBytes()

	// Encode byte slice ke string base64
	base64Str := base64.StdEncoding.EncodeToString(imgBytes)

	return base64Str, nil
}

type FaceDetect struct {
	Nfaces    int    `json:"nfaces,omitempty" bson:"nfaces,omitempty"`
	Base64Str string `json:"base64str,omitempty" bson:"base64str,omitempty"`
}
