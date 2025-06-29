# syntax=docker/dockerfile:1

# Etapa 1: build
FROM node:20-alpine AS build

WORKDIR /app

# Copiar package.json y package-lock.json (o solo package.json)
COPY package*.json ./

# Instalar dependencias
RUN npm install

# Copiar el resto del código
COPY . .

# Construir la app para producción
RUN npm run build

# Etapa 2: Nginx para servir archivos estáticos
FROM nginx:alpine

# BORRAR configuración default de Nginx
RUN rm /etc/nginx/conf.d/default.conf

# Escribir configuración custom directamente en el Dockerfile
RUN echo "server { \
    listen 80; \
    server_name localhost; \
    root /usr/share/nginx/html; \
    index index.html; \
    location / { \
    try_files \$uri /index.html; \
    } \
    }" > /etc/nginx/conf.d/default.conf


# Copiar el build al directorio público de Nginx
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80
