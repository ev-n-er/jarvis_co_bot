FROM golang:1.18 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build ./cmd/bot

FROM alpine:3.14 as production
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder app/bot ./bin
COPY --from=builder app/Procfile ./bin
EXPOSE 8080
CMD ./bin/bot