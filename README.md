# bookservice
GRCP сервис на go + клиент  + mysql on docker
1)запустить в терминале Makefile командой : Makefile
2)Запустите в терминале bs_server командой : ./bs_server
3)Перейдите в другом терминале в каталог databases/mysql/ командой : cd databases/mysql
4)Запустите docker compose командой : docker-compose.yml
5)Запустите в другом терминале bs_client : ./client  (один из методов) параметры("Имя автора\Название книги",Like )
    Методы:
        1)searchbyauthor          2)searchbybook
        Параметр Like активен      Параметр like отключен

Пример запуска с методом и параметрами : ./client  searchbybook "Преступление и наказание" 1
