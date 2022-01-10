FROM golang:1.17.0-bullseye

RUN go version
ENV GOPATH=/


COPY ./ ./

RUN go mod tidy
RUN go mod download

RUN go build -o backend-test ./cmd/main.go

CMD ["./backend-test"]