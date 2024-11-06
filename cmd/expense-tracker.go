package main

import (
	"expense-tracker/internals"
	"flag"
	"fmt"
	"log"
	"os"
)

const filename = "expense.json"

func main() {
	expenseList := &lib.List{}

	if err := expenseList.GetList(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
    //Add Flag
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("description", "", "Expense Desription")
	amount := addCmd.Float64("amount", 0, "Expense Desription")

    //Delete Flag
	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := delCmd.Int("id", 1, "Delete expense")

    //Summary Flag
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if err := expenseList.AddSave(filename, *desc, *amount); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Expense successfully added (ID: %d)", len(*expenseList))

	case "delete":
		delCmd.Parse(os.Args[2:])
		if err := expenseList.DeleteSave(filename, *id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Expense successfully deleted")

	case "summary":
		summaryCmd.Parse(os.Args[2:])
		expenseList.Summary(filename)

	default:
		fmt.Printf("Invalid command %s\n", os.Args[1])
	}

}
