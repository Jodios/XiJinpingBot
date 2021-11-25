FROM golang:alpine as builder
RUN mkdir /build 
ADD ./src /build/
WORKDIR /build 
RUN ls
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY --from=builder /build/config.yaml /
COPY --from=builder /build/banner.txt /
WORKDIR /app
CMD ["./main"]