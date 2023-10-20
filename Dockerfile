FROM golang:1.21-alpine AS BASE
WORKDIR /app
COPY *.sum *.mod ./
RUN go mod tidy
COPY . .
RUN go build -o go-cep cmd/main.go 

FROM alpine
WORKDIR /app
COPY --from=BASE /app/go-cep /app/go-cep
EXPOSE 8080
RUN apk --no-cache add ca-certificates
CMD [ "/app/go-cep" ]