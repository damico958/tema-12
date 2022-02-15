package operations

type Subtraction struct {
	FirstNumber  float64
	OperationSymbol string
	SecondNumber float64
}

func (operation Subtraction) Calculate() float64 {
	return operation.FirstNumber - operation.SecondNumber
}

func Sub(firstNumber float64, secondNumber float64) Subtraction {
	subtract := Subtraction{firstNumber, "-", secondNumber}
	return subtract
}
