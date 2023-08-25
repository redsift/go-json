package errors

import "fmt"

type Annotation interface {
	fmt.Stringer
	Is(Annotation) bool
}

type ErrWithAnnotation struct {
	Err error
	Annotation
}

func Annotate(e error, a Annotation) error {
	if e == nil {
		return nil
	}
	return ErrWithAnnotation{
		Err:        e,
		Annotation: a,
	}
}

func (e ErrWithAnnotation) Error() string {
	if ae, ok := e.Err.(ErrWithAnnotation); ok {
		return fmt.Sprintf("%s%s", e.Annotation, ae.Error())
	}
	return fmt.Sprintf("%s: %s", e.Annotation, e.Err.Error())
}

func (e ErrWithAnnotation) Unwrap() error {
	return e.Err
}

func (e ErrWithAnnotation) Is(o error) bool {
	switch o.(type) {
	case ErrWithAnnotation:
		return true
	}
	return false
}
