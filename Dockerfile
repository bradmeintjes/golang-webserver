FROM golang:latest AS builder

# allows go to retrieve the dependencies for the build step
RUN apt-get install git

RUN mkdir /build/

WORKDIR /build/
ADD . /build/

# compilation
RUN CGO_ENABLED=0 go build -o /build/app .

FROM alpine:latest

# run as non root
RUN adduser -D -u 10000 brad
USER brad

WORKDIR /

COPY --from=builder /build/app /

CMD ["/app"]




