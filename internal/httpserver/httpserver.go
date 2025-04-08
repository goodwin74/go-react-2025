package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Структура для управления HTTP-сервером.
type HTTPServer struct {
	port            string
	mux             *http.ServeMux
	notFoundHandler http.HandlerFunc // Обработчик для неизвестных маршрутов
	server          *http.Server
}

// Новый экземпляр HTTPServer.
func NewHTTPServer(port string) *HTTPServer {
	mux := http.NewServeMux()
	return &HTTPServer{
		port: port,
		mux:  mux,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
	}
}

// Добавляет новый маршрут с обработчиком.
func (s *HTTPServer) AddRoute(path string, handler http.HandlerFunc) {
	if path == "/" {
		// Для главной страницы используем точное совпадение
		s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				handler(w, r)
			} else {
				if s.notFoundHandler != nil {
					s.notFoundHandler(w, r)
				} else {
					http.NotFound(w, r)
				}
				return
			}
		})
	} else {
		// Обычный маршрут
		s.mux.HandleFunc(path, handler)
	}
}

// Устанавливает обработчик для неизвестных маршрутов.
func (s *HTTPServer) SetNotFoundHandler(handler http.HandlerFunc) {
	s.notFoundHandler = handler
}

// Запускает HTTP-сервер.
func (s *HTTPServer) Start() error {
	fmt.Printf("Starting HTTP server on port %s...\n", s.port)
	return s.server.ListenAndServe()
}

// Останавливает HTTP-сервер.
func (s *HTTPServer) Stop() error {
	fmt.Println("Stopping HTTP server...")
	return s.server.Close()
}

// Обработчик для обслуживания статических файлов
func (s *HTTPServer) ServeStatic(dir string) {

	fs := http.FileServer(http.Dir(dir))

	// Регистрируем обработчик для всех маршрутов
	s.mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Если запрос идет на главную страницу ("/"), перенаправляем на index.html
		if r.URL.Path == "/" {
			http.ServeFile(w, r, dir+"/index.html")
			return
		}

		// Иначе передаем управление FileServer для обслуживания статических файлов
		fs.ServeHTTP(w, r)
	}))
}

func (s *HTTPServer) ProxyToVite() {
	// Создаем прокси-обработчик для localhost:3001
	viteURL, _ := url.Parse("http://localhost:3001")
	proxy := httputil.NewSingleHostReverseProxy(viteURL)

	// Регистрируем обработчик для всех маршрутов
	s.mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}))
}

// Добавляет роутер для React приложения из папки dist на главную страницу.
func (s *HTTPServer) AddReactRouter() {
	// Получаем значение переменной окружения "APPMODE"
	appMode := os.Getenv("APPMODE")

	if appMode != "dev" {
		// Если значение не равно "dev", обслуживаем статические файлы из папки "dist"
		s.ServeStatic("./web/dist")
	} else {
		// Если значение равно "dev", запускаем сервер на порту 8080 и перенаправляем запросы на Vite
		fmt.Println("Running in development mode. Proxying requests to Vite...")
		s.ProxyToVite()
	}
}
