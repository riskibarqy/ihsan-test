package constants

// custom error
const (
	UnknownError   string = "internal server error"
	RecordNotFound string = "record not found"
)

// pq error
const (
	SQLErrorDuplicate string = "pq: duplicate key"
)

// general error
const (
	InternalServerError = 1000 + iota
)

// user error
const (
	UserInvalidID = 2000 + iota
	UserNotFound
	UserInvalidRequest
	UserFailedCreateDuplicated
	UserGetBalanceNotFound
)

// user balance history error
const (
	UserBalanceHistoryInvalidRequest = 3000 + iota
	UserBalanceHistoryNotFound
	UserBalanceHistoryInvalidWithdrawAmount
)
