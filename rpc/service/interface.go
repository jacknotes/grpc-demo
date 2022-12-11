package service

const (
	Name = "HelloService"
)

type Service interface {
	Hello(name string, response *string) error
}
