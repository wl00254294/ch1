FROM golang:alpine AS builder
RUN sed -i 's/https/http/' /etc/apk/repositories
#RUN apk add --no-cache curl
RUN apk update && apk add --no-cache git
WORKDIR /eric/go/ch1/
COPY . .
RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/main
#RUN go build -o /go/bin/main
#EXPOSE 8080 8080
#ENTRYPOINT /go/bin/main


FROM scratch
COPY --from=builder /eric/go/ch1/config.json /eric/go/ch1/config.json
COPY --from=builder /go/bin/main /go/bin/main
EXPOSE 8080 8080
CMD ["/go/bin/main"]

