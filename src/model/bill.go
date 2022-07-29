package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const timeLayout = "2006-01-02"

type Bill struct {
	IBAN     string
	Booked   time.Time
	Text     string
	Ballance float32
	Out      float32
	In       float32
}

func (m *Bill) String() string {
	return fmt.Sprintf(
		"IBAN:\t\t%v\nBooked:\t\t%v\nText:\t\t%v\nBallance:\t%v\nOut:\t\t%v\nIn:\t\t%v\n",
		m.IBAN, m.Booked, m.Text, m.Ballance, m.Out, m.In,
	)
}

type Bills struct {
	Bills  []Bill
	Sum    float32
	SumOut float32
	SumIn  float32
}

func (m *Bills) String() string {
	return fmt.Sprintf("Anzahl Daten: %v\n Summe: %v\n", len(m.Bills), m.Sum)
}

func (m *Bills) FilterMonth() map[string]Bills {
	bs := make(map[string]Bills)
	b := []Bill{}

	for month := 1; month <= 12; month++ {
		monthName := ""
		for _, bill := range m.Bills {
			if int(bill.Booked.Month()) == month {
				b = append(b, bill)
				monthName = bill.Booked.Month().String()
			}
		}

		bs[monthName] = initBills(b)
	}
	return bs
}

func initBills(bills []Bill) Bills {
	if len(bills) == 0 {
		return Bills{}
	}
	bs := Bills{Bills: bills}
	for _, bill := range bills {
		bs.SumIn += bill.In
		bs.SumOut += bill.Out
	}
	bs.Sum = bs.SumIn - bs.SumOut
	return bs
}

func New(csvFile string) (Bills, error) {
	bs := Bills{}
	lines, err := readCsv("data/bisSeptember.csv")
	if err != nil {
		panic("Datei konnte nicht gelesen werden")
	}
	bsLines := bills(lines)
	bs = bsLines

	bs.Sum = bs.SumIn - bs.SumOut

	return bs, nil
}

// Liest die CSV Datei
func readCsv(fs string) ([][]string, error) {

	f, err := os.Open(fs)
	if err != nil {
		return [][]string{}, err
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return data, nil
}

// Liest die Daten in das Bills Objekt
func bills(data [][]string) Bills {
	bs := []Bill{}
	bills := Bills{
		SumIn:  0,
		SumOut: 0,
		Sum:    0,
		Bills:  []Bill{},
	}
	for _, lines := range data[1:] {
		b := Bill{}
		if lines[1] == "" {
			continue
		}
		b.IBAN = lines[0]
		date, err := time.Parse(timeLayout, strings.Split(lines[1], " ")[0])
		if err != nil {
			continue
		}
		b.Booked = date
		b.Text = lines[2]
		amout, err := strconv.ParseFloat(lines[3], 32)

		if err != nil {
			log.Fatal(err)
			continue
		}
		if amout > 0 {
			b.In = float32(amout)
			bills.SumIn += b.In
		} else {
			b.Out = float32(amout)
			bills.SumOut += b.Out
		}
		ballance, err := strconv.ParseFloat(lines[3], 32)
		if err != nil {
			log.Fatal(err)
			continue
		}
		b.Ballance = float32(ballance)
		// fmt.Println(b.Text)
		bs = append(bs, b)
	}
	bills.Bills = bs
	return bills
}
