package user

type MaritalStatus string

const (
	Married MaritalStatus = "married"
	Single  MaritalStatus = "single"
)

func ValidateMaritalStatus(status string) error {
	switch status {
	case string(Married), string(Single):
		return nil
	default:
		return ErrInvalidMaritalStatus
	}
}
