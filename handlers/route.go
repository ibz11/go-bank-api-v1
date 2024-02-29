package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"log"
	"github.com/gorilla/mux"
	"github.com/ibz11/go-bank-api-v1.git/models"
	"gorm.io/gorm"
)

func CreateAccount(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Error decoding JSON request", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	result := db.Create(&account)
	if result.Error != nil {

		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
	fmt.Println("CreateAccount")
}

func GetAccount(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var account []models.Account

	result := db.First(&account, params["id"])
	if result.Error != nil {
		// Check if the error is due to a record not found
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(account)
	fmt.Println("Get Account")
}

func MakeTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Make Transaction")
}

func CheckBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Check Balance on Account")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete  Account")
}
