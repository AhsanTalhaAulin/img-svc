FROM golang:1.19 as builder



WORKDIR /src
COPY ./ ./

# RUN go mod tidy
# RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix 'static' -o /app .

#  final image
FROM alpine:latest

COPY --from=builder /app /app
EXPOSE 8080

ENTRYPOINT ["/app"]