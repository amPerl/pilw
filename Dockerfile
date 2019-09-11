FROM golang:1.13-alpine AS base

RUN apk add make

COPY go.mod /app/
COPY go.sum /app/
COPY Makefile /app/

COPY cmd /app/cmd
COPY pkg /app/pkg

WORKDIR /app/

RUN make build

FROM alpine
COPY --from=base /app/pilw /usr/bin/

CMD ["pilw", "user", "info"]