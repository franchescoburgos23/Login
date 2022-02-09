package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/", Login)
	http.HandleFunc("/Sign-up", Signup)

	http.ListenAndServe(":8080", nil)

}

type User struct {
	Name string
}

var t = template.Must(template.ParseGlob("static/*"))

var conexionEstablish = ConnexionBD()

func ConnexionBD() (conexion *sql.DB) {

	Driver := "mysql"
	User := "root"
	Password := ""
	BDName := "Login"

	conexion, err := sql.Open(Driver, User+":"+Password+"@tcp(127.0.0.1)/"+BDName)

	if err != nil {
		panic(err.Error())
	}

	return conexion

}

func Login(w http.ResponseWriter, r *http.Request) {

	name := User{Name: "Franchesco"}

	t.ExecuteTemplate(w, "login", name)

}

func Signup(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "Sign-up", nil)

	insertConexion, err := conexionEstablish.Prepare("INSERT INTO Users(email,password) VALUES('correo@gmail.com','holamundo')")

	if err != nil {
		panic(err)
	}

	insertConexion.Exec()

}
