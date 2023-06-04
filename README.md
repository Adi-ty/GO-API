# GO-API

This repository provides a template project for building a Go API. The API is built with Gorilla Mux, Postgres, Docker and Docker Compose. The API includes basic CRUD (Create, Read, Update, Delete) functionality, allowing you to quickly start developing your own API.

## Installation 

- Clone the repository: 
  ```
  git clone https://github.com/{YOUR_USERNAME}/GO-API.git
  ```
  
- Change to the project directory:
  ```
  cd GO-API
  ```

- Start the server
  ```
  task run 
  ```

## End Points

- Post Comment
  - URL: `/api/v1/comment`
  - Method: POST
  - Authorization: bearer token(with signature "Aditya")
  
    **Request Body**
    ```
    {
    "slug": "hello",
    "body": "new body",
    "author": "Adi"
    }
    ```

    **Response**
    ```
    {
    "ID": "80518c3b-8bf9-434f-856c-829e525b6d13",
    "Slug": "hello",
    "Body": "new body",
    "Author": "Adi"
    }
    ```

- Get Comment by id
  - URL: `/api/v1/comment/{id}`
  - Method: GET

    **Response**
    ```
    {
    "ID": "80518c3b-8bf9-434f-856c-829e525b6d13",
    "Slug": "hello",
    "Body": "new body",
    "Author": "Adi"
    }
    ```

- Update Comment
  - URL: `/api/v1/comment/{id}`
  - Method: PUT
  - Authorization: bearer token(with signature "Aditya")
  
    **Request Body**
    ```
    {
    "slug": "Testing put",
    "author": "Aditya Singh",
    "body": "body vo to bnva do"
    }
    ```

    **Response**
    ```
    {
    "ID": "b9408103-d836-4fc4-a47e-0c82c49cef09",
    "Slug": "Testing put",
    "Body": "body",
    "Author": "Aditya Singh"
    }
    ```

- Delete Comment
  - URL: `/api/v1/comment/{id}`
  - Method: DELETE
  - Authorization: bearer token(with signature "Aditya")
  
    **Response**
    ```
    {
    "Message": "Successfully Deleted"
    }
    ```

## JWT Authentication 

To authorize requests for creating, updating, or deleting comments, a bearer token with the signature "Aditya" needs to be included in the Authorization header. You can generate the JWT token using a tool like [jwt.io](https://jwt.io/). Please make sure to keep the token secure and only share it with trusted individuals.

## Contributing 

If you would like to contribute to this project, you are welcome to do so! You can contribute by:

- Opening an issue to report a bug or suggest a new feature.
- Opening a pull request to propose changes or fixes.

Please provide a clear description of the issue or your proposed changes when opening an issue or pull request. Your contributions are greatly appreciated!

