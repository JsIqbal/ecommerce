# Ecommerce

Your Comprehensive Guide for the Project

## This project is deployed in ROBI Cloud:

-   Endpoint:

## Table of Contents

1. [Seed Database](#seed-database)
2. [Start Application Locally](#start-application-locally)
3. [Start Application with Docker](#start-application-with-docker)

<a name="seed-database"></a>

## 1. Seed Database(locally) - you need to run postgresql locally

To seed the database:

-   Configure your `.env` file.
-   Navigate to the root of the application.
-   Run the following command in the terminal:

```bash
make seed

```

## 1.1. Seed Database(Docker)

To seed the database:

-   Configure your `.env` file.
-   Navigate to the root of the application.
-   Run the following command in the terminal:

```bash
make serve
```

-   After executing the command you will be prompt with 2 options:

    -   option 1: Docker
    -   option 2: local

-   Choose the Docker
-   You will see an error saying that connection failed but a db container will be created
-   Run the following command in another termianl in the same directory:

```bash
make seed
```

<!-- Start Application -->

# Start Application

To start the application locally

-   configure the .env file
-   go to root of the application
-   run the following command in the terminal

```bash
make serve
```

-   After executing the command you will be prompt with 2 options:

    -   option 1: Docker
    -   option 2: local

-   Choose the Docker
-   You will see an error saying that connection failed but a db container will be created
-   Run the following command in another termianl in the same directory:

```bash
make seed
```

# API Docs:

# Brand APIs:

## End-point: Create brand

### Method: POST

> ```
> http://localhost:5000/api/brands
> ```

### Body (**raw**)

```json
{
    "name": "ASUS",
    "status_id": 1
}
```

## End-point: Get brand

### Method: GET

> ```
> http://localhost:5000/api/brands/:id
> ```

## End-point: Update Brand

### Method: PUT

> ```
> http://localhost:5000/api/brands/:id
> ```

### Body (**raw**)

```json
{
    "name": "Lenevo",
    "status_id": 1
}
```

## End-point: Delete Brand

### Method: DELETE

> ```
> http://localhost:5000/api/brands/:id
> ```

## End-point: Get brands

### Method: GET

> ```
> http://localhost:5000/api/brands?page=1&limit=2
> ```

### Query Parameters:

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 2     |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

# Product

## End-point: Create product

### Method: POST

> ```
> http://localhost:5000/api/products
> ```

### Body (**raw**)

```json
{
    "name": "Lenovo Think V2",
    "description": "Powerful laptop for professional use.",
    "brand_id": "5f2dc58e-d3a8-4580-b4fb-0e72d93f0afe",
    "category_id": "8ace9e3f-3bca-4deb-8128-e0f67b0c0924",
    "supplier_id": "5d96a2df-370b-4afd-a7c4-cfcc1e7241d2",
    "unit_price": 50.05,
    "discount_price": 12.54,
    "tags": ["business", "professional"],
    "status_id": 1,
    "stock_quantity": 100
}
```

## End-point: Get product

### Method: GET

> ```
> http://localhost:5000/api/products/:id
> ```

## End-point: Update product

### Method: PUT

> ```
> http://localhost:5000/api/products/:id
> ```

### Body (**raw**)

```json
{
    "productName": "Lenovo Think V2",
    "productDescription": "A powerful laptop designed for professional use.",
    "brandId": "5f2dc58e-d3a8-4580-b4fb-0e72d93f0afe",
    "categoryId": "8ace9e3f-3bca-4deb-8128-e0f67b0c0924",
    "supplierId": "5d96a2df-370b-4afd-a7c4-cfcc1e7241d2",
    "unitPrice": 50.05,
    "discountPrice": 12.54,
    "tags": ["abc", "xyz"],
    "statusId": 1,
    "stockQuantity": 100
}
```

## End-point: Delete product

### Method: DELETE

> ```
> http://localhost:5000/api/products/:id
> ```

## End-point: Get products

### Method: GET

> ```
> http://localhost:5000/api/products?page=1&limit=20
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 20    |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

# Supplier APIs

## End-point: Create supplier

### Method: POST

> ```
> http://localhost:5000/api/suppliers
> ```

### Body (**raw**)

```json
{
    "name": "Iqbal Hossain",
    "email": "zafar.iq3089@gmail.com",
    "phone": "01403229479",
    "status_id": 1,
    "is_verified_supplier": true
}
```

## End-point: Get supplier

### Method: GET

> ```
> http://localhost:5000/api/suppliers/:id
> ```

## End-point: Update supplier

### Method: PUT

> ```
> http://localhost:5000/api/suppliers/:id
> ```

### Body (**raw**)

```json
{
    "name": "THE KRAKEN",
    "email": "kraken@gmail.com",
    "phone": "01403229479",
    "status_id": 1,
    "is_verified_supplier": true
}
```

## End-point: Delete supplier

### Method: DELETE

> ```
> http://localhost:5000/api/suppliers/:id
> ```

## End-point: Get suppliers

### Method: GET

> ```
> http://localhost:5000/api/suppliers?page=1&limit=5
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 5     |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

# Category APIs:

## End-point: Create category

### Method: POST

> ```
> http://localhost:5000/api/categories
> ```

### Body (**raw**)

```json
{
    "name": "Android",
    "parent_id": "pef438e9-2c04-4e12-961d-d35e2d75e5cf",
    "status_id": 1
}
```

## End-point: Get category

### Method: GET

> ```
> http://localhost:5000/api/categories/:id
> ```

## End-point: Update categories

### Method: PUT

> ```
> http://localhost:5000/api/categories/:id
> ```

### Body (**raw**)

```json
{
    "name": "iphone",
    "status_id": 1
}
```

## End-point: Delete Brand

### Method: DELETE

> ```
> http://localhost:5000/api/categories/:id
> ```

## End-point: Get categories

### Method: GET

> ```
> http://localhost:5000/api/categories?page=1&limit=5
> ```

### Query Params

| Param | value |
| ----- | ----- |
| page  | 1     |
| limit | 5     |

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get category tree

### Method: GET

> ```
> http://localhost:5000/api/categories/tree
> ```

## Thank You!

Thank you for your time and assistance! 🙌 If you have any more questions or need further help, feel free to [reach out](https://github.com/JsIqbal). Have a great day!
