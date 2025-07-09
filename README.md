# Тестовое задание asys.one

утилита для потоковой загрузки больших файлов без выгрузки в память

## Описание
 - Загрузка файла на HTTP сервер через POST
 - Отображение прогресс-бара во время передачи
 - Чтение конфигурации из YAML файла
 - Консольный интерфейс

## Библиотеки
 - io
 - os
 - mime/multipart
 - http
 - github.com/schollz/progressbar/v3
 - github.com/spf13/viper
 - github.com/spf13/cobra

## Установка
```
git clone https://github.com/altaastro/datatool.git
cd datatool
go build -o datatool ./cmd/main.go
```

## Конфигурация

Файл конфигурации:configs/config.yaml

- file_path — полный путь к файлу для загрузки
- server_url — адрес эндпойнта, принимающего файл

## Использованиие
```
./datatool upload --config=configs/config.yaml
```
