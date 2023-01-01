package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Need a registrar struct that contains this as an object.
var (
	stu []Student
	db  = connectDB()
)

type Student struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Id       string `json:"id"`
	Date     string `json:"date"`
}

// INSERT || UPDATE DATA
func insertIntoDB(w http.ResponseWriter, r *http.Request) {
	log.Println("Inserted into database")
	params := mux.Vars(r)
	stmt, err := db.Prepare("INSERT INTO Registers(Fullname, Email, Phone,UUID,RegisterDate) VALUES(?,?,?,?,?)")
	checkErr(err)
	stmt.Exec(params["fName"], params["email"], params["pNum"], uuid.New(), time.Now())
}
func updateName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Exec("UPDATE Registers SET Fullname = ? WHERE UUID = ?", params["fName"], params["id"])
}
func updateNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Exec("UPDATE Registers SET Phone = ? WHERE UUID = ?", params["pNum"], params["id"])
}
func updateEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Exec("UPDATE Registers SET Email = ? WHERE UUID = ?", params["email"], params["id"])
}

// DELETE DATA
func deleteWithId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.Exec("DELETE FROM Registers WHERE UUID = ?", params["id"])
}
func deleteAllFromDB(w http.ResponseWriter, r *http.Request) {
	db.Exec("DELETE FROM Registers")
}

// GET DATA
func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("SELECT * FROM Registers")
	checkErr(err)
	st := Student{}
	for rows.Next() {
		err := rows.Scan(&st.Fullname, &st.Email, &st.Phone, &st.Id, &st.Date)
		checkErr(err)
		stu = append(stu, st)
	}
	json.NewEncoder(w).Encode(stu)
}

// Reloading duplicates the same entry for both READ requests. Probably something to do in the For loop. It is unknown how to fix this.
/*Potential solution to this issue is to ensure the For loop is outside of the function. So when the request associated
Function is called, the loop isn't repeated again and again.
*/
func getWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	rows, err := db.Query("SELECT Fullname, Email, Phone, UUID,RegisterDate FROM Registers WHERE UUID = ?", params["id"])
	checkErr(err)
	st := Student{}
	for rows.Next() {
		err := rows.Scan(&st.Fullname, &st.Email, &st.Phone, &st.Id, &st.Date)
		checkErr(err)
		stu = append(stu, st)
	}
	json.NewEncoder(w).Encode(stu)
}

// ETC
func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./Database/MockDB.db")
	checkErr(err)
	return db
}
