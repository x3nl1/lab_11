# Лабораторная работа №11

## Студент

ФИО: Агапов Степан Александрович,
Группа: 221331

## Вариант

9

## Выполненные задания

### Средний уровень

* 1: Dockerfile для Python
* 3: Dockerfile для Rust
* 9: Ограничение ресурсов контейнеров

### Повышенный уровень

* 1: Go static + scratch
* 3: CI/CD pipeline

## Структура проекта

* m_task_api/ — Go сервис
* python_client/ — Python клиент
* rust_service/ — Rust сервис
* docker-compose.yml
* .github/workflows/ci.yml

## Запуск

docker compose up --build

## CI/CD

При push в main:

* собираются образы
* отправляются в Docker Hub

## Технологии

* Go
* Python
* Rust
* Docker
* GitHub Actions
