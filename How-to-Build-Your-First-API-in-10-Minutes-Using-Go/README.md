# A step-by-step guide for the above code using markdown.

## usersHandler

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

## usersCreateHandler

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

# A Quick Step-by-Step Guide:

* Clone this repo.

* run ```go run main.go```
