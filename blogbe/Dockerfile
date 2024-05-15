# Sử dụng image Golang chính thức làm base
FROM golang:1.22.3

# Thiết lập thư mục làm việc
WORKDIR /

# Sao chép mã nguồn vào thư mục làm việc trong container
COPY . .

# Tải và cài đặt dependencies
RUN go mod download

# Biên dịch ứng dụng
RUN go build -o main .

# Expose cổng mà ứng dụng sử dụng
EXPOSE 3000

# Chạy ứng dụng khi container được khởi chạy
CMD ["./main"]
