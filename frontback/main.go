package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=13111995 dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
	<body  bgcolor="Lightskyblue">  
		<form action="/create" method="POST">

			<label for="firstname">FirstName:</label>
			<input type="text" id="firstname" name="firstname"><br><br>

			<label for="middlename">MiddleName:</label>
			<input type="text" id="middlename" name="middlename"><br><br>
			
			<label for="lastname">LastName:</label>
			<input type="text" id="lastname" name="lastname"><br><br>

			<label for="gender">Gender:</label>
			<input type="text" id="gender" name="gender"><br><br>
			
			<label for="phone">Phone:</label>
			<input type="text" id="phone" name="phone"><br><br>

			<label for="address">Address:</label>
			<input type="text" id="address" name="address"><br><br>
			
			<label for="email">Email:</label>
			<input type="email" id="email" name="email"><br><br>

			<label for="password">Password:</label>
			<input type="Password" id="password" name="password"><br><br>
			
			<input type="submit" value="Submit">
		</form>
	</body>
</html>  `)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract form data
	firstname := r.FormValue("firstname")
	middlename := r.FormValue("middlename")
	lastname := r.FormValue("lastname")
	gender := r.FormValue("gender")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Insert data into the database
	db, err := CreateConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO students (firstname, middlename,lastname,gender,phone,address,email,password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", firstname, middlename, lastname, gender, phone, address, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to index page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/create", CreateHandler).Methods("POST")

	http.ListenAndServe(":8099", r)
}
