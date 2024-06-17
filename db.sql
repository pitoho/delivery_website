-- Создание таблицы для хранения пользовательских данных
CREATE TABLE User_Data (
  id_user SERIAL PRIMARY KEY,
  user_name VARCHAR(255),
  last_name VARCHAR(255),
  phone_number VARCHAR(20),
  user_email VARCHAR(255),
  user_password VARCHAR(255)
);
-- Создание таблицы для хранения тегов
CREATE TABLE Tags (
  id_tag SERIAL PRIMARY KEY,
  tag VARCHAR(255)
);

-- Создание таблицы для хранения адресов доставки
CREATE TABLE Delivery_Addresses (
  id_address SERIAL PRIMARY KEY,
  street VARCHAR(255),
  house INT,
  corpus_building INT,
  flat INT
);

-- Создание таблицы для хранения информации о блюдах
CREATE TABLE Dish_Storage (
  id_dish SERIAL PRIMARY KEY,
  dish_name VARCHAR(255),
  dish_image_path VARCHAR(255),
  price INT,
  tags_id INT REFERENCES Tags (id_tag)
);

-- Создание таблицы для хранения заказов
CREATE TABLE Order_Storage (
  id_order SERIAL PRIMARY KEY,
  customer_id INT REFERENCES User_Data (id_user),
  dish_id INT REFERENCES Dish_Storage (id_dish),
  order_time TIME,
  delivery_address_id INT REFERENCES Delivery_Addresses (id_address),
  total_price INT,
  order_status VARCHAR(255)
);

------------------------------------------------
--Процедура на добавление данных в User Data

CREATE OR REPLACE PROCEDURE add_user(
    user_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_number VARCHAR(20),
    user_email VARCHAR(255),
    user_password VARCHAR(255)
)
AS $$
BEGIN
    INSERT INTO User_Data (user_name, last_name, phone_number, user_email, user_password)
    VALUES (user_name, last_name, phone_number, user_email, user_password);
END;
$$ LANGUAGE plpgsql;

--Процедура на добавлени е данных в Delivery

CREATE OR REPLACE PROCEDURE add_order_with_delivery_address(
    customer_name VARCHAR(255),
    customer_last_name VARCHAR(255),
    customer_phone_number VARCHAR(20),
    customer_email VARCHAR(255),
    customer_password VARCHAR(255),
    dish_id INT,
    order_time TIME,
    street VARCHAR(255),
    house INT,
    corpus_building INT,
    flat INT,
    total_price INT,
    order_status VARCHAR(255)
)
AS $$
DECLARE
    last_delivery_address_id INT;
    last_customer_id INT;
BEGIN
    -- Добавляем нового пользователя и получаем его id
    INSERT INTO User_Data (user_name, last_name, phone_number, email, password)
    VALUES (customer_name, customer_last_name, customer_phone_number, customer_email, customer_password)
    RETURNING id_user INTO last_customer_id;

    -- Добавляем новый адрес доставки и получаем его id
    INSERT INTO Delivery_Addresses (street, house, corpus_building, flat)
    VALUES (street, house, corpus_building, flat)
    RETURNING id_address INTO last_delivery_address_id;

    -- Вставляем новый заказ с полученными id_address и id_user
    INSERT INTO Order_Storage (customer_id, dish_id, order_time, delivery_address_id, total_price, order_status)
    VALUES (last_customer_id, dish_id, order_time, last_delivery_address_id, total_price, order_status);
END;
$$ LANGUAGE plpgsql;


