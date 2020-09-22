package libapi

type ActionResult struct {
	Result interface{}
}

func ObjResult(retVal interface{}) (*ActionResult, error) {
	return &ActionResult{
		Result: retVal,
	}, nil
}

func ErrResult(errorEnc error) (*ActionResult, error) {
	return nil, errorEnc
}

func NilResult() (*ActionResult, error) {
	return &ActionResult{
		Result: nil,
	}, nil
}
