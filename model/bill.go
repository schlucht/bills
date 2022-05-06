package bill

import 
(
  "fmt"
  "time"
)

type Bill struct {
  IBAN string
  Booked time.Time
  Text string
  Amount float32
  Ballance float32
}

func (m *Bill) ToString() string {
  return fmt.Sprintf("%s\t, %.2f\t\t, %.2f\t\t", m.Text, m.Amount, m.Ballance)
}