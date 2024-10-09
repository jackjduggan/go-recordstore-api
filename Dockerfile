FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

#RUN go build -o recordapi .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/recordapi

ENV PORT=8080

EXPOSE 8080

CMD ["./recordapi"]