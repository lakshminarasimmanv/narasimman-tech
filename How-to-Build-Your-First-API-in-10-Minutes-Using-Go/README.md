# How to Build Your First API in 10 Minutes Using Go

## A step-by-step guide for the above code.

### usersHandler

This is the handler for the `/users` endpoint. It returns a JSON-encoded list of users.

```go
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```

### usersCreateHandler

This is the handler for the `/users/create` endpoint. It creates a new user from the request body, and appends it to the `users` slice.

```go
func usersCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}
```

## A Quick Step-by-Step Guide:

* Clone this repo.

* run ```go run main.go``` This will start a sever at localhost:8080

* Using the terminal, you can test the API by running, ```curl http://localhost:8080/users```

* You can send a POST request by, ```curl -i -X POST -d '{"name":"Mia","age":25}' http://localhost:8080/users/create```

Congratulations you have successfully built you first API and tested using Go!
