# HumanResourceAPI

## Description

The **HumanResourceAPI** is a simple REST API built with Go and MongoDB. It allows you to perform CRUD operations on a "human" resource, including creating, retrieving, updating, and deleting individuals. You can also search for people by name. This API is suitable for various applications that require managing human resources.

## Getting Started

To set up and run the HumanResourceAPI locally, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/human-resource-api.git
   ```
2. Set the DatabaseUrl environment variable to your MongoDB database URL:

   ```shell
     Copy code
     export DatabaseUrl=mongodb://localhost:27017/your_database
   ```
3. Run the API:
   ```shell
     Copy code
     go run main.go
   ```
The API will be accessible locally at the specified endpoints on port 80.

## Usage

Use the API's endpoints to interact with human resources data. You can perform actions such as retrieving all records, getting a person by ID, updating a person's information, deleting a person, and creating a new person.

For more detailed information and examples of API usage, please refer to the [documentation.md](./DOCUMENTATION.md) file.

