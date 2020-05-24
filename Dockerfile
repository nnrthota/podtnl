FROM golang:1.13.4-stretch as builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 make build


FROM alpine:3.10
WORKDIR /app
COPY --from=builder /app/podtnl .

EXPOSE 5000

ENTRYPOINT ["./podtnl"]
