## Установка
1. Скачать и распаковать исходные данные в папку data_source.
[Ссылка на исходные данные для скачивания](https://disk.yandex.ru/d/Uh54JX7qJV0TEg) 
( где я эти данные взял в упор не помню )
1. Запустить базу данных . Есть готовый для запуска в докере ```make start_db```
2. [Поставить golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
3. Запустить миграции БД ```make migrate```
   
