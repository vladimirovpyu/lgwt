package main

import (
	"context"
	"log/slog"
	"os"

	"go.uber.org/fx"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func main() {
	fx.New(
		// Регистрируем конструкторы (аналог autowire в services.yaml)
		fx.Provide(
			NewLogger,
			func() string { return "postgres://localhost:5432/mydb" },
			fx.Annotate(
				NewPostgresUserRepo,
				fx.As(new(UserRepository)), // связываем реализацию с интерфейсом
			),
			NewUserService,
		),
		// Запускаем код, куда fx сам подставит готовый *UserService
		fx.Invoke(func(svc *UserService) {
			_ = svc.Register(context.Background(), "dev@example.com")
		}),
	).Run()
}
