package ktp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"gocv.io/x/gocv"
)

func DetectandCropKTP(ktpdt *KTPProps) (buf *gocv.NativeByteBuffer, err error) {
	// Decode base64 string into gocv.Mat
	img, err := base64ToMat(ktpdt.Base64Str)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64: %v", err)
	}
	defer img.Close()

	// Process image (detect KTP and perform perspective transformation)
	processedImg, err := processImage(img)
	if err != nil {
		return nil, fmt.Errorf("error processing image: %v", err)
	}
	defer processedImg.Close()

	// Encode gocv.Mat to byte slice (JPEG format)
	buf, err = gocv.IMEncode(gocv.JPEGFileExt, processedImg)
	if err != nil {
		return nil, fmt.Errorf("error encoding image to JPEG: %v", err)
	}
	defer buf.Close()

	// Convert the buffer to a base64 string
	ktpdt.Base64Str = base64.StdEncoding.EncodeToString(buf.GetBytes())

	return buf, nil
}

// Function to decode base64 string into gocv.Mat
func base64ToMat(base64Str string) (gocv.Mat, error) {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return gocv.NewMat(), fmt.Errorf("error decoding base64: %v", err)
	}

	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return gocv.NewMat(), fmt.Errorf("error decoding JPEG: %v", err)
	}

	// Encode the image back to bytes buffer
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return gocv.NewMat(), fmt.Errorf("error encoding JPEG: %v", err)
	}

	// Decode the encoded image into gocv.Mat
	mat, err := gocv.IMDecode(buf.Bytes(), gocv.IMReadColor)
	if err != nil {
		return gocv.NewMat(), fmt.Errorf("error decoding image into Mat: %v", err)
	}

	return mat, nil
}

// Function to process image (detect KTP and perform perspective transformation)
func processImage(img gocv.Mat) (gocv.Mat, error) {
	// Convert image to grayscale
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// Detect edges using Canny edge detection
	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(gray, &edges, 50, 150)

	// Find contours in the edges
	contours := gocv.FindContours(edges, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer contours.Close()

	// Filter contours based on size and shape to find the largest rectangle
	var largestContour gocv.PointVector
	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		area := gocv.ContourArea(contour)
		if area < 1000 {
			continue
		}

		// Approximate the contour to get a polygon with fewer vertices
		epsilon := 0.02 * gocv.ArcLength(contour, true)
		approx := gocv.ApproxPolyDP(contour, epsilon, true)
		defer approx.Close()

		// Check if the polygon is a quadrilateral (4 sides)
		if approx.Size() == 4 {
			largestContour = approx
			break
		}
	}

	// Check if a valid contour was found
	if largestContour.Size() != 4 {
		return gocv.NewMat(), fmt.Errorf("no valid rectangle contour found")
	}

	// Order the vertices of the rectangle
	orderedPoints := orderPoints(largestContour)

	// Define destination points for perspective transformation
	dst := gocv.NewPointVectorFromPoints([]image.Point{
		{0, 0},
		{300, 0},
		{300, 200},
		{0, 200},
	})

	// Calculate the perspective transform matrix
	src := gocv.NewPointVectorFromPoints(orderedPoints)
	warpMat := gocv.GetPerspectiveTransform(src, dst)

	// Apply perspective transformation
	warped := gocv.NewMat()
	gocv.WarpPerspective(img, &warped, warpMat, image.Point{300, 200})

	return warped, nil
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
