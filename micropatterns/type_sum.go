package payment

import "fmt"

// PaymentResult — это наш закрытый интерфейс (sum type).
// Никто за пределами пакета payment не сможет создать новый вариант,
// так как метод isPaymentResult() начинается с маленькой буквы.
type PaymentResult interface {
	isPaymentResult()
}

// Вариант 1: Успешная оплата
type Success struct {
	TransactionID string
	Amount        int
}

func (Success) isPaymentResult() {}

// Вариант 2: Ошибка оплаты
type Failure struct {
	Reason string
	Code   int
}

func (Failure) isPaymentResult() {}

// Функция, возвращающая тип-сумму
func ProcessPayment(amount int) PaymentResult {
	if amount <= 0 {
		return Failure{Reason: "некорректная сумма", Code: 400}
	}
	return Success{TransactionID: "tx_98765", Amount: amount}
}

// Пример обработки всех вариантов через type switch
func HandleResult(res PaymentResult) {
	switch v := res.(type) {
	case Success:
		fmt.Printf("Оплата прошла успешно! ID: %s, Сумма: %d\n", v.TransactionID, v.Amount)
	case Failure:
		fmt.Printf("Ошибка %d: %s\n", v.Code, v.Reason)
	default:
		// В полноценных sum types блок default не нужен, но в Go
		// компилятор не проверяет полноту перебора (exhaustiveness).
		panic("неизвестный тип результата")
	}
}
