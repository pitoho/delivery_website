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

insert into DISH_STORAGE( 
  dish_name ,
  dish_image_path ,
  price ,
  tags_id 
) values (
	
)

insert into Tags (
  tag 
) values 
	('garnir_fried'),
	('garnir_boiled'),
	('base_sea'),
	('base_meat'),
	('base_vegan'),
	('drink_tea'),
	('drink_alc'),
	('drink_gas')

select * from Tags

CREATE OR REPLACE FUNCTION fill_dish_storage() RETURNS void AS $$
BEGIN
    -- Insert values for garnir_fried
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('French Fries', 'https://img.freepik.com/free-photo/crispy-french-fries-with-ketchup-mayonnaise_1150-26588.jpg?t=st=1718892793~exp=1718896393~hmac=3d0dc33e249f5ae10c3ccdb811ab58d9fc3535eb8983510faea53c45ab5407a6&w=740', 150, 1),
    ('Onion Rings', 'https://img.freepik.com/free-photo/high-angle-delicious-fast-food-drink_23-2149235974.jpg?t=st=1718893242~exp=1718896842~hmac=188cc2fbf57b731076fac061583064a97a614eeb1cb7b424457f7cce1218b5f9&w=740', 130, 1),
    ('Fried Zucchini', 'https://img.freepik.com/free-photo/baked-zucchini-sticks-with-cheese-bread-crumbs_2829-10864.jpg?t=st=1718893262~exp=1718896862~hmac=138cfc29a053d6580fd3e236722481cef1e4399426a5ac8e3bc59703638b2ab0&w=740', 140, 1),
    ('Fried Mushrooms', 'https://img.freepik.com/free-photo/top-view-delicious-cooked-mushrooms-with-greens-dark-background-dish-dinner-meal-wild-plant-food_140725-96261.jpg?t=st=1718893285~exp=1718896885~hmac=cbc0b679d592d448480e4279052c34cf47fa0af10ac7954b45ec444a78a87dd2&w=740', 160, 1);
    
    -- Insert values for garnir_boiled
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Boiled Potatoes', 'https://img.freepik.com/free-photo/potatoes-meal-top-view-composition_23-2148619111.jpg?t=st=1718893318~exp=1718896918~hmac=0a617337e1c6e39fdb80e495d288126f419f87575d40f091d6b12181f3d3a296&w=740', 120, 2),
    ('Boiled Carrots', 'https://img.freepik.com/free-photo/whole-slice-carrots-with-parsley-bowl-blue-table_114579-89875.jpg?t=st=1718893344~exp=1718896944~hmac=b236dcfa92523302a547087b4bff70a4a90a2900563bb6ab5b682a71c7a6188e&w=740', 110, 2),
    ('Boiled Rice', 'https://img.freepik.com/free-photo/elegant-minimalistic-rice-bowl_23-2149483989.jpg?t=st=1718893366~exp=1718896966~hmac=7ca7869f174cbf02abfa958ab51db76eb9054a56f5be926a0a63b115d31e1f7a&w=740', 130, 2),
    ('Boiled Broccoli', 'https://img.freepik.com/free-photo/fresh-raw-green-broccoli-frying-pan-yellow-surface_176474-447.jpg?t=st=1718893395~exp=1718896995~hmac=7979faa7bab73e32a4ae622259b2d0c246425ee853407b7b171ead6380281473&w=740', 140, 2);
    
    -- Insert values for base_sea
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Grilled Salmon', 'https://example.com/images/grilled_salmon.jpg', 450, 3),
    ('Shrimp Cocktail', 'https://example.com/images/shrimp_cocktail.jpg', 400, 3),
    ('Fish Tacos', 'https://example.com/images/fish_tacos.jpg', 350, 3),
    ('Lobster Roll', 'https://example.com/images/lobster_roll.jpg', 500, 3);
    
    -- Insert values for base_meat
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Beef Steak', 'https://example.com/images/beef_steak.jpg', 550, 4),
    ('Pork Chops', 'https://example.com/images/pork_chops.jpg', 500, 4),
    ('Chicken Breast', 'https://example.com/images/chicken_breast.jpg', 350, 4),
    ('Lamb Ribs', 'https://example.com/images/lamb_ribs.jpg', 600, 4);
    
    -- Insert values for base_vegan
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Vegan Burger', 'https://example.com/images/vegan_burger.jpg', 300, 5),
    ('Grilled Tofu', 'https://example.com/images/grilled_tofu.jpg', 250, 5),
    ('Vegan Pizza', 'https://example.com/images/vegan_pizza.jpg', 350, 5),
    ('Quinoa Salad', 'https://example.com/images/quinoa_salad.jpg', 200, 5);
    
    -- Insert values for drink_tea
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Green Tea', 'https://example.com/images/green_tea.jpg', 80, 6),
    ('Black Tea', 'https://example.com/images/black_tea.jpg', 70, 6),
    ('Herbal Tea', 'https://example.com/images/herbal_tea.jpg', 90, 6),
    ('Chai Tea', 'https://example.com/images/chai_tea.jpg', 100, 6);
    
    -- Insert values for drink_alc
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Red Wine', 'https://example.com/images/red_wine.jpg', 300, 7),
    ('Beer', 'https://example.com/images/beer.jpg', 150, 7),
    ('Whiskey', 'https://example.com/images/whiskey.jpg', 400, 7),
    ('Vodka', 'https://example.com/images/vodka.jpg', 350, 7);
    
    -- Insert values for drink_gas
    INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES 
    ('Coca Cola', 'https://example.com/images/coca_cola.jpg', 100, 8),
    ('Sprite', 'https://example.com/images/sprite.jpg', 90, 8),
    ('Fanta', 'https://example.com/images/fanta.jpg', 90, 8),
    ('Pepsi', 'https://example.com/images/pepsi.jpg', 100, 8);
END;
$$ LANGUAGE plpgsql;

-- Call the function to fill the table
SELECT fill_dish_storage();

