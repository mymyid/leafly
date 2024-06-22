package ktp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"

	"gocv.io/x/gocv"
)

func DetectandCropKTP(ktpdt *KTPProps) (buf *gocv.NativeByteBuffer, err error) {
	// Dekode string base64 menjadi gambar
	img, err := base64ToMat(ktpdt.Base64Str)
	if err != nil {
		log.Fatalf("Error decoding base64: %v", err)
	}
	defer img.Close()

	// Proses gambar (deteksi KTP dan transformasi perspektif)
	processedImg, err := processImage(img)
	if err != nil {
		log.Fatalf("Error processing image: %v", err)
	}
	defer processedImg.Close()

	// Encode gocv.Mat to byte slice
	buf, err = gocv.IMEncode(gocv.JPEGFileExt, processedImg)
	if err != nil {
		return
	}
	defer buf.Close()
	// Convert the buffer to a base64 string
	ktpdt.Base64Str = base64.StdEncoding.EncodeToString(buf.GetBytes())

	// Tampilkan gambar hasil
	//window := gocv.NewWindow("Processed Image")
	//defer window.Close()
	//window.IMShow(processedImg)
	//window.WaitKey(0)
	return
}

// Fungsi untuk mendekode base64 menjadi gocv.Mat
func base64ToMat(base64Str string) (gocv.Mat, error) {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return gocv.NewMat(), err
	}
	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return gocv.NewMat(), err
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return gocv.NewMat(), err
	}

	mat, err := gocv.IMDecode(buf.Bytes(), gocv.IMReadColor)
	if err != nil {
		return gocv.NewMat(), err
	}
	return mat, nil
}

// Fungsi untuk memproses gambar (deteksi KTP dan transformasi perspektif)
func processImage(img gocv.Mat) (gocv.Mat, error) {
	// Konversi ke grayscale
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// Deteksi tepi menggunakan Canny
	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(gray, &edges, 50, 150)

	// Temukan kontur
	contours := gocv.FindContours(edges, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	// Filter kontur berdasarkan ukuran dan bentuk
	var largestContour gocv.PointVector
	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		area := gocv.ContourArea(contour)
		if area < 1000 {
			continue
		}

		// ApproxPolyDP untuk mendapatkan bentuk kontur yang lebih tepat
		approx := gocv.ApproxPolyDP(contour, 0.02*gocv.ArcLength(contour, true), true)
		defer approx.Close()
		if approx.Size() == 4 {
			largestContour = approx
			break
		}
	}

	// Jika kontur ditemukan
	if largestContour.Size() == 4 {
		// Urutkan titik berdasarkan posisi (Top-left, Top-right, Bottom-right, Bottom-left)
		orderedPoints := orderPoints(largestContour)

		// Tentukan titik tujuan untuk transformasi perspektif
		dst := gocv.NewPointVectorFromPoints([]image.Point{
			{0, 0},
			{300, 0},
			{300, 200},
			{0, 200},
		})

		// Hitung matriks transformasi perspektif
		src := gocv.NewPointVectorFromPoints(orderedPoints)
		warpMat := gocv.GetPerspectiveTransform(src, dst)

		// Terapkan transformasi perspektif
		warped := gocv.NewMat()
		gocv.WarpPerspective(img, &warped, warpMat, image.Point{300, 200})

		return warped, nil
	} else {
		return gocv.NewMat(), fmt.Errorf("KTP tidak ditemukan!")
	}
}

// Fungsi untuk mengurutkan titik sudut (Top-left, Top-right, Bottom-right, Bottom-left)
func orderPoints(points gocv.PointVector) []image.Point {
	// Asumsikan bahwa points memiliki panjang 4
	sorted := make([]image.Point, 4)
	add := func(p1, p2 image.Point) image.Point { return image.Point{p1.X + p2.X, p1.Y + p2.Y} }
	diff := func(p1, p2 image.Point) image.Point { return image.Point{p1.X - p2.X, p1.Y - p2.Y} }

	minSum, maxSum := points.At(0), points.At(0)
	minDiff, maxDiff := points.At(0), points.At(0)
	for i := 1; i < points.Size(); i++ {
		p := points.At(i)
		if add(minSum, p).X < add(minSum, minSum).X {
			minSum = p
		}
		if add(maxSum, p).X > add(maxSum, maxSum).X {
			maxSum = p
		}
		if diff(minDiff, p).X < diff(minDiff, minDiff).X {
			minDiff = p
		}
		if diff(maxDiff, p).X > diff(maxDiff, maxDiff).X {
			maxDiff = p
		}
	}
	sorted[0] = minSum
	sorted[2] = maxSum
	sorted[1] = minDiff
	sorted[3] = maxDiff
	return sorted
}
