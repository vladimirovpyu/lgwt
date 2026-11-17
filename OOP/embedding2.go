package main

// это пример переопределения метода

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

// 1. Встраиваемая структура с общими методами
type BaseHandler struct {
	Logger *slog.Logger
}

func (b *BaseHandler) RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		b.Logger.Error("failed to encode response", "err", err)
	}
}

func (b *BaseHandler) LogRequest(r *http.Request) {
	b.Logger.Info("incoming request", "method", r.Method, "path", r.URL.Path)
}

// 2. Внешняя структура, которая встраивает BaseHandler
type UserHandler struct {
	BaseHandler // Встраивание (анонимное поле без имени)
	ServiceName string
}

// Собственный метод UserHandler, который использует «поднятые» методы BaseHandler
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Вызываем метод BaseHandler напрямую через h, как будто он объявлен у UserHandler
	h.LogRequest(r)

	user := map[string]string{"name": "Alex", "service": h.ServiceName}
	h.RespondJSON(w, http.StatusOK, user)
}

// 3. Переопределение (затенение / shadowing) метода
func (h *UserHandler) LogRequest(r *http.Request) {
	fmt.Println("--- Custom UserHandler Log ---")
	// При необходимости можно вызвать «родительский» метод явно по имени типа
	h.BaseHandler.LogRequest(r)
}

func main() {
	handler := &UserHandler{
		BaseHandler: BaseHandler{
			Logger: slog.Default(),
		},
		ServiceName: "users-api",
	}

	// Вызовется переопределенный метод UserHandler.LogRequest
	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	handler.LogRequest(req)

	// Вызовется поднятый метод BaseHandler.RespondJSON
	// Также UserHandler автоматически реализует любой интерфейс, требующий RespondJSON(...)
}
