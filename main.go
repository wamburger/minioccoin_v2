package main

import (
    "context"
    "log"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
    ctx := context.Background()

    // Создание экспортёра для OTLP трассировок с использованием gRPC
    exporter, err := otlptracegrpc.New(ctx,
        otlptracegrpc.WithInsecure(),               // Отключение TLS (для тестовых целей)
        otlptracegrpc.WithEndpoint("localhost:4317"), // Адрес и порт сервера OTLP
    )
    if err != nil {
        log.Fatalf("Ошибка создания экспортёра: %v", err)
    }

    // Инициализация TracerProvider с экспортёром
    tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
    defer func() { _ = tp.Shutdown(ctx) }()

    log.Println("Tracer успешно инициализирован.")
}
