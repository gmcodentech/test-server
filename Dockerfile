FROM golang AS build
RUN mkdir /Example
WORKDIR /Example
COPY index.go .
RUN go build -o main index.go



FROM ubuntu:latest
WORKDIR /tmp
COPY --from=build /Example/main .
EXPOSE 8080
CMD ["./main"]