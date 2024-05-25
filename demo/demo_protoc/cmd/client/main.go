package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/conf"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}
	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	res, error := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%v", res)

}
