package constants

// pq error
const (
	SQLErrorDuplicate string = "pq: duplicate key"
)

// user error
const (
	UserInvalidID = 1000 + iota
	UserNotFound
	UserInvalidRequest
	UserFailedCreate
	UserFailedCreateDuplicated
)
