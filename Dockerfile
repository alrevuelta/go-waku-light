FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /main

FROM golang:1.20

WORKDIR /

COPY --from=build /main /main

ENTRYPOINT ["/main"]