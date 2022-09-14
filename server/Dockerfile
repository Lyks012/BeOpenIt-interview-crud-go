FROM golang:1.18.1-alpine3.15
WORKDIR /app
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN chmod +rx wait-for entrypoint.sh

ENTRYPOINT [ "sh", "entrypoint.sh" ]
# EXPOSE 8081
# CMD go run .