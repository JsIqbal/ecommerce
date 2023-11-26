package db

var DbSchema = `
	DROP TABLE IF EXISTS product_stocks;
	DROP TABLE IF EXISTS products;
	DROP TABLE IF EXISTS brands;
	DROP TABLE IF EXISTS categories;
	DROP TABLE IF EXISTS suppliers;

	CREATE TABLE IF NOT EXISTS brands (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		status_id INTEGER NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS categories (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		parent_id UUID,
		sequence INTEGER,
		status_id INTEGER NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS suppliers (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		phone VARCHAR(20),
		status_id INTEGER NOT NULL,
		is_verified_supplier BOOLEAN NOT NULL,
		created_at BIGINT NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS products (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name VARCHAR(255) NOT NULL,
		description TEXT,
		specifications TEXT,
		brand_id UUID REFERENCES brands(id) NOT NULL,
		category_id UUID REFERENCES categories(id) NOT NULL,
		supplier_id UUID REFERENCES suppliers(id) NOT NULL,
		unit_price NUMERIC NOT NULL,
		discount_price NUMERIC,
		tags VARCHAR(255)[],
		status_id INTEGER NOT NULL,
		created_at BIGINT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS product_stocks (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		product_id UUID REFERENCES products(id) NOT NULL,
		stock_quantity INTEGER NOT NULL,
		updated_at BIGINT NOT NULL
	);
`
