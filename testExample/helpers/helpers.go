package helpers


type Helpers interface {
	Sum(num1, num2 int) int
	SumMultiple(num2 []int) int
}

type helpersImpl struct{}

func NewHelpers() Helpers {
	return &helpersImpl{}
}

func (h *helpersImpl) Sum(num1, num2 int) int {
 return num1 + num2
}

func (h *helpersImpl) SumMultiple(num []int) int {
	acum := 0

	for _, n := range num {
		acum = acum + n
	}

	return acum
}