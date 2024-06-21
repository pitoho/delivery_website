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

CREATE TABLE Order_Dish (
  id_order INT REFERENCES Order_Storage (id_order),
  id_dish INT REFERENCES Dish_Storage (id_dish),
  PRIMARY KEY (id_order, id_dish)
);
-- Создание таблицы для хранения заказов

CREATE TABLE Order_Storage (
  id_order SERIAL PRIMARY KEY,
  customer_id INT REFERENCES User_Data (id_user),
  order_time TIMESTAMP WITH TIME ZONE DEFAULT timezone('UTC+3', NOW()), 
  delivery_address_id INT REFERENCES Delivery_Addresses (id_address),
  total_price INT,
  order_status VARCHAR(255)
);

-- Создание таблицы для хранения сессий
CREATE TABLE Sessions (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES User_Data (id_user),
  token VARCHAR(255)
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

--Процедура на добавление данных в Delivery Address

CREATE OR REPLACE PROCEDURE check_delivery_address_uniqueness(
  IN street1 VARCHAR(255),
  IN house1 INT,
  IN corpus_building1 INT,
  IN flat1 INT,
  OUT existing_id INT
) AS $$
BEGIN
  SELECT id_address INTO existing_id 
  FROM Delivery_Addresses d
  WHERE d.street = street1 AND d.house = house1 AND d.corpus_building = corpus_building1 AND d.flat = flat1;
  IF existing_id IS NULL THEN
    INSERT INTO Delivery_Addresses (street, house, corpus_building, flat)
    VALUES (street1, house1, corpus_building1, flat1)
    RETURNING id_address INTO existing_id;
  END IF;
END;
$$ LANGUAGE plpgsql;

CALL check_delivery_address_uniqueness(

--Процедура на выбор блюд с тегами

CREATE OR REPLACE FUNCTION get_dish_with_tags()
RETURNS TABLE (
  id_dish INT,
  dish_name VARCHAR(255),
  dish_image_path VARCHAR(255),
  price INT,
  tag VARCHAR(255)
) AS $$
BEGIN
  RETURN QUERY 
  SELECT 
    ds.id_dish,
    ds.dish_name,
    ds.dish_image_path,
    ds.price,
    t.tag
  FROM Dish_Storage ds
  JOIN Tags t ON ds.tags_id = t.id_tag;
END;
$$ LANGUAGE plpgsql;

--Процедура выборки данных из User_Data по полю email

CREATE OR REPLACE PROCEDURE get_user_id_by_email(
  IN email VARCHAR(255),
  OUT user_id INT
) AS $$
BEGIN
  SELECT id_user INTO user_id
  FROM User_Data
  WHERE user_email = email;
END;
$$ LANGUAGE plpgsql;


--Процедура сравнения хеша

CREATE OR REPLACE PROCEDURE check_token_and_create_session(
  IN token VARCHAR(255)
) AS $$
DECLARE
  user_record RECORD;
  hashed_email VARCHAR(32);
BEGIN
  FOR user_record IN SELECT * FROM User_Data LOOP
    SELECT MD5(user_record.user_email) INTO hashed_email;
    IF hashed_email = token THEN
      INSERT INTO Sessions (user_id, token)
      VALUES (user_record.id_user, token);
      EXIT;
    END IF;
  END LOOP;
END;
$$ LANGUAGE plpgsql;

CALL check_token_and_create_session('7e36d74c46009a0ff436f4cb4584fd94')
	SELECT * FROM Sessions

--Процедура на добавление данных в order_storage
	
CREATE OR REPLACE PROCEDURE add_order(
  IN customer_id INT,
  IN delivery_address_id INT,
  IN total_price INT,
  IN order_status VARCHAR(255),
  OUT new_order_id INT
) AS $$
BEGIN
  INSERT INTO Order_Storage (customer_id, delivery_address_id, total_price, order_status)
  VALUES (customer_id, delivery_address_id, total_price, order_status)
  RETURNING id_order INTO new_order_id;
END;
$$ LANGUAGE plpgsql;

--Процедура на добавление данных в order_dish

CREATE OR REPLACE PROCEDURE add_order_dish(
    IN id_order INT,
    IN id_dish INT
) AS $$
BEGIN
    INSERT INTO Order_Dish (id_order, id_dish)
    VALUES (id_order, id_dish);
END;
$$ LANGUAGE plpgsql;
------------------------------------------------


