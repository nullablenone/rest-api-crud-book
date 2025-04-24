# 📚 REST API CRUD Book

A simple RESTful API for managing book data using Golang and GORM. This project is designed to facilitate CRUD operations (Create, Read, Update, Delete) on book data, with a PostgreSQL database connection.

---

## 🛠️ Technologies Used

- Golang  
- GORM  
- PostgreSQL  
- Godotenv (for environment variable management)


---

## ⚙️ Setup & Installation

### 1. Clone the Repository

```bash
git clone https://github.com/nullablenone/rest-api-crud-book.git
cd rest-api-crud-book
```

### 2. Create `.env` File

Copy `.env.example` to `.env` and edit the values based on your local database configuration:

```bash
cp .env.example .env
```

Example `.env` content:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=yourdatabase
DB_SSLMODE=disable
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run main.go
```

The API will be running at `http://localhost:8080`.

---

## 📚 API Endpoints

| Method | Endpoint       | Description               |
|--------|----------------|---------------------------|
| GET    | /books         | Get all books             |
| GET    | /books/{id}    | Get book by ID            |
| POST   | /books         | Add a new book            |
| PUT    | /books/{id}    | Update existing book data |
| DELETE | /books/{id}    | Delete a book             |

---


## 🤝 Contributions

Contributions are very welcome! Feel free to fork this repository and submit a pull request for improvements or new features.

---


If you need further assistance or would like to add a specific feature, please feel free to contact me!

