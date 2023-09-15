# HumanResourceAPI Documentation

## Introduction

The **HumanResourceAPI** is a simple REST API designed to perform CRUD (Create, Read, Update, Delete) operations on a "human" resource. It interfaces with MongoDB as the database of choice and is built using the Go programming language. This API is capable of dynamically handling parameters, allowing you to add or retrieve a person by name and perform various other operations.

## Setup and Deployment

### Local Development

1. **Install Go:** Make sure you have Go installed on your local machine.

2. **Clone the Repository:** Clone the API repository from GitHub
  
3. **Set Database URL:**

Export or set the environment variable `DatabaseUrl` to specify your MongoDB database URL. This is essential for connecting the API to the database.

Example:

```shell
export DatabaseUrl="mongodb://localhost:27017/your_database"
```
4. Run the API: Start the API by running the main Go file:
```shell
go run main.go
```
## Endpoints

### Get All (GET /api)

- **HTTP Method:** GET
- **Description:** Retrieve a list of all human resources.
- **Request Parameters:** None
- **Response Format:** JSON

#### Example Output:
```json
HTTP 200 OK
{
    "data": [
        {
            "user_id": "65045f1342ac6c035118d7d9",
            "user_name": "Mustapha",
            "created_at": "2023-09-15T13:41:39.369Z",
            "updated_at": "2023-09-15T13:41:39.369Z"
        },
        {
            "user_id": "6504600f42ac6c035118d7da",
            "user_name": "Awwal",
            "created_at": "2023-09-15T13:45:51.728Z",
            "updated_at": "2023-09-15T13:45:51.728Z"
        }
    ],
    "status": "success"
}
```

### Get by ID (GET /api/:user_id)

- **HTTP Method**: GET
- **Description**: Retrieve a specific human resource by their ID.
- **Request Parameters**: None
- **user_id (Path parameter)**: The ID of the human resource to retrieve.
- **Response Format:** JSON

#### Example Output:
```json
HTTP 200 OK
{
    "data": {
        "user_id": "650469e6dbfe09bd1a62ab98",
        "user_name": "third",
        "created_at": "2023-09-15T14:27:50.071Z",
        "updated_at": "2023-09-15T14:27:50.071Z"
    },
    "status": "success"
}
```

### Update (PATCH /api/:user_id)
- **HTTP Method**: PATCH
- **Description*: Update a specific human resource by their ID.
- **Request Parameters**:
- **user_id (Path parameter)**: The ID of the human resource to update.
- **Request Body**: JSON with updated name.
```
#### Example Request body with a new name :
```json
{
    "user_name": "moon"
}
```

#### Example Output:
```json
HTTP 200 OK
{
    "data": {
        "user_id": "650469e6dbfe09bd1a62ab98",
        "user_name": "moon",
        "created_at": "2023-09-15T14:27:50.071Z",
        "updated_at": "2023-09-15T14:27:50.071Z"
    },
    "status": "success"
}
```

### Create New (POST /api)
- **HTTP Method**: POST
- **Description**: Create a new human resource.
- **Request Body**: JSON with new resource data.
#### Example Request body with a name :
```json
{
    "user_name": "NewName"
}
```

#### Example Output:
```json
HTTP 200 OK
{
    "data": {
        "user_id": "650469e6dbfe09bd1a62ab98",
        "user_name": "NewName",
        "created_at": "2023-09-15T14:27:60.071Z",
        "updated_at": "2023-09-15T14:27:60.071Z"
    },
    "status": "success"
}
```

### Delete (DELETE /api/:user_id)
- **HTTP Method**: DELETE
- **Description**: Delete a specific human resource by their ID.
- **user_id (Path parameter)**: The ID of the human resource to delete.
#### Example Output:
```json
HTTP 204 No content
```

## Conclusion

Congratulations! You have successfully explored the **HumanResourceAPI** documentation. This API empowers you to perform CRUD operations on human resources, dynamically handle parameters, and interface with a MongoDB database.

- Use the provided endpoints to retrieve, update, create, or delete human resources.
- Set up the API locally for development using the provided instructions.
- Refer to the API's GitHub repository for advanced deployment and configuration options.

If you have any questions, encounter issues, or wish to contribute to the development of this API, please feel free to [open an issue](https://github.com/your-api-repository/issues) on the GitHub repository. Your feedback and contributions are highly appreciated.

Thank you for using the **HumanResourceAPI**. We hope it serves as a valuable resource for your projects and applications.


 
