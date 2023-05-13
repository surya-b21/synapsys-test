FROM golang:1.20.4-alpine AS builder

# install upx
RUN apk add upx

#install goupx
RUN go get github.com/pwaller/goupx

WORKDIR /app
COPY . .

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#build app
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-s -w -extldflags "-static"' -o smartdio.run

RUN goupx synapsys.run

FROM alpine:latest

RUN apk add ca-certificates

COPY --from=builder /app/synapsys /server

CMD ["/server"]