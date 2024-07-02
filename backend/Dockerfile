# 基於官方 Golang 映像建立 Docker 映像
FROM golang:1.22-alpine

# 設置工作目錄
WORKDIR /app

# 複製所有源代碼到容器中
COPY . .

# 下載依賴並建置應用程式
RUN go mod download
RUN go build -o main .

# 暴露端口，用於本地測試
EXPOSE 8080

# 定義啟動命令
CMD ["./main"]
