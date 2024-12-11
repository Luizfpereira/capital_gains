# Capital Gains

## Description
This Go CLI program calculates the capital gains tax based on buy and sell stock operations in the financial market. The program reads input as JSON-formatted transactions, processes the data, and calculates the applicable tax.

## How it works

1. Input:

    The program reads multiple JSON-formatted lists of stock operations from standard input.
    Each operation contains:
    operation: "buy" or "sell".

2. unit-cost: 

    The price per unit of the stock.
    quantity: The number of stocks in the transaction.


    For every sell operation, calculates the profit:
    profit = (sell price - buy price) × quantity.
    Exempts taxes if total monthly sales are below R$ 20,000.00.
    Applies a 20% tax rate on profits exceeding R$ 20,000.00.

3. Output:

    The program outputs the calculated tax or indicates no tax is due.

### Example

The user will input multiple lists of JSON-formatted objects and hit enter to indicate the end of the input. To get the result, the user will hit the enter again.

**Input:**

```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":20.00, "quantity": 5000}]
[{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
{"operation":"sell", "unit-cost":10.00, "quantity": 5000}]

```

**Output:**

```json
[{"tax":0},{"tax":10000}]
[{"tax":0},{"tax":0}]
```

For every list informed by the user, a new list in a new line will be printed in the terminal. A tax object will exist for every operation, buy or sell.

---

# Requirements
* **Go**: Version 1.18 or higher
* **Docker**

---

# Installation and Usage

## Clone the repository:

```bash
git clone https://github.com/Luizfpereira/capital_gains.git
cd capital_gains

```

## Build and run the program:

Go to the /cmd folder:

```shell
cd cmd
```

Then, run the command:
```shell
go run main.go
```

The terminal will be ready to receive user's inputs.

### Input format

Provide stock operations as JSON-formatted lists, one per line:

```shell
[{"operation":"buy", "unit-cost":10.00, "quantity": 100},
{"operation":"sell", "unit-cost":15.00, "quantity": 50},
{"operation":"sell", "unit-cost":15.00, "quantity": 50}]
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":20.00, "quantity": 5000},
{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]
```

## Run with docker:

After cloning the repository, inside the capital_gains folder, build the application:


```shell
docker build -t my-app .
```

Execute the container with the command:

```shell
docker run --rm -it my-app
```

You will be redirected to the **/cmd** folder inside the container. Then, run the application

```shell
go run main.go
```

To leave the container, type **exit**.