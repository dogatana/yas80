package util

import "yas80/object"

// prog(*object.BlockObject) を単一階層の slice に変換
func FlattenObject(obj object.Object) []object.Object {
	objs := []object.Object{}

	switch obj := obj.(type) {
	case *object.BlockObject:
		for _, o := range obj.Block {
			objs = append(objs, FlattenObject(o)...)
		}
	default:
		objs = append(objs, obj)
	}
	return objs
}
