# Курс по Го в ШАД
https://gitlab.com/slon/shad-go

### Решенные задачи вычеркиваю в [hometasks.md](./hometasks.md)



Для работы с кодом нужен go 1.18 или выше.

```sh
git clone https://gitlab.com/slon/shad-go.git
cd shad-go
```

Проект можно открыть в любой go IDE с поддержкой модулей.

## Информация

- [Программа курса](docs/syllabus.md)
- [Слайды](https://p.go.manytask.org/)
- [Как прислать патч](CONTRIBUTING.md)


2. Проверьте, что ваше решение проходит тесты локально.

```shell
cd shad-go
# Из корня репозитория.
go test ./sum/...
```
   
3. Проверьте, что код проходит линтер. Линтер нужно установить [по инструкции](https://github.com/golangci/golangci-lint#binary).

```shell
# Из корня репозитория.
golangci-lint run ./sum/...
```
