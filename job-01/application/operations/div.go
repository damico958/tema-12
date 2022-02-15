package operations

type Division struct {
	FirstNumber  float64
	OperationSymbol string
	SecondNumber float64
}

func (operation Division) Calculate() float64 {
	return operation.FirstNumber / operation.SecondNumber
}

func Div(firstNumber float64, secondNumber float64) Division {
	division := Division{firstNumber, "/" ,secondNumber}
	return division
}
