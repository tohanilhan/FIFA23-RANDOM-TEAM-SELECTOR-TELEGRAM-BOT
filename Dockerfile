# Dockerfile for 'klearis-cron-manager'
# build stage
# define base image
FROM golang:1.20.7-alpine AS build


# create work directory
RUN mkdir /prog

# switch to work directory
WORKDIR /prog

# copy all files
ADD . .

# download dependencies
RUN go mod download

# build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o prog .

# build stage #1
FROM alpine:3.17 AS deploy

# install timezone
RUN apk add --no-cache tzdata

# create work directory
RUN mkdir /prog

# switch to work directory
WORKDIR /prog

# copy application artifacts to current directory
COPY --from=build /prog/prog .

# run application
ENTRYPOINT ["/prog/prog","6041537593:AAEVwq-4ntBXO_bWh9zyyRhfyJuuiAEt9Os"]
