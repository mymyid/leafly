package face

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func DetectandCropFace(msg *FaceDetect) (buf *gocv.NativeByteBuffer, err error) {
	// Decode base64 string
	imgData, err := base64.StdEncoding.DecodeString(msg.Base64Str)
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
		err = errors.New("Error loading cascade file")
		return
	}

	// Deteksi wajah di gambar
	rects := classifier.DetectMultiScale(img)
	msg.Nfaces = len(rects)

	// Tandai setiap wajah yang terdeteksi dengan persegi panjang
	//for _, r := range rects {
	//	gocv.Rectangle(&img, r, color.RGBA{0, 255, 0, 0}, 3)
	//}

	// Crop wajah pertama yang terdeteksi
	face := img.Region(rects[0])
	defer face.Close()

	// Encode gocv.Mat to byte slice
	buf, err = gocv.IMEncode(gocv.JPEGFileExt, face)
	if err != nil {
		return
	}
	defer buf.Close()
	// Convert the buffer to a base64 string
	msg.Base64Str = base64.StdEncoding.EncodeToString(buf.GetBytes())

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

func saveIMG(img gocv.Mat, resultFile string) bool {
	// Simpan hasil gambar dengan deteksi wajah ke file
	return gocv.IMWrite(resultFile, img)

}
