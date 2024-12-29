# 🚀 Rapid Key-Value Database

A lightweight, persistent, and fast **Key-Value Database** inspired by the **Bitcask storage model** written in Golang. It is optimized for high-speed reads and writes, using in-memory indexing and an append-only log structure.

# 📦Features

-   **Append-Only Write Log:** Ensures write speed and durability.
-   **In-Memory Index:** Quick access to data using in-memory maps.
-   **Crash Recovery:** Log files ensure data durability.
-   **HTTP API:** Simple APIs for `PUT`, `GET`, and future `MERGE` operations.

## 🔧 **Installation**
1. Clone the repository:
> git clone https://github.com/yourusername/rapidkv-db.git
> cd rapidkv-db
2. Build and run:
> go build -o rapidkv-db ./cmd/main.go 
> go run .\cmd\main.go

## Test APIs using **curl** or tools like **Postman**:
1.  Add key-value pair.
	> curl --location 'http://localhost:9090/put' \
--header 'Content-Type: application/json' \
--data '{
    "key": "hans",
    "value": "ravi"
}'

2. Retrieve a value by key
    > curl --location 'http://localhost:9090/get?key=hans'

## 🛠️ **Configuration**
-   Default port: `9090`
-   Default log directory: `./data`

## 📚 **API Endpoints**

| Method | Endpoint | Description |
| -------| -----------| -------|
| POST   | `/put`|Add a key-value pair
| GET| `/get`| Retrieve a value
| POST| `/merge`| Manually Trigger data merge
| GET| `/health`| Check server health

## 👥 **Contributing**

-   Fork the repository
-   Create a new branch for your feature/bugfix
-   Submit a pull request

## 📝 **License**

This project is licensed under the **MIT License**.


## 🚀 **TODO List**

### 🔨 **Core Features**

- [x] Implement `PUT` handler for writing key-value pairs.
- [x]  Implement `GET` handler for retrieving key-value pairs.
- [ ]  Implement `MERGE` operation to compact and optimize log files.
- [ ]  Add TTL (Time-To-Live) support for keys.
- [ ]  Add data persistence validation on startup.
- [ ]  Thread Safety.