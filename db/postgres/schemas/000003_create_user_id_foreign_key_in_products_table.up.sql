ALTER TABLE IF EXISTS products
ADD CONSTRAINT products_user_id_fkey
FOREIGN KEY (user_id)
REFERENCES users (id);