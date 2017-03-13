package data

//Go 默认是不支持泛型的，定义一个接口结构体，即可当泛型使用
type T interface {

}

type BaseResponse struct {
	Code int
	Msg  string
	Data T
}