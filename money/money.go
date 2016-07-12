package money

import "fmt"

const (
	EUR = "EUR"
	USD = "USD"
)

var Currencies = map[string]float64{
	EUR: 1.0, // default currency is EUR
	USD: 0.902730761,
}

type Money struct {
	AmountInCents int
	Currency      string
}

func (m *Money) Add(other *Money) *Money {
	return &Money{AmountInCents: m.ToEUR().AmountInCents + other.ToEUR().AmountInCents, Currency: EUR}

}

func (m *Money) AmountInEURCents() int {
	factor, ok := Currencies[m.Currency]
	if !ok {
		panic("currency " + m.Currency + " not supported")
	}
	return int(float64(m.AmountInCents) * factor)
}

func (m *Money) PerDays(n int) *Money {
	return &Money{AmountInCents: n * m.AmountInCents, Currency: m.Currency}
}

func (m *Money) ToEUR() *Money {
	return &Money{
		AmountInCents: m.AmountInEURCents(),
		Currency:      EUR,
	}
}

func New(amountInCents int, currency string) *Money {
	return &Money{AmountInCents: amountInCents, Currency: currency}
}

func (m *Money) String() string {
	return fmt.Sprintf("%.02f EUR", float64(m.ToEUR().AmountInCents)/100.0)
}
