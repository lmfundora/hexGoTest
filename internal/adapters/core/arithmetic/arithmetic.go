package arithmetic

type Adapter struct {
  
}

func NewAdapter() *Adapter  {
  return &Adapter{}
}

func (adith Adapter) Addition(a, b int32) (int32, error) {
  return a + b, nil
}

func (adith Adapter) Subtraction(a, b int32) (int32, error) {
  return a - b, nil
}

func (adith Adapter) Multiplication(a, b int32) (int32, error) {
  return a * b, nil
}

func (adith Adapter) Division(a, b int32) (int32, error) {
  return a / b, nil
}
