
## install
0. Download and unpack the source data to the data_source folder. [Link to the source data for downloading](https://disk.yandex.ru/d/Uh54JX7qJV0TEg)
( I don 't remember where I took this data)
1. Start the database . There is a ready-made case to run in the docker ```make start_db```
2. [Install golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
3. Run DB migrations ```make migrate```
4. ```make build```
5. ```./bin/import data_sorce/file.txt```
6. ```make dbshell```
7. Enjoy

## Установка
0. Скачать и распаковать исходные данные в папку data_source.
[Ссылка на исходные данные для скачивания](https://disk.yandex.ru/d/Uh54JX7qJV0TEg) 
( где я эти данные взял в упор не помню )
1. Запустить базу данных . Есть готовый вариант для запуска в докере ```make start_db```
2. [Поставить golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
3. Запустить миграции БД ```make migrate```
4. ```make build```
5. ```./bin/import data_sorce/file.txt```
6. ```make dbshell```
7. Enjoy

