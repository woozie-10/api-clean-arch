FROM golang
WORKDIR /app
COPY go.mod go.sum ./
EXPOSE 8081
RUN go mod download
COPY . ./
RUN go build -o api-clean-arch ./cmd
CMD ["./api-clean-arch"]