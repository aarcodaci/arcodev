package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aarcodaci/arcodev/db"

	"gopkg.in/gomail.v2"
)

type consulta struct {
	Apellido    string `json:"Apellido"`
	Nombre      string `json:"Nombre"`
	Email       string `json:"email"`
	Email2      string `json:"email2"`
	Genero      string `json:"genero"`
	Codigoemail string `json:"codigoemail"`
}

func renderTemplate2(w http.ResponseWriter, tmpl string, p any) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func insertContact(dbc *sql.DB, datos consulta) {

	insertSql := "INSERT INTO arcodev.cam_user_mails(email, apellido, nombre, alta_fh, verified_code, varified_status) 		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, $4, $5)"

	_, err := dbc.Exec(insertSql, datos.Email, datos.Apellido, datos.Nombre, datos.Codigoemail, "OK")
	db.CheckError(err)
}

func presentarFormConsulta2(w http.ResponseWriter, r *http.Request) {
	//	tmpl := template.Must(template.ParseFiles("registermail.html"))
	var verifcode string
	if r.FormValue("Apellido") == "" {
		verifcode = ""
	} else {
		rand.Seed(time.Now().UnixNano())
		min := 10000
		max := 99999
		verifcode = strconv.Itoa(rand.Intn(max-min+1) + min)
	}
	details := consulta{
		Apellido:    r.FormValue("Apellido"),
		Nombre:      r.FormValue("Nombre"),
		Email:       r.FormValue("email"),
		Email2:      r.FormValue("email2"),
		Genero:      r.FormValue("genero"),
		Codigoemail: verifcode,
	}
	renderTemplate2(w, "registermail", details)

	if verifcode != "" {
		enviarMail(details.Email, "Codigo de verificación de mail", "El código generado es "+verifcode)

	}

	//	if r.Method != http.MethodPost {
	//tmpl.Execute(w, details)
	//return
	//	}

	// do something with details
	outputFile := "salidaConsulta2.txt"
	//	db.DoPostgress()

	var err error
	var jsonConsutlaTxt []byte
	jsonConsutlaTxt, err = json.MarshalIndent(details, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(outputFile, jsonConsutlaTxt, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//tmpl.Execute(w, struct{ Success bool }{true})

}

func enviarMail(mailto string, subject string, texto string) {

	msg := gomail.NewMessage()
	msg.SetHeader("From", "arcobook@gmail.com")
	msg.SetHeader("To", mailto)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", texto)
	//msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.gmail.com", 587, "arcobook@gmail.com", "hixmaeoqaxifomhs")
	//aatqmqrofcdhtsay
	// Send the email

	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

}

func registrarConsulta(w http.ResponseWriter, r2 *http.Request) {

	details := consulta{
		Apellido:    r2.FormValue("Apellido"),
		Nombre:      r2.FormValue("Nombre"),
		Email:       r2.FormValue("email"),
		Email2:      r2.FormValue("email2"),
		Genero:      r2.FormValue("genero"),
		Codigoemail: r2.FormValue("codigoemail"),
	}
	renderTemplate2(w, "verifiedmail", details)

	db := db.DbConnect()
	defer db.Close()

	insertContact(db, details)

}

/*
func homeLink(w http.ResponseWriter, _ *http.Request) {

	_, _ = fmt.Fprintf(w, "This new Site comming soon!")
}
*/

var templates = template.Must(template.ParseFiles("quiensoy.html", "testimonios.html", "libros.html", "mainpage.html", "registermail.html", "verifiedmail.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, rep any) {

	err := templates.ExecuteTemplate(w, tmpl+".html", rep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func ppalHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "mainpage", nil)
}

func librosHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "libros", nil)
}

func quiensoyHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "quiensoy", nil)
}
func testimoniosHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "testimonios", nil)
}

//var validPath = regexp.MustCompile("^/(edit|save|view|ppal|assets|registermail|mainpage)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fn(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	/*
		router := mux.NewRouter().StrictSlash(true)
		//	router.HandleFunc("/", homeLink)
		router.HandleFunc("/kitchen", homeLink)
		router.HandleFunc("/formconsulta", presentarFormConsulta)
		router.HandleFunc("/formconsulta2", presentarFormConsulta2)
		router.HandleFunc("/event", createEvent).Methods("POST")
		router.HandleFunc("/register", registrarConsulta).Methods("POST")
		router.HandleFunc("/event2", createEvent).Methods("POST")
		router.HandleFunc("/event3", createEvent).Methods("POST")
		router.HandleFunc("/events", getAllEvents).Methods("GET")

		//	http.HandleFunc("/formconsulta3", makeHandler(presentarFormConsulta2))

		//	fs := http.FileServer(http.Dir("./assets/css/"))
		//	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", fs))

		router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/css"))))
		router.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./assets/css"))))
		router.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
		router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
		router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	*/
	//	http.HandleFunc("/contactForm/", makeHandler(contactFormHandler))
	http.HandleFunc("/", makeHandler(ppalHandler))
	http.HandleFunc("/mainpage/", makeHandler(ppalHandler))
	http.HandleFunc("/libros/", makeHandler(librosHandler))
	http.HandleFunc("/quiensoy/", makeHandler(quiensoyHandler))
	http.HandleFunc("/testimonios/", makeHandler(testimoniosHandler))
	http.HandleFunc("/register/", makeHandler(presentarFormConsulta2))
	http.HandleFunc("/verified/", makeHandler(registrarConsulta))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/css"))))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./assets/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
