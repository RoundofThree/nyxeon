FROM golang:1.16-alpine AS builder 

RUN mkdir /build 
COPY . /build/ 
WORKDIR /build 
RUN go build 

FROM alpine 
RUN adduser -S -D -H -h /app appuser 
USER appuser 
COPY --from=builder /build/nyxeon /app/ 
COPY config/ /app/config 
WORKDIR /app 
CMD ["/app/nyxeon", "-m", "production"] 