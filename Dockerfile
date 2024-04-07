# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.22-alpine
ARG TARGETARCH
ARG TARGETVARIANT
RUN apk --no-cache add tzdata

WORKDIR /src/

COPY go.mod go.sum /src/

RUN go mod download

COPY pkg /src/pkg
COPY internal /src/internal

ARG MAINFILE=bot
COPY cmd/${MAINFILE} /src/cmd/${MAINFILE}
RUN cd /src/cmd/${MAINFILE} && CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT//v/} go build -o /src/main .

FROM scratch

COPY --from=0 /src/main /

WORKDIR /
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=America/New_York
USER nobody

ENTRYPOINT [ "/main" ]