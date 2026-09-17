package main

import "fmt"


type RepositoryWriter interface {
	SaveOrder(order Order) error
}

type Notifier interface {
	Send(customer string, message string) error
}

type InMemoryRepo struct {
	Orders []Order
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{}
}

func (r *InMemoryRepo) SaveOrder(order Order) error {
	r.Orders = append(r.Orders, order)
	fmt.Printf("Заказ сохранен в память. Всего заказов: %d\n", len(r.Orders))
	return nil
}

// уведомления

type EmailSender struct{}

func (e *EmailSender) Send(customer string, message string) error {
	fmt.Printf("[EMAIL] Клиенту %s: %s\n", customer, message)
	return nil
}

type SMSSender struct{}

func (s *SMSSender) Send(customer string, message string) error {
	fmt.Printf("[SMS] Клиенту %s: %s\n", customer, message)
	return nil
}
