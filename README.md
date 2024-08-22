# LINEAR-STATS

This program calculates the statistical measures of a dataset provided in a text file (data.txt). The file must strictly hold decimal numbers as program won't run if a non decimal number character is detected. Also the numbers must be within the float64 range else program won't run. Empty lines will be skipped. Fractions are not to be used. The supported statistical measures are:

    - Average
    - Linear Regression Line
    - Pearson Correlation Coefficient

## Functions

The program consists of several functions to perform specific tasks in the statistical analysis process.

## Getting Started
### Dependencies

The program requires the Go programming language (version 1.16 or higher) to be installed on your system. You can install it by running the command:
```bash
go mod init linear-stats
```
### Installing

Clone the repository by running the command:
```bash
git clone https://learn.zone01kisumu.ke/git/mombewa/linear-stats.git
```
Navigate to the project directory.
```bash
cd linear-stats
```

### Executing program

Prepare a text file containing the dataset, with one number per line.
Run the compiled program with the text file as an argument:

```bash
go run main.go "data.txt" 
```

The program will output the calculated statistical measures.

### Built With

    Go - Programming language

### Author

- [@MauriceOmbewa]([https://learn.zone01kisumu.ke/git/mombewa](https://github.com/MauriceOmbewa))

### License

[MIT](LICENSE)

