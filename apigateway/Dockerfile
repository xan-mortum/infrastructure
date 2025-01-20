FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/files
RUN swag init -g ./cmd/apigateway/apigateway.go -o cmd/docs
RUN go build -o apigateway ./cmd/apigateway
EXPOSE 8081
CMD ["./apigateway"]