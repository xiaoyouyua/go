package hanlder

//名称冲突问题
const HelloServiceName = "handler/HelloService"

type NewHelloService struct {
}

func (s *NewHelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello," + request
	return nil
}
