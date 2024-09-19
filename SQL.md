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
-- Users
INSERT INTO users (name, email, password) VALUES
('John Doe', 'john@example.com', 'password123'),
('Jane Smith', 'jane@example.com', 'password456');

-- Categories
INSERT INTO categories (name, description) VALUES
('Electronics', 'Electronic devices and accessories'),
('Clothing', 'Apparel and fashion items'),
('Books', 'Physical and digital books');

-- Products
INSERT INTO products (name, description, price, stock_quantity) VALUES
('Smartphone', 'High-end smartphone with advanced features', 699.99, 50),
('Laptop', 'Powerful laptop for work and gaming', 1299.99, 30),
('T-shirt', 'Comfortable cotton t-shirt', 19.99, 100),
('Jeans', 'Classic blue jeans', 49.99, 75),
('Novel', 'Bestselling fiction novel', 14.99, 200);

-- Product Categories
INSERT INTO product_categories (product_id, category_id) VALUES
(1, 1), -- Smartphone in Electronics
(2, 1), -- Laptop in Electronics
(3, 2), -- T-shirt in Clothing
(4, 2), -- Jeans in Clothing
(5, 3); -- Novel in Books

-- Cart Items
INSERT INTO cart_items (user_id, product_id, quantity) VALUES
(1, 1, 1), -- John has a smartphone in his cart
(1, 3, 2), -- John has two t-shirts in his cart
(2, 2, 1), -- Jane has a laptop in her cart
(2, 5, 3); -- Jane has three novels in her cart

-- Orders
INSERT INTO orders (user_id, total_amount, status) VALUES
(1, 739.97, 'Completed'),
(2, 1344.96, 'Processing');

-- Order Items
INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
(1, 1, 1, 699.99), -- John's order: 1 smartphone
(1, 3, 2, 19.99),  -- John's order: 2 t-shirts
(2, 2, 1, 1299.99), -- Jane's order: 1 laptop
(2, 5, 3, 14.99);   -- Jane's order: 3 novels