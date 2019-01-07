
FROM golang:1.11-alpine as builder
RUN apk update && apk add --no-cache git

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /go/bin/app

FROM alpine:3.8

COPY /templates/ /go/bin/templates
COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 9000

ENTRYPOINT ["./go/bin/app"]


