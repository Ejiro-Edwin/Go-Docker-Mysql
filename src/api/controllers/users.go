package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("List Users"))
}
func CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Create User"))
}
func GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Get User"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Update Users"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Delete User"))
}