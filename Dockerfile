FROM golang:1.21.5-alpine3.19

WORKDIR /app

COPY . ./
RUN go install github.com/fzipp/gofind@latest
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-initializr

EXPOSE 8080

CMD ["/go-initializr"]