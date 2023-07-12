# --- builder ---
FROM golang:1.20.6-alpine3.17 as builder
LABEL stage=builder
RUN apk add git
WORKDIR /build

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -ldflags "-X 'main.serverBuildString=pretendo.pikmin3.docker'" -v -o server

# --- runner ---
FROM alpine:3.17 as runner
WORKDIR /build

COPY --from=builder /build/server /build/
CMD ["/build/server"]
