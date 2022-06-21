# Base image for my docker
FROM golang:1.17-alpine

ENV GO111MODULE=on


RUN mkdir calendarAPI
RUN mkdir calendarAPI/businessLayer
RUN mkdir calendarAPI/model
RUN mkdir calendarAPI/module
RUN mkdir calendarAPI/routes

COPY go.mod calendarAPI/
COPY go.sum calendarAPI/

COPY businessLayer calendarAPI/businessLayer
COPY model calendarAPI/model
COPY module calendarAPI/module
COPY routes calendarAPI/routes
COPY main.go calendarAPI/

WORKDIR calendarAPI

RUN go mod download

RUN go build -o main .

EXPOSE 80

CMD ["./main"]