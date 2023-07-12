FROM golang AS builder
WORKDIR /target-services
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o ts .
CMD ["./ts"]
