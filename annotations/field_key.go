package annotations

import "github.com/goccy/go-json/internal/errors"

func FieldKey(e error, key string) error {
	if e == nil {
		return nil
	}
	return errors.ErrWithAnnotation{
		Err:        e,
		Annotation: AnnotationFieldKey{Key: key},
	}
}

type AnnotationFieldKey struct {
	Key string
}

func (a AnnotationFieldKey) String() string {
	return "." + a.Key
}

func (a AnnotationFieldKey) Is(o errors.Annotation) bool {
	_, ok := o.(AnnotationFieldKey)
	return ok
}
