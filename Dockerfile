FROM golang:1.20.4-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go build -o short-links-app ./cmd/web/main.go
CMD ["./short-links-app", "--storage=cache"]