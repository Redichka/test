# Используем образ Golang
FROM golang:1.16

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Отключаем поддержку модулей Go
ENV GO111MODULE=off

# Копируем исходный код в контейнер
COPY . .

# Компилируем приложение
RUN go build -o app .

# Указываем порт, на котором будет работать приложение
EXPOSE 8000

# Запускаем приложение при старте контейнера
CMD ["./app"]
