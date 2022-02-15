package operations

type Addition struct {
	FirstNumber  float64
	OperationSymbol string
	SecondNumber float64
}

func (operation Addition) Calculate() float64 {
	return operation.FirstNumber + operation.SecondNumber
}

func Add(firstNumber float64, secondNumber float64) Addition {
	sum := Addition{firstNumber, "+", secondNumber}
	return sum
}
