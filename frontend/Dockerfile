# syntax=docker/dockerfile:1

# Etapa 1: build
FROM node:20-alpine AS build

WORKDIR /app

# Copiar package.json y package-lock.json (o package.json y package.json solo)
COPY package*.json ./

# Instalar dependencias
RUN npm install

# Copiar el resto del código
COPY . .

# Construir la app para producción
RUN npm run build

# Etapa 2: Nginx para servir archivos estáticos
FROM nginx:alpine

# Copiar el build al directorio público de Nginx
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80
