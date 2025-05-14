package main

func DoSomething(input int) (err error) {

	defer func() {
		p := recover()
		if p == nil {
			return
		}

		if e, ok := p.(error); ok {
			err = e
			return
		}

		//		err := fmt.Errorf("panit %T %s", p, p)
	}()

	switch input {
	case 0:
		return nil
	case 1:
		panic("one")
	}

	return nil
}

func main() {

}
