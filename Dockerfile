FROM golang:alpine AS build
WORKDIR /go/src/watersupplynamanger
COPY . .
WORKDIR /go/src/watersupplynamanger/app
RUN go build -o /go/bin/appbinary

FROM scratch
COPY --from=build /go/bin/appbinary /go/bin/appbinary
EXPOSE 8080
ENTRYPOINT ["/go/bin/appbinary"]
