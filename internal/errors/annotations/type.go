package annotations

import (
	"reflect"

	"github.com/goccy/go-json/internal/errors"
	"github.com/goccy/go-json/internal/runtime"
)

func Type(e error, t *runtime.Type) error {
	if e == nil {
		return nil
	}

	name := ""
	for {
		switch t.Kind() {
		case reflect.Pointer:
			t = t.Elem()
			continue
		default:
			name += t.Name()
		}
		break
	}
	return errors.ErrWithAnnotation{
		Err:        e,
		Annotation: AnnotationTypeName{TypeName: name},
	}
}

type AnnotationTypeName struct {
	TypeName string
}

func (a AnnotationTypeName) String() string {
	return a.TypeName
}

func (a AnnotationTypeName) Is(o errors.Annotation) bool {
	_, ok := o.(AnnotationTypeName)
	return ok
}
