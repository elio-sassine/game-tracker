FROM golang:1.24.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd
RUN mkdir -p /certs

EXPOSE 8080 8443

CMD ["/server"]