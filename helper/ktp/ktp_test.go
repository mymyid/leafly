package ktp

import (
	"testing"
)

// func TestBase64ToMat(t *testing.T) {
// 	base64Str := "YOUR_BASE64_STRING_HERE"
// 	_, err := base64ToMat(base64Str)
// 	if err != nil {
// 		t.Errorf("base64ToMat() error = %v", err)
// 	}
// }

// func TestProcessImage(t *testing.T) {
// 	base64Str := "YOUR_BASE64_STRING_HERE"
// 	img, err := base64ToMat(base64Str)
// 	if err != nil {
// 		t.Fatalf("Error decoding base64: %v", err)
// 	}
// 	defer img.Close()

// 	_, err = processImage(img)
// 	if err != nil {
// 		t.Errorf("processImage() error = %v", err)
// 	}
// }

// func TestOrderPoints(t *testing.T) {
// 	points := gocv.NewPointVector()
// 	points.Append(image.Pt(0, 0))
// 	points.Append(image.Pt(1, 1))
// 	points.Append(image.Pt(1, 0))
// 	points.Append(image.Pt(0, 1))

// 	ordered := orderPoints(points)
// 	if len(ordered) != 4 {
// 		t.Errorf("orderPoints() expected length = 4, got = %d", len(ordered))
// 	}
// }

func TestDetectandCropKTP(t *testing.T) {
	ktpProps := &KTPProps{
		Base64Str: "YOUR_BASE64_STRING_HERE",
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DetectandCropKTP() panicked: %v", r)
		}
	}()

	DetectandCropKTP(ktpProps)
}
