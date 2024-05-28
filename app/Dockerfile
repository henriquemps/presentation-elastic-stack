# Etapa de construção
FROM golang:1.20 as builder

# Diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod download

RUN go mod tidy

# Copie o restante do código fonte
COPY .. .

# Compile o aplicativo Go
RUN go build -o app

# Porta exposta pelo aplicativo
EXPOSE 8080

# Comando para rodar o aplicativo
CMD ["./app"]
