package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// User struct represents a user with name, age, and friends.
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

var (
	db *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://username:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Post("/create", createUserHandler)
	r.Post("/make_friends", makeFriendsHandler)
	r.Delete("/user/{id}", deleteUserHandler)
	r.Get("/friends/{id}", getFriendsHandler)
	r.Put("/user/{id}", updateAgeHandler)

	http.ListenAndServe(":8080", r)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	row := db.QueryRow(statement, newUser.Name, newUser.Age)
	err = row.Scan(&newUser.ID)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}{
		ID:     newUser.ID,
		Status: "created",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func makeFriendsHandler(w http.ResponseWriter, r *http.Request) {
	var friendRequest struct {
		SourceID int `json:"source_id"`
		TargetID int `json:"target_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&friendRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sourceUser, err := getUserByID(friendRequest.SourceID)
	if err != nil {
		http.Error(w, "Source user not found", http.StatusNotFound)
		return
	}

	targetUser, err := getUserByID(friendRequest.TargetID)
	if err != nil {
		http.Error(w, "Target user not found", http.StatusNotFound)
		return
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	err = updateUser(sourceUser)
	if err != nil {
		http.Error(w, "Failed to update source user", http.StatusInternalServerError)
		return
	}

	err = updateUser(targetUser)
	if err != nil {
		http.Error(w, "Failed to update target user", http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
	}{
		Status: fmt.Sprintf("%s and %s are now friends", sourceUser.Name, targetUser.Name),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = deleteUser(user)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	response := struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}{
		Name:   user.Name,
		Status: "deleted",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getFriendsHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	friends, err := getUsersByIDs(user.Friends)
	if err != nil {
		http.Error(w, "Failed to retrieve friends", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friends)
}

func updateAgeHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updateAge struct {
		NewAge int `json:"new_age"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateAge)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Age = updateAge.NewAge

	err = updateUser(user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
	}{
		Status: "User age successfully updated",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getUserByID(userID int) (User, error) {
	statement := `SELECT id, name, age, friends FROM users WHERE id = $1`
	row := db.QueryRow(statement, userID)
	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.Age, pq.Array(&user.Friends))
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func getUsersByIDs(userIDs []int) ([]User, error) {
	users := make([]User, 0, len(userIDs))
	for _, userID := range userIDs {
		user, err := getUserByID(userID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func updateUser(user User) error {
	statement := `UPDATE users SET name = $1, age = $2, friends = $3 WHERE id = $4`
	_, err := db.Exec(statement, user.Name, user.Age, pq.Array(user.Friends), user.ID)
	return err
}

func deleteUser(user User) error {
	statement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(statement, user.ID)
	return err
}
