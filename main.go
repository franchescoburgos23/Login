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
	http.HandleFunc("/registred", Registred)
	http.HandleFunc("/check", Check)
	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/errml", Errml)
	http.HandleFunc("/errps", Errps)

	http.ListenAndServe(":8080", nil)

}

type User struct {
	Id       string
	Email    string
	Password string
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

	t.ExecuteTemplate(w, "login", nil)

}

func Signup(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "Sign-up", nil)

}

func Registred(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		email := r.FormValue("email")
		password := r.FormValue("password")

		insertConexion, err := conexionEstablish.Prepare("INSERT INTO Users(email,password) VALUES(?,?)")

		if err != nil {
			panic(err)
		}

		insertConexion.Exec(email, password)

		http.Redirect(w, r, "/", 301)
	}
}

func Check(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		record, err := conexionEstablish.Query("SELECT *FROM Users")
		if err != nil {

			panic(err.Error())

		}

		user := User{}
		arrayuser := []User{}

		for record.Next() {
			var id int
			var email string
			var password string
			err = record.Scan(&id, &email, &password)
			if err != nil {
				panic(err.Error())
			}

			user.Email = email
			user.Password = password

			arrayuser = append(arrayuser, user)
		}

		VAemail := r.FormValue("mail")
		VApassword := r.FormValue("ps")

		if VAemail == user.Email {

			if VApassword == user.Password {

				http.Redirect(w, r, "/dashboard", 301)
			} else {

				http.Redirect(w, r, "/errps", 301)
			}

		} else {
			http.Redirect(w, r, "/errml", 301)
		}

	}

}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "dashboard", nil)

}

func Errml(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "Loginerrml", nil)

}

func Errps(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "Loginerrps", nil)
}
