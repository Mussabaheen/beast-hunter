FROM golang:latest as build

WORKDIR /app/beast-hunter

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

ADD cmd ./cmd
ADD config ./config
ADD internal ./internal

RUN CGO_ENABLED=0 go build cmd/beast-hunter/main.go


FROM alpine:latest as deploy

WORKDIR /app/beast-hunter

COPY --from=build /app/beast-hunter/main ./

ADD config ./config

RUN chmod +x ./main

CMD [ "./main" ]