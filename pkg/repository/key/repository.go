package key

type Reader interface {
	Get() error
}

type Writer interface {
	//Store() (int64, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
