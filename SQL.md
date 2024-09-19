-- Create tables

CREATE TABLE "users" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text,
    "email" text,
    "password" text,
    PRIMARY KEY ("id"),
    CONSTRAINT "uni_users_email" UNIQUE ("email")
);

CREATE TABLE "products" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text,
    "description" text,
    "image_url" text,
    "price" decimal,
    "stock" bigint,
    PRIMARY KEY ("id")
);

CREATE TABLE "categories" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "name" text,
    "description" text,
    PRIMARY KEY ("id")
);

CREATE TABLE "product_categories" (
    "product_id" bigint,
    "category_id" bigint
);

CREATE TABLE "cart_items" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" bigint,
    "product_id" bigint,
    "quantity" bigint,
    "added_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_cart_items_product" FOREIGN KEY ("product_id") REFERENCES "products"("id")
);

CREATE TABLE "orders" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" bigint,
    "total" decimal,
    "status" text,
    PRIMARY KEY ("id") 
);

CREATE TABLE "order_items" (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "order_id" bigint,
    "product_id" bigint,
    "quantity" bigint,
    "price" decimal,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_orders_order_items" FOREIGN KEY ("order_id") REFERENCES "orders"("id")
);



-- Insert dummy data
-- Insert Categories
INSERT INTO categories (name, description) VALUES
('Electronics', 'Gadgets and electronic devices'),
('Clothing', 'Apparel for men and women'),
('Books', 'Various genres of books'),
('Home Appliances', 'Appliances for everyday use'),
('Sports Equipment', 'Gear for various sports'),
('Toys', 'Fun toys for children'),
('Health & Beauty', 'Products for health and personal care'),
('Automotive', 'Accessories and parts for vehicles');

-- Insert Products for each Category
-- Electronics
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Smartphone', 'Latest smartphone with advanced features', 'https://example.com/smartphone.jpg', 450000, 30),
('Laptop', 'High-performance laptop for work and gaming', 'https://example.com/laptop.jpg', 500000, 20),
('Headphones', 'Noise-cancelling over-ear headphones', 'https://example.com/headphones.jpg', 250000, 50),
('Smartwatch', 'Stylish smartwatch with health tracking', 'https://example.com/smartwatch.jpg', 300000, 40),
('Bluetooth Speaker', 'Portable Bluetooth speaker with great sound', 'https://example.com/speaker.jpg', 150000, 60),
('Camera', 'Digital camera with high resolution', 'https://example.com/camera.jpg', 400000, 15),
('LED TV', '55 inch LED TV with stunning picture quality', 'https://example.com/tv.jpg', 450000, 10),
('Tablet', '10 inch tablet for browsing and entertainment', 'https://example.com/tablet.jpg', 350000, 25);

-- Clothing
INSERT INTO products (name, description, image_url, price, stock) VALUES
('T-shirt', 'Cotton t-shirt available in various colors', 'https://example.com/tshirt.jpg', 100000, 100),
('Jeans', 'Stylish blue jeans for casual wear', 'https://example.com/jeans.jpg', 200000, 50),
('Jacket', 'Warm jacket for colder weather', 'https://example.com/jacket.jpg', 350000, 30),
('Dress', 'Elegant dress for special occasions', 'https://example.com/dress.jpg', 400000, 20),
('Sneakers', 'Comfortable sneakers for everyday use', 'https://example.com/sneakers.jpg', 250000, 40),
('Scarf', 'Soft scarf to keep warm', 'https://example.com/scarf.jpg', 80000, 60),
('Hat', 'Fashionable hat for sunny days', 'https://example.com/hat.jpg', 120000, 55),
('Belt', 'Leather belt for men and women', 'https://example.com/belt.jpg', 150000, 70);

-- Books
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Novel', 'Bestselling fiction novel', 'https://example.com/novel.jpg', 150000, 100),
('Cookbook', 'Delicious recipes for home cooking', 'https://example.com/cookbook.jpg', 200000, 50),
('Self-help Book', 'Motivational book for personal growth', 'https://example.com/selfhelp.jpg', 100000, 75),
('Children\'s Book', 'Fun stories for kids', 'https://example.com/childrensbook.jpg', 90000, 80),
('Biography', 'Life story of a famous personality', 'https://example.com/biography.jpg', 250000, 40),
('Science Fiction', 'Futuristic science fiction novel', 'https://example.com/scifi.jpg', 180000, 60),
('Mystery', 'Thrilling mystery novel', 'https://example.com/mystery.jpg', 220000, 30),
('Travel Guide', 'Guide to amazing travel destinations', 'https://example.com/travelguide.jpg', 170000, 50);

