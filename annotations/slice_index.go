package annotations

import (
	"fmt"

	"github.com/goccy/go-json/internal/errors"
)

func SliceIndex(e error, idx int) error {
	if e == nil {
		return nil
	}
	return errors.ErrWithAnnotation{
		Err:        e,
		Annotation: AnnotationSliceIndex{Idx: idx},
	}
}

type AnnotationSliceIndex struct {
	Idx int
}

func (a AnnotationSliceIndex) String() string {
	return fmt.Sprintf("[%d]", a.Idx)
}

func (a AnnotationSliceIndex) Is(o errors.Annotation) bool {
	_, ok := o.(AnnotationSliceIndex)
	return ok
}
