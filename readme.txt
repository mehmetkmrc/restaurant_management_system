# Restaurant Management System (Golang)

This project is a **Restaurant Management System** backend built with Golang. It provides a set of APIs to handle various operations in a restaurant, such as managing menus, orders, invoices, tables, and users. This system is designed with scalability in mind and integrates with MongoDB for database operations.

## Features

- Manage restaurant menus, including creating, updating, and deleting menu items.
- Handle customer orders and track their details.
- Manage invoices, including creating and updating payment statuses.
- User authentication and role-based access control for secure operations.
- Table management to organize seating and table information.
- Error handling and validation for robust API performance.

## Tech Stack

- **Backend**: Golang (Gin Framework)
- **Database**: MongoDB
- **Authentication**: JWT
- **API Testing**: Postman (Collection included)

## Getting Started

To run the project locally:

1. Clone the repository:
   ```bash
   git clone https://github.com/mehmetkmrc/restaurant_management_system.git
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Set up MongoDB connection in your environment variables.

4. Run the server:
   ```bash
   go run main.go
   ```

## Postman Collection

A Postman collection is included in the `docs/postman/restaurant_api_collection.json` file to help you easily test the API endpoints.

To test the project, you can use the Postman collection by following these steps:

1. Download the `postman_collection/restaurant_management_system.postman_collection.json` file.

2. Open the Postman application and click the "Import" button.

3. Select the collection and run API requests.


## Postman Koleksiyonu
Projeyi test etmek için, aşağıdaki adımları izleyerek Postman koleksiyonunu kullanabilirsiniz:

1. `postman_collection/restaurant_management_system.postman_collection.json` dosyasını indirin.
2. Postman uygulamasını açın ve "Import" tuşuna tıklayın.
3. Koleksiyonu seçin ve API isteklerini çalıştırın.