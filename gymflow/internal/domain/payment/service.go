package payment

type Service interface {
	CreatePayment(userID uint, req CreatePaymentRequest) (*Payment, error)
	ListPayments(userID uint) ([]Payment, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePayment(userID uint, req CreatePaymentRequest) (*Payment, error) {
	p := &Payment{
		UserID: userID,
		Amount: req.Amount,
		Method: req.Method,
		Status: "paid", // тут можно добавить интеграцию с реальным провайдером
		BookingID: req.BookingID,
	}
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *service) ListPayments(userID uint) ([]Payment, error) {
	return s.repo.ListByUser(userID)
}
