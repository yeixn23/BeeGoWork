package models

/*
人结构体的定义，包含三个段：
*/
//type Person struct {
//	Name string
//	Age int
//	Sex string
//}
/*
定义一个register,包含4个段：
*/
type Register struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Address  string `json:"address"`
	Nick     string `json:"nick"`
}
