package main

import "fmt"

type Request struct {
	URL    string
	UserID int
	Body   string
}

type Response struct {
	Status  int
	Message string
}

// Тип обработчика: функция, которая принимает запрос и следующий обработчик
type Handler func(req *Request) *Response

// Middleware — функция, которая оборачивает обработчик
type Middleware func(next Handler) Handler

// Middleware 1: логирование
func LoggingMiddleware(next Handler) Handler {
	return func(req *Request) *Response {
		fmt.Printf("📝 [LOG] Запрос: %s\n", req.URL)
		resp := next(req)
		fmt.Printf("📝 [LOG] Ответ: %d\n", resp.Status)
		return resp
	}
}

// Middleware 2: аутентификация
func AuthMiddleware(next Handler) Handler {
	return func(req *Request) *Response {
		if req.UserID == 0 {
			fmt.Println("🚫 [AUTH] Пользователь не авторизован")
			return &Response{Status: 401, Message: "Unauthorized"}
		}
		fmt.Printf("✅ [AUTH] Пользователь %d авторизован\n", req.UserID)
		return next(req)
	}
}

// Middleware 3: проверка валидности запроса
func ValidationMiddleware(next Handler) Handler {
	return func(req *Request) *Response {
		if req.Body == "" {
			return &Response{Status: 400, Message: "Bad Request: body is empty"}
		}
		fmt.Println("✅ [VALID] Запрос прошёл валидацию")
		return next(req)
	}
}

// Основной обработчик
func MainHandler(req *Request) *Response {
	fmt.Printf("🎯 [MAIN] Обрабатываем запрос: %s\n", req.URL)
	return &Response{Status: 200, Message: "OK"}
}

// Функция для применения цепочки middleware
func Apply(handler Handler, middlewares ...Middleware) Handler {
	// Применяем в обратном порядке, чтобы первый middleware выполнился первым
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func main() {
	// Выстраиваем цепочку: Logging → Auth → Validation → Main
	chain := Apply(MainHandler, LoggingMiddleware, AuthMiddleware, ValidationMiddleware)

	fmt.Println("=== Запрос без авторизации ===")
	resp := chain(&Request{URL: "/api/data", UserID: 0, Body: "test"})
	fmt.Printf("Результат: %d %s\n\n", resp.Status, resp.Message)

	fmt.Println("=== Авторизованный запрос ===")
	resp = chain(&Request{URL: "/api/data", UserID: 42, Body: "hello"})
	fmt.Printf("Результат: %d %s\n", resp.Status, resp.Message)
}
