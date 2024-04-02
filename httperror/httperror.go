package httperror

type HttpError struct {
	Error string `json:"error"`
}

func NewHttpError(err error) HttpError {
	return HttpError{Error: err.Error()}
}

type Kind int

func (kind Kind) Error() string {
	switch kind {
	case EntityNotFound:
		return "Could not find entity in database"
	case UnAllowedToUseKey:
		return "You are not allowed to use the API-Key"
	case AuthRequired:
		return "You have to be logged in"
	case InvalidIdType:
		return "Id in path must be a valid int"
	}
	// Unreachable
	return ""
}

const (
	EntityNotFound    Kind = 1
	UnAllowedToUseKey Kind = 2
	AuthRequired      Kind = 3
	InvalidIdType     Kind = 4
)
