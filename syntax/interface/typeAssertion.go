package main

/**
	类型断言

 */
func main() {
	var s interface{}
	s = "string"
	ret,ok := s.(bool)  //ret 是接口绑定值实例值的副本（实例可能是指针值）
	if !ok {
		println("接口类型判断错误。。。", ok)
	} else {
		println("接口判读正确")
	}
	println(ret)


	// switch 方式类型查询
	switch v:= s.(type) {
	case bool:
		println("type is bool",v)
	case string:
		println("type is string",v)
	default:
		println("i don't know",v)
	}
}
