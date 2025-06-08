package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// HandleRoot отправляет html форму из index.html
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// Получаем путь до текущей рабочей директории
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "не удалось получить рабочую директорию", http.StatusInternalServerError)
		return
	}

	// Собираем путь до index.html на уровень выше
	path := filepath.Join(wd, "..", "index.html")

	// Отдаем файл как ответ
	http.ServeFile(w, r, path)
}

// HandleUpload обрабатывает загрузку файла
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // ограничение на размер файла: 10 МБ
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
