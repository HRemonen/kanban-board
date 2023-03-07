FROM golang:1.20.0

WORKDIR /usr/src/app

RUN go install github.com/HRemonen/kanban-board/src/server@latest

COPY . .

RUN go mod tidy

CMD ["air", "./cmd/main.go", "-b", "0.0.0.0"]