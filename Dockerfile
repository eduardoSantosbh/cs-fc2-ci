FROM golang:1.19

WORKDIR /app

COPY . .

RUN go build math.go

RUN chmod +x math

CMD ["./math"]