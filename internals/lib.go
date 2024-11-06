package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Expense struct {
	Description string
	Amount      float64
	Date        time.Time
}

type List []Expense

func (l *List) AddSave(filename, desc string, amt float64) error {
	new_expense := Expense{
		Description: desc,
		Amount:      amt,
		Date:        time.Now(),
	}

	*l = append(*l, new_expense)

	if err := l.SaveList(filename); err != nil {
		return err
	}

	return nil
}

func (l *List) DeleteSave(filename string, id int) error {
	lp := *l
	if id <= 0 || id > len(*l) {
		return fmt.Errorf("ID does not exist")
	}

	*l = append(lp[:id-1], lp[id:]...)

	if err := l.SaveList(filename); err != nil {
		return err
	}

	return nil
}

func (l *List) GetList(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)

}

func (l *List) SaveList(filename string) error {
	jsfile, err := json.Marshal(l)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsfile, 0644)
}

func (l *List) Summary(filename string) {
	err := l.GetList(filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(*l) != 0 {
		fmt.Println(" ID\tDate\t\tDescription\tAmount")
		for index, exp := range *l {
			fmt.Printf("%3d\t%s\t%s\t$%3.2f\n", index+1, exp.Date.Format("2006-01-02"), exp.Description, exp.Amount)
		}
	} else {
		fmt.Println("Expense List is empty.")
	}
}
