# bookservice
#### GRCP сервис на go + клиент  + mysql on docker
## Для сборки проекта в терминале используйте команду : *./make*

Пример сборки проекта  :


![image](https://user-images.githubusercontent.com/109175607/178629154-8b171825-732b-473e-8ef0-203ecdfae589.png)


## Запустите сервер в терминале используя команду  : *./bs_server*

Пример запуска сервера  :


 ![image](https://user-images.githubusercontent.com/109175607/178629172-92912e3d-1e81-4a01-9b7f-5df25cdc9d31.png)


## Для запуска докер контейнера с mysql перейдите в каталог  */databases/mysql*

## Запустите docker compose в отдельном терминале командой:

*sudo docker-compose up*

Пример запуска контейнера :


![image](https://user-images.githubusercontent.com/109175607/178629186-745fd377-391e-4eeb-bd61-dae108b549dc.png)


Контейнер успешно развернут :


![image](https://user-images.githubusercontent.com/109175607/178629220-599a490b-8ed5-4894-ab22-120113c53148.png)


## Запустите в другом терминале клиент командой :
 
*./client  (один из методов) параметры("Имя автора\Название книги",Like )*

### Методы :

1 *searchbyauthor*  Параметр Like используется

2 *searchbybook*  Параметр like неиспользуется, пишите точные названия              

Пример запуска с методом и параметрами : 

*./client  searchbybook "Преступление и наказание" 1*

Пример запуск клиента :

![image](https://user-images.githubusercontent.com/109175607/178629245-46fa946c-a68a-475a-8bca-ed3a2adf7fb6.png)