-- Home Appliances
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Blender', 'Powerful blender for smoothies', 'https://example.com/blender.jpg', 250000, 20),
('Microwave', 'Compact microwave for quick meals', 'https://example.com/microwave.jpg', 400000, 15),
('Toaster', '2-slice toaster for breakfast', 'https://example.com/toaster.jpg', 150000, 25),
('Vacuum Cleaner', 'Efficient vacuum cleaner for home', 'https://example.com/vacuum.jpg', 350000, 10),
('Electric Kettle', 'Quick boiling electric kettle', 'https://example.com/kettle.jpg', 200000, 30),
('Rice Cooker', 'Automatic rice cooker for perfect rice', 'https://example.com/ricecooker.jpg', 300000, 18),
('Air Fryer', 'Healthier way to fry food', 'https://example.com/airfryer.jpg', 450000, 12),
('Food Processor', 'Versatile food processor for cooking', 'https://example.com/foodprocessor.jpg', 400000, 15);

-- Sports Equipment
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Tennis Racket', 'High-quality tennis racket', 'https://example.com/racket.jpg', 300000, 10),
('Soccer Ball', 'Official size soccer ball', 'https://example.com/soccerball.jpg', 150000, 50),
('Yoga Mat', 'Comfortable yoga mat for workouts', 'https://example.com/yogamat.jpg', 100000, 75),
('Dumbbells', 'Set of dumbbells for strength training', 'https://example.com/dumbbells.jpg', 250000, 30),
('Bicycle', 'Mountain bike for outdoor adventures', 'https://example.com/bicycle.jpg', 500000, 5),
('Treadmill', 'Electric treadmill for home workouts', 'https://example.com/treadmill.jpg', 400000, 8),
('Baseball Glove', 'Soft baseball glove for training', 'https://example.com/baseballglove.jpg', 200000, 25),
('Skipping Rope', 'Durable skipping rope for fitness', 'https://example.com/skippingrope.jpg', 50000, 100);

-- Toys
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Building Blocks', 'Colorful building blocks for kids', 'https://example.com/buildingblocks.jpg', 150000, 50),
('Doll', 'Cute doll for imaginative play', 'https://example.com/doll.jpg', 200000, 40),
('Action Figure', 'Collectible action figure', 'https://example.com/actionfigure.jpg', 250000, 30),
('Puzzle', 'Challenging puzzle for all ages', 'https://example.com/puzzle.jpg', 100000, 60),
('Board Game', 'Fun board game for family nights', 'https://example.com/boardgame.jpg', 200000, 25),
('Toy Car', 'Fast toy car for racing', 'https://example.com/toycar.jpg', 50000, 80),
('Stuffed Animal', 'Soft and cuddly stuffed animal', 'https://example.com/stuffedanimal.jpg', 150000, 50),
('Kite', 'Colorful kite for outdoor fun', 'https://example.com/kite.jpg', 80000, 70);

-- Health & Beauty
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Shampoo', 'Nourishing shampoo for healthy hair', 'https://example.com/shampoo.jpg', 100000, 100),
('Face Cream', 'Moisturizing cream for daily use', 'https://example.com/facecream.jpg', 150000, 80),
('Perfume', 'Elegant perfume for women', 'https://example.com/perfume.jpg', 300000, 40),
('Makeup Kit', 'Complete makeup kit for beauty enthusiasts', 'https://example.com/makeupkit.jpg', 400000, 25),
('Sunscreen', 'SPF 50 sunscreen for skin protection', 'https://example.com/sunscreen.jpg', 200000, 60),
('Toothpaste', 'Whitening toothpaste for bright smile', 'https://example.com/toothpaste.jpg', 50000, 100),
('Lip Balm', 'Moisturizing lip balm for dry lips', 'https://example.com/lipbalm.jpg', 30000, 90),
('Body Wash', 'Refreshing body wash for daily use', 'https://example.com/bodywash.jpg', 100000, 75);

