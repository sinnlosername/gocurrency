package gocurrency

type Money struct {
	Value float64
	Currency string
}

func (m *Money) ToEUR() {
	m.Value = ConvertToEUR(m.Currency, m.Value)
}

func (m *Money) To(currency string) {
	m.Value = ConvertFromTo(m.Currency, currency, m.Value)
	m.Currency = currency
}

func (m *Money) Format(precision int) (string) {
	return FormatNumber(m.Value, precision)
}

func (m *Money) IsValid() (bool) {
	return IsNumberValid(m.Value)
}