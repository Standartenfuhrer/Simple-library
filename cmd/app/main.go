package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//Регистрируем обработчик для пути "/"
	//Все запросы на адрес http://localhost/8080 будут обрабатываться функцией homeHandler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health_check", healthCheckHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/notfound", notFoundHandler)
	
	//Определяем порт, на котором будет работать сервер
	port := ":8080"
	fmt.Printf("Запускаем сервер на порту %s", port)

	//Запускаем сервер
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err) //Если сервер не запустился - логируем и выходим
	}

}

// Наш первый обработчик
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Library"))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//1. Создаем структуру для ответа, чтобы ее было легко превратить в JSON
	response := struct {
		Status      string `json:"status"`
		ProjectName string `json:"project_name"`
	}{
		Status:      "available",
		ProjectName: "LibraryAPI",
	}

	//2. Сериаулизуем структуру в JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//3. Устанавливаем заголовок
	w.Header().Set("Content-Type", "application/json")

	//4. Устанавливаем статус код 200 ОК
	w.WriteHeader(http.StatusOK)

	//5. Пишем JSON для ответа
	w.Write(jsonData)
}

//Вывод информации
func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status     string `json:"status"`
		Version    string `json:"version"`
		ServerTime string `json:"server_time"`
	}{
		Status:     "available",
		Version:    "1.0",
		ServerTime: time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	w.Write(jsonData)
}
//Не найдена страница
func notFoundHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Страница не найдена"))
}
