package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"testApi/models"
)

type Usuario struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	Nombre        string
	Nombreusuario string
	Password      string
	Tipousuario   string
	Inventario    string
}

type UserResponse struct {
	Token      string
	Mensaje    string
	Id         string
	Nombre     string
	Inventario string
	Tipo       string
	Alias      string
}

var LoginUser = func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	nombreUsarioForm := r.Form.Get("usuario")
	contrasenaForm := r.Form.Get("contrasena")

	client := models.GetClient()
	usuariosCollection := client.Database("GMDS").Collection("Usuarios")

	var result Usuario

	filter := bson.D{{"nombreusuario", nombreUsarioForm}}

	err := usuariosCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	fmt.Printf("Found a single document: %+v\n", result)

	userPassword2 := contrasenaForm
	hasFromDatabase := []byte(result.Password)

	if err := bcrypt.CompareHashAndPassword(hasFromDatabase, []byte(userPassword2)); err != nil {
		w.WriteHeader(404)
		w.Write([]byte("El nombre de usuario o contraseña no son validas"))
		return
	}

	fmt.Println("Password was correct")

	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}

	clave := os.Getenv("CLAVE_SECRETA")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Usuario": result.Nombreusuario,
		"pass":    os.Getenv("CLAVE_TOKENS"),
		"nbf":     time.Date(2019, 7, 20, 10, 25, 15, 10, time.UTC).Unix(),
	})

	var jwtKey = []byte(clave)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	respuesta := UserResponse{tokenString, "Autenticacion exitosa", result.Id.Hex(), result.Nombre, result.Inventario, result.Tipousuario, result.Nombreusuario}

	js, err := json.Marshal(respuesta)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
