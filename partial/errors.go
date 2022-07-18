package partial

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Errors map[string]error

func (es Errors) Error() string {
	return validation.Errors(es).Error()
}

// Filter removes all nils from Errors and returns back the updated Errors as an error.
// If the length of Errors becomes 0, it will return nil.
func (es Errors) Filter() error {
	err := validation.Errors(es).Filter()
	if err != nil {
		if errs, ok := err.(validation.Errors); ok {
			return Errors(errs)
		}
	}
	return err
}

func FilterErrors(err error, fields []string) error {
	fieldMap := make(map[string]bool)
	for _, field := range fields {
		fieldMap[field] = true
	}
	errMap, ok := err.(Errors)
	if !ok {
		return err
	}
	for field := range errMap {
		if !fieldMap[field] {
			errMap[field] = nil
		}
	}
	return errMap.Filter()
}
