package main

import (
	"fmt"
	"log"
)

// главный сервис заказов
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

	fmt.Println("Сохраняем заказ...")
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

	// Используем хранение в памяти, как вариант замены БД
	repo := NewInMemoryRepo()

	// Можно менять EmailSender на SMSSender, запустится нужный интерфейс
	notifier := &EmailSender{}

	service := NewOrderService(repo, notifier)

	err := service.CreateOrder("Семён Семёныч", []string{"Джойстик+1", "Playstation"}, 75000.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Программа успешно завершила работу!")
}
