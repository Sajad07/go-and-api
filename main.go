package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

//User data type
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
}

//UserStore for store a data
type UserStore struct {
	sync.Mutex
	store map[string]User
}

//methode for management POST and GET request for create a user and read all user
func (u *UserStore) postOrGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		u.getAllUsers(w, r)
		return
	case "POST":
		u.createOneUser(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

//this func for read all users
func (u *UserStore) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]User, len(u.store))
	u.Lock()
	i := 0
	for _, user := range u.store {
		users[i] = user
		i++
	}
	u.Unlock()
	jsonByte, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

//this func for create a user
func (u *UserStore) createOneUser(w http.ResponseWriter, r *http.Request) {
	jsonByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json',but got %v", ct)))
		return
	}
	var user User
	err = json.Unmarshal(jsonByte, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if user.ID == "" {
		user.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	u.Lock()
	u.store[user.ID] = user
	defer u.Unlock()
	fmt.Fprintf(w, "Successfully Create a User !")

}

//this func for selecte random user form users
func (u *UserStore) getRandomUser(w http.ResponseWriter, r *http.Request) {
	ids := make([]string, len(u.store))
	u.Lock()
	i := 0
	for id := range u.store {
		ids[i] = id
		i++
	}
	defer u.Unlock()
	var target string
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if len(ids) == 1 {
		target = ids[0]
	} else {
		rand.Seed(time.Now().UnixNano())
		target = ids[rand.Intn(len(ids))]
	}
	w.Header().Add("location", fmt.Sprintf("/users/%s", target))
	w.WriteHeader(http.StatusFound)
}

//this func for read a user
func (u *UserStore) getOneUser(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if parts[2] == "random" {
		u.getRandomUser(w, r)
		return
	}
	u.Lock()
	user, ok := u.store[parts[2]]
	u.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	jsonByte, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

//this func for creat a empty data store
func newUser() *UserStore {
	return &UserStore{
		store: map[string]User{},
	}
}

//Admin a struct for admin
type Admin struct {
	password string
}

//this func for create a admin
func newAdmin() *Admin {
	os.Setenv("ADMIN_PASSWORD", "abc123")
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("Required env var ADMIN_PASSWORD not set")
	}
	return &Admin{password: password}
}

//this func for handle admin request
func (a *Admin) adminHandle(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 Unauthorized"))
		return
	}
	w.Write([]byte("<html><h1>Super Secret Admin Portal</h1></html>"))
}

//main func for handler all function
func main() {
	admin := newAdmin()
	user := newUser()
	http.HandleFunc("/users", user.getAllUsers)
	http.HandleFunc("/users/", user.postOrGet)
	http.HandleFunc("/admin", admin.adminHandle)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
