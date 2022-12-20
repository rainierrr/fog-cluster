####################### Build stage #######################
FROM golang:1.19-alpine AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main /go/src/main.go

####################### Run stage #######################
FROM alpine:3.17

WORKDIR /go/src

RUN apk upgrade --update

COPY --from=builder /go/src/main .


EXPOSE 8080

CMD [ "/go/src/main" ]
