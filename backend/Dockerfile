#syntax=docker/dockerfile:1

FROM golang:1.24.1-alpine
WORKDIR /app

#copiar solo los archivos de dependecias primero 
COPY go.mod ./
COPY go.sum ./

#Descargar dependencias
COPY . .
RUN go mod tidy
RUN go mod download
RUN go build -o /backend

EXPOSE 8080
CMD ["sh", "-c", "sleep 10 && /backend"]