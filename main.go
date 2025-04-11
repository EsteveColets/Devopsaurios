package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestType := r.Method
	log.Printf("Received %s request for %s", r.Method, r.URL.Path)
	fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Página de ejemplo</title>
            <style>
                body {
                    background-color:rgb(40, 43, 240); /* Azul claro */
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                    height: 100vh;
                    margin: 0;
                    font-family: Arial, sans-serif;
                }
                h1 {
                    color: white;
                    text-align: center;
                }
                img {
                    display: block;
                    margin: 20px auto;
                }
                .button {
                    background-color: #4CAF50; /* Verde */
                    border: none;
                    color: white;
                    padding: 15px 32px;
                    text-align: center;
                    text-decoration: none;
                    display: inline-block;
                    font-size: 16px;
                    margin: 4px 2px;
                    cursor: pointer;
                }
            </style>
        </head>
        <body>
            <h1>Soy alumno de la UOC</h1>
            <p>Tengo una petición tipo: %s</p>
            <a href="https://sites.google.com/uoc.edu/devopsauriosoficial/miembros-del-equipo" class="button" target="_blank">DevOpsaurios</a>
            <img src="static/logo.jpeg" alt="Imagen de ejemplo">
        </body>
        </html>
    `, requestType)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
