package player


type ErrorNoUser struct {
    e string
}
func (err ErrorNoUser) Error() string {
    return err.e
}

type ErrorWrongPass struct {
    e string
}
func (err ErrorWrongPass) Error() string {
    return err.e
}

type ErrorUserExist struct {
    e string
}
func (err ErrorUserExist) Error() string {
    return err.e
}

