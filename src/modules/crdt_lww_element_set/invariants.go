package mod_crdt_lww_element_set

import "reflect"

func VerifyLWWLaws(a, b, c *LWWSet) error {
	ab1 := a.Clone()
	if err := ab1.Merge(b); err != nil {
		return err
	}
	ba1 := b.Clone()
	if err := ba1.Merge(a); err != nil {
		return err
	}
	if !reflect.DeepEqual(ab1.Elements(), ba1.Elements()) {
		return ErrLWWInvariant
	}
	dup := ab1.Clone()
	if err := dup.Merge(ab1); err != nil || !reflect.DeepEqual(dup.Elements(), ab1.Elements()) {
		return ErrLWWInvariant
	}
	left := ab1.Clone()
	if err := left.Merge(c); err != nil {
		return err
	}
	bc := b.Clone()
	if err := bc.Merge(c); err != nil {
		return err
	}
	right := a.Clone()
	if err := right.Merge(bc); err != nil || !reflect.DeepEqual(left.Elements(), right.Elements()) {
		return ErrLWWInvariant
	}
	return nil
}
