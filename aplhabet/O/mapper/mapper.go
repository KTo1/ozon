package mapper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func CopyFieldByPath(src interface{}, srcPath string, dst interface{}, dstPath string) error {
	srcSegs, err := parsePath(srcPath)
	if err != nil {
		return fmt.Errorf("source path parse error: %w", err)
	}

	dstSegs, err := parsePath(dstPath)
	if err != nil {
		return fmt.Errorf("destination path parse error: %w", err)
	}

	srcVal, err := getOrCreateFieldBySegments(src, srcSegs)
	if err != nil {
		return fmt.Errorf("source value error: %w", err)
	}

	dstVal, err := getOrCreateFieldBySegments(dst, dstSegs)
	if err != nil {
		return fmt.Errorf("destination value error: %w", err)
	}

	if !dstVal.CanSet() {
		return fmt.Errorf("cannot set destination field")
	}

	if srcVal.Type() != dstVal.Type() {
		if srcVal.Type().ConvertibleTo(dstVal.Type()) {
			srcVal = srcVal.Convert(dstVal.Type())
		} else {
			return fmt.Errorf("cannot convert %v to %v", srcVal.Type(), dstVal.Type())
		}
	}

	dstVal.Set(srcVal)

	return nil
}

// private

type pathSegment struct {
	Field string
	Index *int // если это слайс, то индекс
}

func parsePath(path string) ([]pathSegment, error) {
	var result []pathSegment

	segments := strings.Split(path, ".")
	for _, seg := range segments {
		var ps pathSegment

		bracketIdx := strings.Index(seg, "[")
		if bracketIdx == -1 {
			ps.Field = seg
		} else {
			if !strings.HasSuffix(seg, "]") {
				return nil, fmt.Errorf("invalid segment: %s", seg)
			}

			ps.Field = seg[:bracketIdx]
			idxStr := seg[bracketIdx+1 : len(seg)-1]
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return nil, fmt.Errorf("invalid index in segment '%s': %w", seg, err)
			}

			ps.Index = &idx
		}

		result = append(result, ps)
	}

	return result, nil
}

func getOrCreateFieldBySegments(object interface{}, segments []pathSegment) (reflect.Value, error) {
	v := reflect.ValueOf(object)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return reflect.Value{}, fmt.Errorf("expected non-nil pointer to struct")
	}

	v = v.Elem()
	for i, seg := range segments {
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				v.Set(reflect.New(v.Type().Elem()))
			}

			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return reflect.Value{}, fmt.Errorf("segment %d: %s is not a struct", i, seg.Field)
		}

		v = v.FieldByName(seg.Field)
		if !v.IsValid() {
			return reflect.Value{}, fmt.Errorf("field %q not found", seg.Field)
		}

		if seg.Index != nil {
			if v.Kind() != reflect.Slice {
				return reflect.Value{}, fmt.Errorf("field %q is not a slice", seg.Field)
			}

			if v.Len() <= *seg.Index {
				needed := *seg.Index + 1 - v.Len()
				extension := reflect.MakeSlice(v.Type(), needed, needed)
				v.Set(reflect.AppendSlice(v, extension))
			}

			v = v.Index(*seg.Index)
		}
	}

	return v, nil
}
