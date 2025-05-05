FROM golang AS build
RUN mkdir /Example
WORKDIR /Example
COPY index.go .
RUN go build -o main index.go



FROM alpine:latest
WORKDIR /tmp
COPY --from=build /Example/main .
CMD ["./main"]