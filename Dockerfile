FROM golang:1.23.5

# Enviroment variable
WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

#Copying files to work directory
COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

# Run and expose the server on port 3000
EXPOSE 3000