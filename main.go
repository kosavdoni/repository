package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

    "./main/models"
    "./main/repository"
	_ "github.com/lib/pq" // Драйвер для PostgreSQL
	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type OrderProduct struct {
	OrderID   int `json:"orderId"`
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
	Product   Product
}

type Order struct {
	ID         int            `json:"id"`
	UserID     int            `json:"userId"`
	TotalPrice float64        `json:"totalPrice"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Products   []OrderProduct `json:"products"`
}

type OrderStatus int

const (
	Pending OrderStatus = iota
	Processed
	Delivered
)

var users []User
var products []Product
var orders []Order

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// Начало чего-то страшного
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	// Выполнение запроса на удаление пользователя из базы данных
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		log.Fatal(err)
	}

	// Возврат успешного статуса в ответе
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	// Выполнение запроса на удаление продукта из базы данных
	_, err := db.Exec("DELETE FROM products WHERE id = $1", productID)
	if err != nil {
		log.Fatal(err)
	}

	// Возврат успешного статуса в ответе
	w.WriteHeader(http.StatusOK)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderID := params["id"]

	// Выполнение запроса на удаление заказа из базы данных
	_, err := db.Exec("DELETE FROM orders WHERE id = $1", orderID)
	if err != nil {
		log.Fatal(err)
	}

	// Возврат успешного статуса в ответе
	w.WriteHeader(http.StatusOK)
}

var db *sql.DB

//и конец

func main() {
	// Подключение к базе данных
	db, err := sql.Open("postgres", "user=myuser password= dbname=mydatabase sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Инициализация репозиториев
	// Инициализация репозиториев
	user_repository := &repository.UserRepository{DB: db}

	router := mux.NewRouter()

	// Установка обработчиков запросов
	router.HandleFunc("/users", GetUserHandler(userRepo)).Methods("GET")
	router.HandleFunc("/users", CreateUserHandler(userRepo)).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
