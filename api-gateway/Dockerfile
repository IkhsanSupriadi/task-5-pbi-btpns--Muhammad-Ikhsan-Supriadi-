FROM golang:1.20.4

WORKDIR /app 

COPY . ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-gateway

EXPOSE 3000

CMD [ "/api-gateway" ]