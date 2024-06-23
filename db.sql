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

-- CALL check_delivery_address_uniqueness(

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

--Функция на выбор данных из order_storage по user_id

CREATE OR REPLACE FUNCTION get_user_orders(
  user_id INT
)
RETURNS TABLE(
  got_id_order INT,
  got_order_time TIMESTAMP WITH TIME ZONE,
  got_total_price INT,
  got_order_status VARCHAR(255)
) AS $$
BEGIN
  RETURN QUERY SELECT id_order, order_time, total_price, order_status 
  FROM Order_Storage WHERE customer_id = user_id;
END;
$$ LANGUAGE plpgsql;

select * from Order_Storage 
select * from get_user_orders(18);
------------------------------------------------

select * from tags
	
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


--Insert values for garnir_fried
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('French Fries', 'https://img.freepik.com/free-photo/crispy-french-fries-with-ketchup-mayonnaise_1150-26588.jpg?t=st=1718892793~exp=1718896393~hmac=3d0dc33e249f5ae10c3ccdb811ab58d9fc3535eb8983510faea53c45ab5407a6&w=740', 150, 1),
('Onion Rings', 'https://img.freepik.com/free-photo/high-angle-delicious-fast-food-drink_23-2149235974.jpg?t=st=1718893242~exp=1718896842~hmac=188cc2fbf57b731076fac061583064a97a614eeb1cb7b424457f7cce1218b5f9&w=740', 130, 1),
('Fried Zucchini', 'https://img.freepik.com/free-photo/baked-zucchini-sticks-with-cheese-bread-crumbs_2829-10864.jpg?t=st=1718893262~exp=1718896862~hmac=138cfc29a053d6580fd3e236722481cef1e4399426a5ac8e3bc59703638b2ab0&w=740', 140, 1),
('Fried Mushrooms', 'https://img.freepik.com/free-photo/top-view-delicious-cooked-mushrooms-with-greens-dark-background-dish-dinner-meal-wild-plant-food_140725-96261.jpg?t=st=1718893285~exp=1718896885~hmac=cbc0b679d592d448480e4279052c34cf47fa0af10ac7954b45ec444a78a87dd2&w=740', 160, 1);

--Insert values for garnir_boiled
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Boiled Potatoes', 'https://img.freepik.com/free-photo/potatoes-meal-top-view-composition_23-2148619111.jpg?t=st=1718893318~exp=1718896918~hmac=0a617337e1c6e39fdb80e495d288126f419f87575d40f091d6b12181f3d3a296&w=740', 120, 2),
('Boiled Carrots', 'https://img.freepik.com/free-photo/whole-slice-carrots-with-parsley-bowl-blue-table_114579-89875.jpg?t=st=1718893344~exp=1718896944~hmac=b236dcfa92523302a547087b4bff70a4a90a2900563bb6ab5b682a71c7a6188e&w=740', 110, 2),
('Boiled Rice', 'https://img.freepik.com/free-photo/elegant-minimalistic-rice-bowl_23-2149483989.jpg?t=st=1718893366~exp=1718896966~hmac=7ca7869f174cbf02abfa958ab51db76eb9054a56f5be926a0a63b115d31e1f7a&w=740', 130, 2),
('Boiled Broccoli', 'https://img.freepik.com/free-photo/fresh-raw-green-broccoli-frying-pan-yellow-surface_176474-447.jpg?t=st=1718893395~exp=1718896995~hmac=7979faa7bab73e32a4ae622259b2d0c246425ee853407b7b171ead6380281473&w=740', 140, 2);

--Insert values for base_sea
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Grilled Salmon', 'https://img.freepik.com/free-photo/top-view-delicious-meal-black-tray-dark-wooden-background-with-decorations_176474-3955.jpg?t=st=1719087512~exp=1719091112~hmac=4c9a07f909833437a261a991e5ff9c22cbe340352e36133c49a26285d7468bc9&w=740', 450, 3),
('Shrimp Cocktail', 'https://img.freepik.com/free-photo/tails-shrimps-with-fresh-lemon-rosemary-plate_2829-14150.jpg?t=st=1719087542~exp=1719091142~hmac=b4c26d4d801a262c2811c0a47c67d45791fcf59613b31d105f68be0f4118b8ca&w=740', 400, 3),
('Fish Tacos', 'https://img.freepik.com/free-photo/top-view-mexican-food-concept_23-2148629367.jpg?t=st=1719087559~exp=1719091159~hmac=be8e78f3b97caec0ab7dbeb8d818763c2d7d9e11a9e9093a8eba84333448e5fa&w=740', 350, 3),
('Lobster Roll', 'https://img.freepik.com/free-photo/kegs-pancakes-with-red-fish_2829-14009.jpg?t=st=1719087582~exp=1719091182~hmac=ccc42b2dc0a508bd4d22923e001fe17d17847c84119d152edd56ccee19d7c502&w=740', 500, 3);

-- Insert values for base_meat
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Beef Steak', 'https://img.freepik.com/free-photo/top-view-grilled-chop-meat-with-vegetable-salad-sauce-blackboard_141793-4036.jpg?t=st=1719087602~exp=1719091202~hmac=c14e77ee4d1f891ceaa04f0030c02cf81dde16267259316ec306a3582939fa49&w=740', 550, 4),
('Pork Chops', 'https://img.freepik.com/free-photo/duck-breast-steak_1203-2299.jpg?t=st=1719087627~exp=1719091227~hmac=76a855d3aa3354e6093e68f9a6736d4e59117a1a27f42be1bd8678c30e7db442&w=740', 500, 4),
('Chicken Breast', 'https://img.freepik.com/premium-photo/grilled-chicken-breasts-vegetables_2829-8830.jpg?w=740', 350, 4),
('Lamb Ribs', 'https://img.freepik.com/free-photo/wooden-board-with-tasty-cooked-meat_23-2148599800.jpg?t=st=1719087672~exp=1719091272~hmac=26c2b0cb961e681e8ef3bc07cf9d466c22fc8a207576b64f4975dcda81dbcdb7&w=740', 600,4);

--Insert values for base_vegan
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Vegan Burger', 'https://img.freepik.com/free-photo/front-view-vegetarian-burger-counter-with-tomatoes_23-2148784525.jpg?t=st=1719087702~exp=1719091302~hmac=47de66ef3f61c0be05f8db8e51bf9782bd1c0e151cc7edc7ee3b7806bb6ba50c&w=996', 300, 5),
('Grilled Tofu', 'https://img.freepik.com/free-photo/spring-rolls_74190-1407.jpg?t=st=1719087725~exp=1719091325~hmac=6a6ba4b7cfea46307a8247219e31de72ebc8cf2fad54767821c9a5bf1b07426a&w=740', 250, 5),
('Vegan Pizza', 'https://img.freepik.com/free-photo/arugula-pizza-with-white-background_23-2148574291.jpg?t=st=1719087759~exp=1719091359~hmac=f3a0f30b8d7bdfbe61249479dc2673a52858b7534ea3f507263f118498f7b27b&w=740', 350, 5),
('Quinoa Salad', 'https://img.freepik.com/free-photo/tabbouleh-salad_2829-10886.jpg?t=st=1719087843~exp=1719091443~hmac=6cd8fd7fa75cc44c1d37583f77233a1aa3e3a16e37c027a0db050eb8abbb0a6d&w=740', 200, 5);

--Insert values for drink_tea
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Green Tea', 'https://img.freepik.com/free-photo/leaf-plate-wood-object-healthy-eating_1172-451.jpg?t=st=1719091653~exp=1719095253~hmac=8c10f7a32a9cb0e02a541324fd2bdb05288b76de224eee4472dd9a1d225e056c&w=740', 80, 6),
('Black Tea', 'https://img.freepik.com/free-photo/flat-lay-cup-tea-with-infuser-marble-background_23-2148316954.jpg?t=st=1719091675~exp=1719095275~hmac=19c86daa85b2d5900d3d9003db15c411584b23cb520f5002814a57cb57424d52&w=740', 70, 6),
('Herbal Tea', 'https://img.freepik.com/free-photo/top-view-assortment-dried-plants_23-2148799537.jpg?t=st=1719091697~exp=1719095297~hmac=23035b89d4cffa9271d5d2734e4dc8466e6964f346c941c9d17e9aac5a2ba7cc&w=740', 90, 6),
('Chai Tea', 'https://img.freepik.com/free-photo/brown-sugar-near-cup-tea_23-2147764928.jpg?t=st=1719091721~exp=1719095321~hmac=92ee9764bcd1f196a19e6d408badd94fd4239453e5ec3007c5db638c2b007060&w=740', 100, 6);

--Insert values for drink_alc
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Red Wine', 'https://img.freepik.com/free-photo/red-wine_144627-33208.jpg?t=st=1719091758~exp=1719095358~hmac=e36b0206bc0237f4fabc6d364dedd1cb0ae8c23559c51cae8891b44510cb67bb&w=996', 300, 7),
('Beer', 'https://img.freepik.com/free-photo/fresh-light-beer-mug_140725-17.jpg?t=st=1719091780~exp=1719095380~hmac=c57bdc486354abbf7215ab717fa9e7646ca80c485c29bc0c903ca95ae2adcdb3&w=900', 150, 7),
('Whiskey', 'https://img.freepik.com/free-photo/glass-whiskey-bourbon-only-with-ice_144627-42842.jpg?t=st=1719091800~exp=1719095400~hmac=ae8e61e453cc47ba4f2a485bfa48072cd0fff2b0b99955fa22871941205b5f2d&w=740', 400, 7),
('Vodka', 'https://img.freepik.com/free-photo/frozen-glasses-with-cold-alochol-drink_144627-19386.jpg?t=st=1719091821~exp=1719095421~hmac=9bc868fb1b464c85f596a044cd2ea9b744a9b3e494bac508f4365e3ceab82479&w=740', 350, 7);

-- Insert values for drink_gas
INSERT INTO Dish_Storage(dish_name, dish_image_path, price, tags_id) VALUES
('Coca Cola', 'https://img.freepik.com/free-photo/fresh-cola-drink-glass_144627-16201.jpg?t=st=1719091846~exp=1719095446~hmac=5c740fc8c0bcb6104f9808e417d8c5b84d88ee340f4d19ac693c0c677b08dc71&w=740', 100, 8),
('Sprite', 'https://img.freepik.com/free-vector/hand-drawn-fresh-michelada-illustration_23-2149212111.jpg?t=st=1719091877~exp=1719095477~hmac=a7071592b9dc42035cc99eb904fe5d54bab9c8663ede50350b72b0b9b791c3c3&w=740', 90, 8),
('Fanta', 'https://img.freepik.com/free-photo/view-abstract-fluid-monochrome-palette_23-2150635167.jpg?t=st=1719091893~exp=1719095493~hmac=0879ac29e290be9133c597ec7c4acab8401abf7e8198b8beae7ade2106e52f27&w=740', 90, 8),
('Pepsi', 'https://img.freepik.com/free-photo/front-view-energy-drink-can-dark-pink-background-color-water-soda-darkness-drink_140725-158055.jpg?t=st=1719091911~exp=1719095511~hmac=5a9574edcf20369e357d6496a8a0fc8774eb009e2e0ccc81829b9c750ab2a9e9&w=740', 100, 8);

delete from public.user_data

select * from public.user_data