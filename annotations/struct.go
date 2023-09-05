package annotations

import "github.com/goccy/go-json/internal/errors"

func StructName(e error, structName string) error {
	if e == nil {
		return nil
	}
	return errors.ErrWithAnnotation{
		Err:        e,
		Annotation: AnnotationStructName{StructName: structName},
	}
}

type AnnotationStructName struct {
	StructName string
}

func (a AnnotationStructName) String() string {
	return a.StructName
}

func (a AnnotationStructName) Is(o errors.Annotation) bool {
	_, ok := o.(AnnotationStructName)
	return ok
}
