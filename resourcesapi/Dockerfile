FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/files
RUN swag init -g ./cmd/resourceapi/resourceapi.go -o cmd/docs
RUN go build -o resourceapi ./cmd/resourceapi
EXPOSE 8082
CMD ["./resourceapi"]