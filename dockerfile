# Use a imagem oficial do Golang como base
FROM golang:1.20.2

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho atual
COPY go.mod go.sum ./

# Baixe todas as dependências
RUN go mod download

# Copie o código-fonte para o diretório de trabalho atual
COPY . .

# Compile o aplicativo
RUN go build -o main ./src

# Exponha a porta que o aplicativo irá rodar
EXPOSE 3330

# Comando para iniciar o aplicativo
CMD ["./main"]