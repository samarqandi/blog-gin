FROM golang:1.20.4-alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

EXPOSE 8080

CMD ["go", "run", "."]
