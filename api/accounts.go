package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"swaggo-example/models"
)

var accounts = map[string]models.Account{}

// GetAccount retrieves an account by ID
// GetAccount godoc
// @Summary Retrieve an account
// @Description Get account details by ID
// @Tags Accounts
// @Param id path int true "Account ID"
// @Success 200 {object} models.Response[models.Account]
// @Failure 400 {object} models.Response[any]
// @Failure 404 {object} models.Response[any]
// @Router /account/{id} [get]
func GetAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		writeJSON(w, http.StatusBadRequest, models.Response[any]{
			StatusCode: 400,
			Message:    "Invalid URL path",
		})
		return
	}

	id := pathParts[4]

	account, exists := accounts[id]
	if !exists {
		writeJSON(w, http.StatusNotFound, models.Response[any]{
			StatusCode: 404,
			Message:    "Account not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, models.Response[models.Account]{
		StatusCode: 200,
		Message:    "Success",
		Data:       account,
	})
}

// CreateAccount creates a new account
// CreateAccount godoc
// @Summary Create a new account
// @Description Add a new account. Returns an error if the ID already exists.
// @Tags Accounts
// @Param account body models.Account true "Account object"
// @Success 201 {object} models.Response[any]
// @Failure 400 {object} models.Response[any]
// @Failure 409 {object} models.Response[any]
// @Router /account [put]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newAccount models.Account
	if err := json.NewDecoder(r.Body).Decode(&newAccount); err != nil {
		writeJSON(w, http.StatusBadRequest, models.Response[any]{
			StatusCode: 400,
			Message:    "Invalid request payload",
		})
		return
	}

	if _, exists := accounts[newAccount.ID]; exists {
		writeJSON(w, http.StatusConflict, models.Response[any]{
			StatusCode: 409,
			Message:    "Account ID already exists",
		})
		return
	}

	accounts[newAccount.ID] = newAccount
	writeJSON(w, http.StatusCreated, models.Response[any]{
		StatusCode: 201,
		Message:    "Account created successfully",
	})
}

// UpdateAccount updates an existing account
// UpdateAccount godoc
// @Summary Update an account
// @Description Update account fields except for ID
// @Tags Accounts
// @Param account body models.Account true "Account object"
// @Success 200 {object} models.Response[any]
// @Failure 400 {object} models.Response[any]
// @Failure 500 {object} models.Response[any]
// @Router /update-account [post]
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedAccount models.Account
	if err := json.NewDecoder(r.Body).Decode(&updatedAccount); err != nil {
		writeJSON(w, http.StatusBadRequest, models.Response[any]{
			StatusCode: 400,
			Message:    "Invalid request payload",
		})
		return
	}

	account, exists := accounts[updatedAccount.ID]
	if !exists {
		writeJSON(w, http.StatusInternalServerError, models.Response[any]{
			StatusCode: 500,
			Message:    "Account not found",
		})
		return
	}

	account.FirstName = updatedAccount.FirstName
	account.LastName = updatedAccount.LastName
	account.Status = updatedAccount.Status
	account.Balance = updatedAccount.Balance

	accounts[updatedAccount.ID] = account
	writeJSON(w, http.StatusOK, models.Response[any]{
		StatusCode: 200,
		Message:    "Account updated successfully",
	})
}

// Utility function to write JSON responses
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
