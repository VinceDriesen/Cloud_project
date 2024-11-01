# Gebruik de officiële Go image als builder
FROM golang:1.23.2-alpine AS builder

# Stel de werkdirectory in
WORKDIR /app

# Installeer wget en git voor het downloaden van de migratietool
RUN apk add --no-cache wget git

# Download en installeer de migrate tool
RUN wget https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/ && \
    rm migrate.linux-amd64.tar.gz

# Kopieer go.mod en go.sum om afhankelijkheden te downloaden
COPY go.mod go.sum ./
RUN go mod download

# Kopieer de rest van de broncode
COPY . ./

# Bouw de applicatie
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Expose de poort
EXPOSE 8082

# Start de server
CMD ["/server"]
