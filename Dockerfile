FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/jgsqware/promtail-sidecar/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/promtail-sidecar
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
COPY --from=builder /go/bin/promtail-sidecar /go/bin/promtail-sidecar
ENTRYPOINT ["/go/bin/promtail-sidecar"]