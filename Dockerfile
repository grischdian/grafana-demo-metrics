# Verwende das offizielle Go-Image als Basis
FROM golang:1.18-alpine as builder

# Setze das Arbeitsverzeichnis im Container
WORKDIR /app

# Kopiere den Quellcode in den Container
COPY . .

# Kompiliere die Anwendung
RUN go env -w GO111MODULE=off && \
  go build -o /metrics-server main.go && \
  go build -o /metrics-server2 main2.go

# Starte einen neuen, leichten Container
FROM alpine:latest  

WORKDIR /

# Kopiere die kompilierte Anwendung aus dem vorherigen Schritt
COPY --from=builder /metrics-server /metrics-server
COPY --from=builder /metrics-server2 /metrics-server2

# Exponiere den Port, den die Anwendung verwendet
EXPOSE 8080

# Startbefehl der Anwendung
CMD ["/metrics-server"]
