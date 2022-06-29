package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"proto_demo/proto/userService"
)

func main() {
	u := &userService.Userinfo{
		Username: "test",
		Age:      10,
		Hobby:    []string{"吃饭", "睡觉", "打豆豆"},
	}
	fmt.Println(u)
	fmt.Println(u.GetType())
	fmt.Println(u.GetUsername())
	fmt.Println(u.GetHobby())
	//protobuf的序列化
	data, _ := proto.Marshal(u)
	fmt.Println(data)
	//protobuf的反序列化
	var user userService.Userinfo
	proto.Unmarshal(data, &user)
	fmt.Printf("%#v", user)
}
