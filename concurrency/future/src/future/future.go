package future

// Allows us to launch other functionality when we don't 
// know when a given function will return
// This could be an on success/ on failure/ or on error
// handler that will execute while main continues to run

// Deletegates the action handler to a differet gofunc
// Stack many async calls between them 

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc	FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}