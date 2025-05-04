-- Создаем базу данных
CREATE DATABASE postgres;

-- Назначаем существующего пользователя postgres владельцем базы данных
GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;