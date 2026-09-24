package main

func main() {
	// ============================================================
	// ЭМУЛЯТОР БИРЖИ
	// Не изменять!
	// Генерирует тестовые котировки и отправляет их в Kafka
	// topic: tickers
	// ============================================================

	producer := NewProducer(ProducerConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "tickers",
	})

	go Run(producer)

	select {}
}
