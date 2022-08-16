# Creating a REST API in Go with JSON Web Tokens

To view the tutorial article, click the [link](https://bit.ly/3zPOtKu) below.

**Article:** https://bit.ly/3zPOtKu

## How to run the program?

First, open the terminal and start the application.

```
go run main.go
```

**Output:**

``Starting the application...``

Now, open the second terminal and send a POST request to the ``/authenticate`` endpoint.

```
curl -X POST http://localhost:12345/authenticate -H 'Content-Type: application/json' -d '{"username": "user1", "password": "password1"}'
```

**Output:**

``{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwicGFzc3dvcmQiOiJwYXNzd29yZDEiLCJleHAiOjE1NDU1NjU4NzB9.4cVIXoCFHbE7VYbS9XIxCv4mF2uV7O4Zm4uLpjK05zE"}``

In the response, you will see a JWT token. Now, open the third terminal and send a GET request to the ``/protected`` endpoint along with the token as a query parameter.

```
curl -X GET http://localhost:12345/protected?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwicGFzc3dvcmQiOiJwYXNzd29yZDEiLCJleHAiOjE1NDU1NjU4NzB9.4cVIXoCFHbE7VYbS9XIxCv4mF2uV7O4Zm4uLpjK05zE
```

**Output:**


`{"username":"user1","password":"password1"}`

If you don't pass the token as a query parameter, you will get an error.

```
curl -X GET http://localhost:12345/protected
```

**Output:**

`{"message":"An authorization header is required"}`

If you pass the wrong token, you will get an error.

```
curl -X GET http://localhost:12345/protected?token=abcdef
```

**Output:**

`{"message":"signature is invalid"}`

If you make any other request, you will get an error.

```
curl -X POST http://localhost:12345/protected`
```

**Output:**

`{"message":"method not allowed"}`

If you make any other request, you will get an error.

```
curl -X GET http://localhost:12345/authenticate`
```

**Output:**

`{"message":"method not allowed"}`

## Conclusion

JWT is a very popular way to authenticate users and protect APIs. Go language supports JWT and it's quite easy to use JWT with Go.
