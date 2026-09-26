# gledger
Double Entry Accounting TUI

## VP Goals:

Focus strictly on transaction entry and simple balance reporting

### View 1: Transaction Ledger (The Dashboard)
It reads an underlying ledger file (like a plain-text journal) and displays recent history.

Simple Transaction List: A vertical list showing date, description, and the total amount of the transaction. You can use the bubbles/list component or a simple viewport.

Running Account Balances: A side panel or top banner displaying the current totals for major account types (Assets, Liabilities, Equity, Income, Expenses).

### View 2: Transaction Entry Form (The Core Utility)
Double-entry requires at least two postings per transaction.

Multi-Field Input Forms: Use bubbles/textinput to capture the Date, Description, Account 1, Amount 1, Account 2, and Amount 2.

Real-time Balance Validation: A visual indicator (like a green checkmark or red warning text) showing whether Amount 1 + Amount 2 == 0.

Auto-balancing Shortcut: If the user enters the first amount (e.g., $50), the second amount field should automatically populate with the inverse (-$50) to save keystrokes.
