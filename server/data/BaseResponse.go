package data

type T interface {

}

type BaseResponse struct {
	Code int
	Msg  string
	Data T
}