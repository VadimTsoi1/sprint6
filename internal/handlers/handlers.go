package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HandleRoot отдаёт index.html, который лежит на уровень выше от internal/handlers
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "не удалось получить рабочую директорию", http.StatusInternalServerError)
		return
	}

	path := filepath.Join(wd, "..", "index.html")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path)
}

// HandleUpload обрабатывает загрузку файла и конвертацию
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка при разборе формы", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Ошибка при конвертации", http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().Format("20060102150405") + ".txt"
	fullPath := filepath.Join(".", filename)

	err = os.WriteFile(fullPath, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Ошибка при записи файла", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
