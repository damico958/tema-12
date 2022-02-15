package operations

type Multiplication struct {
	FirstNumber  float64
	OperationSymbol string
	SecondNumber float64
}

func (operation Multiplication) Calculate() float64 {
	return operation.FirstNumber * operation.SecondNumber
}

func Mul(firstNumber float64, secondNumber float64) Multiplication {
	multiply := Multiplication{firstNumber, "*" , secondNumber}
	return multiply
}
