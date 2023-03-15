FROM golang:1.18-alpine
WORKDIR /app
EXPOSE 1080

COPY . .
ENV GO111MODULE=on
RUN go build -o hello-app ./cmd

CMD [ "./hello-app" ]