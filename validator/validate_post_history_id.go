package validator

func ValidatePostHistoryId(value int32) error {
	return ValidateInt(value, 1)
}