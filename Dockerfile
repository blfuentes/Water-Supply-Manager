FROM golang:alpine AS build
WORKDIR /go/src/watersupplynamanger
COPY . .
WORKDIR /go/src/watersupplynamanger/app
RUN go build -o /go/bin/appbinary

FROM scratch
COPY --from=build /go/bin/appbinary /go/bin/appbinary
ENV MONGODB_URI=mongodb://root:example@mongo:27017/?authSource=admin
EXPOSE 8080
ENTRYPOINT ["/go/bin/appbinary"]
