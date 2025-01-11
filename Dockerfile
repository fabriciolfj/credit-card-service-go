FROM golang:1.23-alpine AS builder

# Adiciona git e dependências de build
RUN apk add --no-cache git

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Download das dependências
RUN go mod download

# Copia todo o código fonte e arquivos de configuração
COPY . .

# Compila a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

# Final stage
FROM alpine:latest

# Adiciona certificados CA e timezone
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copia o binário e arquivos de configuração do builder
COPY --from=builder /app/main .
COPY --from=builder /app/config.properties .

# Configura o entrypoint
CMD ["./main"]