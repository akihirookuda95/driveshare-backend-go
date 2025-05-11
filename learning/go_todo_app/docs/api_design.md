### POST /todos
- **Purpose**: Create a new todo
- **Request (JSON)**
```json
{
  "title": "Buy milk",
  "description": "2L low-fat",
  "is_done": false
}
```
- **Response (JSON)**
```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "2L low-fat",
  "is_done": false
}
```
- **Status Codes**:
  - `201 Created`: Todo created successfully
  - `400 Bad Request`: Invalid input data




### GET /todos
- **Purpose**: Get user's todos
- **Request (JSON)**
  - No requst body
- **Response (JSON)**
```json
[
  {
    "id": 1,
    "title": "Buy milk",
    "description": "2L low-fat",
    "is_done": false
  },
  {
    "id": 2,
    "title": "Read book",
    "description": "Go programming",
    "is_done": true
  }
]
```
- **Status Codes**:
  - `200 OK`: Todos retrieved successfully
  - `401 Unauthorized`: User not authenticated(token expired)



### PUT /todos/{id}
- **Purpose**: Update a todo
- **Request (JSON)**
```json
{
 "id": 1,
 "title": "Buy milk",
 "description": "2L low-fat",
 "is_done": false
}
```
- **Response (JSON)**
```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "2L low-fat",
  "is_done": true
}
```
- **Status Codes**:
 - `200 OK`: Todos retrieved successfully
 - `404 Not Found`: Todo not found



### DELETE /todos/{id}
- **Purpose**: Delete a todo
- **Request (JSON)**
```json
{
  "id": 1
}
```
- **Response (JSON)**
```json
{
  "message": "Todo deleted successfully"
}
```
- **Status Codes**:
  - `200 OK`: Todo deleted successfully
  - `404 Not Found`: Todo not found



### GET /todos/{id}
- **Purpose**: Get a specific todo
- **Request (JSON)**
```json
{
  "id": 1
}
```
- **Response (JSON)**
```json
{
  "id": 1,
  "title": "Buy milk",
  "description": "2L low-fat",
  "is_done": false
}
```
- **Status Codes**:
  - `200 OK`: Todo retrieved successfully
  - `404 Not Found`: Todo not found














