# expense-tracker
A simple application to track and manager your finances.<br />
A Sample solution for the [expense-tracker](https://https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

# Build
Clone the repository and run the following command:
```bash
git clone https://github.com/OLABALADE/expense-tracker.git
cd expense-tracker
```
Build the program with the following command:
```bash
go build cmd/expense-tracker
```
# Usage
Add expense
```bash
#expense-tracker add --description <description> --amount <amount>
expense-tracker add --description "Breakfast" --amount 10
#Expense successfully added (ID: 1)
```
Delete expense
```bash
#expense-tracker delete --id <expense id>
expense-tracker delete --id 1
#Expense successfully deleted
```
Expense Summary
```bash
expense-tracker add --description "Lunch" --amount 20
expense-tracker add --description "Dinner" --amount 10
expense-tracker summary
# ID  Date       Description  Amount
# 1   2024-11-06  Lunch        $20
# 2   2024-11-06  Dinner       $10
