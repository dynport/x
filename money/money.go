package money

import (
	"fmt"
	"strings"
	"sync"
)

const (
	EUR = "EUR"
	USD = "USD"
)

var currenciesMux = &sync.Mutex{}

func SetCurrencies(new map[string]float64) {
	currenciesMux.Lock()
	defer currenciesMux.Unlock()
	exchangeRates = new
}

func ExchangeRates() map[string]float64 {
	currenciesMux.Lock()
	defer currenciesMux.Unlock()
	return exchangeRates
}

var exchangeRates = map[string]float64{
	EUR: 1.0, // default currency is EUR
	//USD: 0.902730761,
	USD: 0.8526,
}

type Money struct {
	AmountInCents int
	Currency      string
}

func (m *Money) Add(other *Money) *Money {
	if other == nil {
		return m
	}
	return &Money{AmountInCents: m.ToEUR().AmountInCents + other.ToEUR().AmountInCents, Currency: EUR}
}

func (m *Money) Sub(other *Money) *Money {
	return &Money{AmountInCents: m.ToEUR().AmountInCents - other.ToEUR().AmountInCents, Currency: EUR}
}

func (m *Money) AmountInEURCentsWithRates(rates map[string]float64) int {
	factor, ok := rates[m.Currency]
	if !ok {
		panic("currency " + m.Currency + " not supported")
	}
	return int(float64(m.AmountInCents) * factor)
}

func (m *Money) AmountInEURCents() int {
	return m.AmountInEURCentsWithRates(ExchangeRates())
}

func (m *Money) Abs() *Money {
	amount := m.AmountInCents
	if amount < 0 {
		amount = -1 * amount
	}
	return &Money{AmountInCents: amount, Currency: m.Currency}
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
	return &Money{AmountInCents: amountInCents, Currency: strings.ToUpper(currency)}
}

func NewEUR(amountInCents int) *Money {
	return New(amountInCents, EUR)
}

func (m *Money) String() string {
	return fmt.Sprintf("%.02f EUR", float64(m.ToEUR().AmountInCents)/100.0)
}
