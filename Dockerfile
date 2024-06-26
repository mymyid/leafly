# Gunakan image dasar golang alpine
FROM golang:1.22.4-alpine

# Install dependencies
RUN apk update && apk add --no-cache \
    build-base \
    cmake \
    git \
    libjpeg-turbo-dev \
    libpng-dev \
    tiff-dev \
    opencv-dev \
    && rm -rf /var/cache/apk/*

# Set environment variable untuk GoCV
ENV CGO_CPPFLAGS="-I/usr/include"
ENV CGO_LDFLAGS="-L/usr/lib -lopencv_core -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_videoio -lopencv_objdetect"

# Buat direktori kerja
WORKDIR /app

# Copy kode Go ke direktori kerja
COPY . .

# Unduh dan instal GoCV
RUN go get -u -d gocv.io/x/gocv

# Build aplikasi Go
RUN go build -o face_detect .

# Menambahkan swap
RUN fallocate -l 512M /swapfile; \
	chmod 0600 /swapfile; \
	mkswap /swapfile;
# Command untuk menjalankan aplikasi
CMD ["./face_detect"]
