FROM golang:1.21.1

WORKDIR /app

COPY go.mod math.go /app/

RUN go build -o math

CMD [ "./math" ]