-- Automotive
INSERT INTO products (name, description, image_url, price, stock) VALUES
('Car Wax', 'High-quality car wax for shine', 'https://example.com/carwax.jpg', 150000, 30),
('Tire Inflator', 'Portable tire inflator for emergencies', 'https://example.com/tireinflator.jpg', 200000, 25),
('Car Phone Mount', 'Secure phone mount for cars', 'https://example.com/phonemount.jpg', 100000, 50),
('Seat Covers', 'Stylish seat covers for cars', 'https://example.com/seatcovers.jpg', 300000, 15),
('Car Vacuum', 'Compact vacuum for car interiors', 'https://example.com/carvacuum.jpg', 250000, 20),
('Dashboard Camera', 'Dash cam for vehicle safety', 'https://example.com/dashboardcamera.jpg', 500000, 5),
('Emergency Kit', 'Car emergency kit with essential tools', 'https://example.com/emergencykit.jpg', 400000, 8),
('Floor Mats', 'Durable floor mats for cars', 'https://example.com/floormats.jpg', 150000, 30);


-- Insert Product-Category Mappings
-- Electronics
INSERT INTO product_categories (product_id, category_id) VALUES
(1, 1), -- Smartphone
(2, 1), -- Laptop
(3, 1), -- Headphones
(4, 1), -- Smartwatch
(5, 1), -- Bluetooth Speaker
(6, 1), -- Camera
(7, 1), -- LED TV
(8, 1); -- Tablet

-- Clothing
INSERT INTO product_categories (product_id, category_id) VALUES
(9, 2), -- T-shirt
(10, 2), -- Jeans
(11, 2), -- Jacket
(12, 2), -- Dress
(13, 2), -- Sneakers
(14, 2), -- Scarf
(15, 2), -- Hat
(16, 2); -- Belt

-- Books
INSERT INTO product_categories (product_id, category_id) VALUES
(17, 3), -- Novel
(18, 3), -- Cookbook
(19, 3), -- Self-help Book
(20, 3), -- Children's Book
(21, 3), -- Biography
(22, 3), -- Science Fiction
(23, 3), -- Mystery
(24, 3); -- Travel Guide

-- Home Appliances
INSERT INTO product_categories (product_id, category_id) VALUES
(25, 4), -- Blender
(26, 4), -- Microwave
(27, 4), -- Toaster
(28, 4), -- Vacuum Cleaner
(29, 4), -- Electric Kettle
(30, 4), -- Rice Cooker
(31, 4), -- Air Fryer
(32, 4); -- Food Processor

-- Sports Equipment
INSERT INTO product_categories (product_id, category_id) VALUES
(33, 5), -- Tennis Racket
(34, 5), -- Soccer Ball
(35, 5), -- Yoga Mat
(36, 5), -- Dumbbells
(37, 5), -- Bicycle
(38, 5), -- Treadmill
(39, 5), -- Baseball Glove
(40, 5); -- Skipping Rope

-- Toys
INSERT INTO product_categories (product_id, category_id) VALUES
(41, 6), -- Building Blocks
(42, 6), -- Doll
(43, 6), -- Action Figure
(44, 6), -- Puzzle
(45, 6), -- Board Game
(46, 6), -- Toy Car
(47, 6), -- Stuffed Animal
(48, 6); -- Kite

-- Health & Beauty
INSERT INTO product_categories (product_id, category_id) VALUES
(49, 7), -- Shampoo
(50, 7), -- Face Cream
(51, 7), -- Perfume
(52, 7), -- Makeup Kit
(53, 7), -- Sunscreen
(54, 7), -- Toothpaste
(55, 7), -- Lip Balm
(56, 7); -- Body Wash

-- Automotive
INSERT INTO product_categories (product_id, category_id) VALUES
(57, 8), -- Car Wax
(58, 8), -- Tire Inflator
(59, 8), -- Car Phone Mount
(60, 8), -- Seat Covers
(61, 8), -- Car Vacuum
(62, 8), -- Dashboard Camera
(63, 8), -- Emergency Kit
(64, 8); -- Floor Mats
