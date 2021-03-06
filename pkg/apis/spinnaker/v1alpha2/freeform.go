package v1alpha2

type FreeForm map[string]interface{}

func (f *FreeForm) DeepCopy() *FreeForm {
	cp := FreeForm{}
	copyInto(*f, cp)
	return &cp
}

func copyInto(m, cp map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			n := make(map[string]interface{})
			cp[k] = copyInto(vm, n)
		} else {
			cp[k] = v
		}
	}
	return cp
}
