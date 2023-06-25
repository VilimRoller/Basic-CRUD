FROM golang:1.19

#Set copy destination
WORKDIR /app

#Download dependencies
COPY go.mod go.sum ./
RUN go mod download

#Copy the source code
COPY . .

#Build
RUN go build -o /Basic-CRUD

EXPOSE 8080

#Run
CMD ["/Basic-CRUD"]
