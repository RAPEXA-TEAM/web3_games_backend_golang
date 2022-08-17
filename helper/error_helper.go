package helper

const YOU_ARE_NOT_IN_GAME = "you are not in the game"
const REQURED_PARAMETERS_IS_NOT_SET = "required parameters is not set"
const TRANSACTION_IS_PENDING = "your transaction is pending try again"
const TRANSACTION_IS_NOT_VALID = "your transaction is not valid"

func ErrorIsInErrorList(error string) bool {
	if YOU_ARE_NOT_IN_GAME == error {
		return true
	} else if TRANSACTION_IS_PENDING == error {
		return false
	} else if TRANSACTION_IS_NOT_VALID == error {
		return true
	} else {
		return false
	}
}
