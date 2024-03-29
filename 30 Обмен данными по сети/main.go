/*
Цель практической работы

Научиться:

    работать с запросами POST, GET, PUT, DELETE;
    применять принципы написания обработчиков HTTP-запросов.


Что нужно сделать

Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:

    1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу. Пример запроса:

POST /create HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"name":"some name","age":"24","friends":[]}

Данный запрос должен возвращать ID пользователя и статус 201.


    2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья. Пример запроса:

    POST /make_friends HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"source_id":"1","target_id":"2"}

Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».


    3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей. Пример запроса:

    DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"target_id":"1"}

Данный запрос должен возвращать 200 и имя удалённого пользователя.


    4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:

    GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close

После /friends/ указывается id пользователя, друзей которого мы хотим увидеть.


    5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:

PUT /user_id HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"new age":"28"}

Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».


Советы и рекомендации

Воспользуйтесь библиотекой Chi.


Что оценивается

Дизайн API и работа в соответствии с функциональными требованиями.


Как отправить работу на проверку

Выполните задание в файле вашей среды разработки и пришлите ссылку на архив с вашим проектом через форму ниже. */

package main

//перед запусокм подключаем библиотеку go get -u github.com/go-chi/chi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// User struct represents a user with name, age, and friends.
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

var (
	users  map[int]User
	nextID int
)

func main() {
	users = make(map[int]User)
	nextID = 1

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

	newUser.ID = nextID
	users[nextID] = newUser
	nextID++

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

	sourceUser, sourceExists := users[friendRequest.SourceID]
	targetUser, targetExists := users[friendRequest.TargetID]

	if !sourceExists || !targetExists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	sourceUser.Friends = append(sourceUser.Friends, targetUser.ID)
	targetUser.Friends = append(targetUser.Friends, sourceUser.ID)

	users[sourceUser.ID] = sourceUser
	users[targetUser.ID] = targetUser

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

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Remove the user from friends' lists
	for _, friendID := range user.Friends {
		friend, friendExists := users[friendID]
		if friendExists {
			friendFriends := friend.Friends[:0]
			for _, friendFriendID := range friend.Friends {
				if friendFriendID != user.ID {
					friendFriends = append(friendFriends, friendFriendID)
				}
			}
			friend.Friends = friendFriends
			users[friend.ID] = friend
		}
	}

	delete(users, userID)

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

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	friends := make([]User, 0, len(user.Friends))
	for _, friendID := range user.Friends {
		friend, friendExists := users[friendID]
		if friendExists {
			friends = append(friends, friend)
		}
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

	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Age = updateAge.NewAge
	users[userID] = user

	response := struct {
		Status string `json:"status"`
	}{
		Status: "User age successfully updated",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
