package domain

type IPayment interface{
	Save(menssage string) error
}