package symbol

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
	. "github.com/kocircuit/kocircuit/lang/go/kit/tree"
)

func (vty *VarietySymbol) Augment(span *Span, fields Fields) (Shape, Effect, error) {
	augmented, err := GroupFieldsToSymbols(span, fields)
	if err != nil {
		return nil, nil, err
	}
	aggregate := append(append(FieldSymbols{}, vty.Arg...), augmented...)
	if err = VerifyNoDuplicateFieldSymbol(span, aggregate); err != nil {
		return nil, nil, span.Errorf(err, "augmenting %s", Sprint(vty))
	}
	return MakeVarietySymbol(vty.Macro, aggregate), nil, nil
}

func GroupFieldsToSymbols(span *Span, fields Fields) (FieldSymbols, error) {
	evalFields := FieldSymbols{}
	for _, fieldGroup := range fields.FieldGroup() {
		groupName := fieldGroup[0].Name
		fieldGroup = FilterEmptyEvalFields(fieldGroup)
		switch len(fieldGroup) {
		case 0: // add a field with an empty symbol value (useful to All macro)
			evalFields = append(evalFields,
				&FieldSymbol{Name: groupName, Monadic: groupName == "", Value: EmptySymbol{}},
			)
		case 1:
			y := fieldGroup[0].Shape.(Symbol)
			evalFields = append(evalFields,
				&FieldSymbol{Name: groupName, Monadic: groupName == "", Value: y},
			)
		default:
			repeatedSymbols := make(Symbols, len(fieldGroup))
			for i, f := range fieldGroup {
				repeatedSymbols[i] = f.Shape.(Symbol)
			}
			if repeated, err := MakeSeriesSymbol(span, repeatedSymbols); err != nil {
				return nil, span.Errorf(err, "repeated field group %s", groupName)
			} else {
				evalFields = append(evalFields,
					&FieldSymbol{Name: groupName, Monadic: groupName == "", Value: repeated},
				)
			}
		}
	}
	return evalFields, nil
}

func FilterEmptyEvalFields(group []Field) (filtered []Field) {
	for _, field := range group {
		if !IsEmptySymbol(field.Shape.(Symbol)) {
			filtered = append(filtered, field)
		}
	}
	return
}

func VerifyNoDuplicateFieldSymbol(span *Span, fieldSymbols FieldSymbols) error {
	seen := map[string]bool{}
	for _, f := range fieldSymbols {
		if seen[f.Name] {
			return span.Errorf(nil, "augmenting duplicate field %s", f.Name)
		} else {
			seen[f.Name] = true
		}
	}
	return nil
}
