FROM golang:1.10-alpine AS builder

WORKDIR /go/src/github.com/alileza/tomato

COPY . ./

RUN apk add --update make git
RUN make build

# ---

FROM alpine

COPY --from=builder /go/src/github.com/alileza/tomato/bin/tomato /bin/tomato

ENTRYPOINT  [ "/bin/tomato" ]
CMD         [ "--config.file=/config.yml", \
              "--features.path=/features/" ]
