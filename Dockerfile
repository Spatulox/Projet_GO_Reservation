# Utiliser l'image Go officielle comme base
FROM golang:1.19-alpine

# Définir le répertoire de travail
WORKDIR /app

# Copier le fichier go.mod
COPY . .

# Installer les dépendances Go
RUN go mod download

# Copier le code source de l'application Go
# COPY ./app .

# Compiler l'application Go
# RUN go build -o main
# Compile error :/

# Exposer le port 8080 pour l'application Go
EXPOSE 8280

# Définir la commande de démarrage
# CMD ["./app/main.go"]
