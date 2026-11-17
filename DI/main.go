package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// ==========================================
// 1. Слой бизнес-логики (пакет service)
// ==========================================

// Контракт (интерфейс), который ожидает сервис
type UserRepository interface {
	Save(ctx context.Context, email string) error
}

// Сам сервис с приватными зависимостями
type UserService struct {
	repo   UserRepository
	logger *slog.Logger
}

// Конструктор (аналог __construct)
func NewUserService(repo UserRepository, logger *slog.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

// Бизнес-метод
func (s *UserService) Register(ctx context.Context, email string) error {
	s.logger.Info("начинаем регистрацию пользователя", "email", email)

	if err := s.repo.Save(ctx, email); err != nil {
		s.logger.Error("ошибка сохранения в БД", "err", err)
		return err
	}

	s.logger.Info("пользователь успешно зарегистрирован", "email", email)
	return nil
}

// ==========================================
// 2. Слой инфраструктуры (пакет repository)
// ==========================================

// Конкретная реализация репозитория (например, для PostgreSQL)
type PostgresUserRepo struct {
	dsn string
}

// Конструктор репозитория
func NewPostgresUserRepo(dsn string) *PostgresUserRepo {
	return &PostgresUserRepo{dsn: dsn}
}

// Реализация метода интерфейса UserRepository
func (r *PostgresUserRepo) Save(ctx context.Context, email string) error {
	// Здесь в реальном проекте был бы SQL-запрос: INSERT INTO users ...
	fmt.Printf("[DB %s] Выполнен запрос: INSERT INTO users (email) VALUES ('%s')\n", r.dsn, email)
	return nil
}

// ==========================================
// 3. Точка сборки приложения (main.go)
// Аналог конфигурации services.yaml в Symfony
// ==========================================

func main() {
	// Шаг 1: Инициализируем базовые зависимости (логгер)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Шаг 2: Инициализируем репозиторий (уровень работы с БД)
	userRepo := NewPostgresUserRepo("postgres://localhost:5432/mydb")

	// Шаг 3: Внедряем зависимости в сервис через конструктор
	// Компилятор сам проверит, что userRepo реализует интерфейс UserRepository
	userService := NewUserService(userRepo, logger)

	// Шаг 4: Вызываем бизнес-логику (например, внутри HTTP-хендлера)
	ctx := context.Background()
	err := userService.Register(ctx, "developer@example.com")
	if err != nil {
		logger.Error("критическая ошибка", "err", err)
		os.Exit(1)
	}
}
