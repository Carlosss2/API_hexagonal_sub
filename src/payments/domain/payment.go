package domain

type Payment struct {
    Id    int32  `json:"id"`
    Menssage string `json:"menssage"`
}

func NewPayment( menssage string) *Payment{
	return &Payment{Id:0,Menssage:menssage}
}