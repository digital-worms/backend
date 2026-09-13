package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectPostgres создаёт пул соединений с PostgreSQL и проверяет,
// что база данных действительно доступна.
func ConnectPostgres(connStr string) (*pgxpool.Pool, error) {
	// Ограничиваем время всей инициализации пула пятью секундами,
	// чтобы приложение не зависло навсегда при недоступной базе данных.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Разбираем строку подключения и преобразуем её в конфигурацию pgxpool.
	// На этом этапе проверяется формат настроек, но соединение с БД ещё не подтверждено.
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse PostgreSQL config: %w", err)
	}

	// Ограничиваем максимальное количество соединений и просим пул
	// поддерживать хотя бы одно соединение готовым к использованию.
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 1

	// Создаём пул на основе подготовленной конфигурации.
	// Успешное создание объекта пула ещё не гарантирует доступность PostgreSQL.
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create PostgreSQL connection pool: %w", err)
	}

	// Выполняем реальную проверку соединения. Если Ping завершается ошибкой,
	// закрываем уже созданный пул, чтобы не оставить занятые ресурсы.
	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	// Возвращаем готовый пул вызывающему коду. Его владелец обязан
	// вызвать Close, когда приложение завершает работу.
	return pool, nil
}
