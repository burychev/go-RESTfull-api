# go-RESTfull-api
RESTfull API Products and Categories

```shell
# download the starter kit
git clone https://github.com/burychev/go-RESTfull-api.git

Create DataBase in PostgreSql
(run the script below)
Connect to db
```
```sql
CREATE TABLE ProductCategory (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255) NOT NULL
);

CREATE TABLE Product (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255) NOT NULL,
    Price DECIMAL(10, 2) NOT NULL,
    CategoryId INT,
    FOREIGN KEY (CategoryId) REFERENCES ProductCategory(Id)
);
```
```shell
#run the script in terminal:
 go run ./cmd/app/main.go
```

At this time, you have a RESTful API server running at `http://localhost:8080`. It provides the following endpoints:

Product Categories:
* `GET /categories`: get all categories
* `POST /categories`: сreate a category
 ```JSON
{
 "name":        "Vegetables",
 "description": "Cucumbers and tomatoes"
}
```
* `GET /categories/:id`: get category by id
* `PUT /categories/:id`: update category by id
 ```JSON
{
 "name":        "New Vegetables",
 "description": "Cucumbers, tomatoes and carrots"
}
```
* `DELETE /categories/:id`: delete category by id



Products:
* `GET /products`: get all products
* `POST /products`: сreate a product
 ```JSON
{
 "name":        "Cucumber",
 "description": "It is Cucumber",
  "price": 12.34,
  "category_id": 1
}
```
* `GET /products/:id`: get product by id
* `PUT /products/:id`: update product by id
 ```JSON
{
 "name":        "New cucumber",
 "description": "It is new cucumber",
  "price": 12.35,
  "category_id": 1
}
```
* `DELETE /products/:id`: delete product by id
