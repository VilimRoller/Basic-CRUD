# Basic-CRUD
Basic-CRUD is a simple CRUD (Create, Read, Update, Delete) application built with Golang. It provides a REST API for managing expenses stored in a Redis database. The application uses Gorilla Mux for HTTP router and Redis for database.

## Features
- Add new expense
- Retrieve and existing expense
- Update an existing expense
- Delete an existing expense

## Prerequisites
- You need Docker installed on your machine

## Installation
1. Clone this repository
```
git clone https://github.com/VilimRoller/Basic-CRUD.git
```
3. Build Docker image
```
docker build -t basic-crud .
```
3. Run Docker containers
```
docker compose up -d
```

## Testing
To run tests, you need Golang installed on your machine.
Run test command from the root directory of the project
```
go test ./...
```

## API Endpoints
### Home
- URL: http://172.18.0.3:8080/Basic-Crud
- Method: Any method
- Description: Checking if API connection works
- Response:
```
Api is running!
```

### Add an expense
- URL: http://172.18.0.3:8080/Basic-Crud/expenses
- Method: POST
- Description: Adding expense to the database
- Body: raw
```
Name: Expense Name
Date: 25-06-2023
Type: Food
Amount: 20.5
Currency: EUR
```
- Response:
```
key = unique_key
```
unique_key is used to retrieve, update or delete the expense.

### Retrieve an expense
- URL: http://172.18.0.3:8080/Basic-Crud/expenses?key=unique_key
- Method: GET
- Description: Retrieving the expense corresponding to unique_key
- Response:
```
Name: Expense Name
Date: 25-06-2023
Type: Food
Amount: 20.5
Currency: EUR
```

Unique key "all" will retrieve all expenses from the database.

### Update an expense
- URL: http://172.18.0.3:8080/Basic-Crud/expenses?key=unique_key
- Method: PUT
- Description: Update the expense corresponding to unique_key
- Body: raw
```
Name: Updated Name
Date: 25-06-2023
Type: Entertainment
Amount: 50.5
Currency: USD
```
- Response:
```
Update successful!
```
### Delete an expense
- URL: http://172.18.0.3:8080/Basic-Crud/expenses?key=unique_key
- Method: DELETE
- Description: Deleting the expense corresponding to unique_key
- Response:
```
Delete successful!
```
Unique key "all" will delete all expenses from the database.
