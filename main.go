package main

import (
	"fmt"
	"log"
)

// OrderService - наш главный сервис
type OrderService struct {
	repo     RepositoryWriter
	notifier Notifier
}

func NewOrderService(repo RepositoryWriter, notifier Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := Order{
		Customer: customer,
		Products: fmt.Sprintf("%v", products),
		Total:    total,
		Status:   "pending",
	}

	fmt.Println("Сохраняем заказ в хранилище...")
	if err := s.repo.SaveOrder(order); err != nil {
		return fmt.Errorf("ошибка сохранения: %w", err)
	}

	msg := fmt.Sprintf("Ваш заказ на %.2f руб. успешно создан!", total)
	fmt.Println("Отправляем уведомление...")
	s.notifier.Send(customer, msg)

	return nil
}

func main() {
	fmt.Println("Запуск системы управления заказами...")

	// Используем хранилище в памяти - никаких баз данных и компиляторов C не нужно!
	repo := NewInMemoryRepo()

	// Попробуй поменять EmailSender на SMSSender, чтобы увидеть разницу
	notifier := &EmailSender{}

	service := NewOrderService(repo, notifier)

	err := service.CreateOrder("Семён Семёныч", []string{"Джойстик+1", "Playstation"}, 75000.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Программа успешно завершила работу!")
}
