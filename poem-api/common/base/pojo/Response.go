package pojo

type Response struct {
	Data    interface{}
	RetCode int
	Message string
	Page    *Page
}
