package main

import (
	"github.com/iesreza/dockertest/components/helloworld"
	"github.com/iesreza/foundation/system"
)

func main()  {
	system.PreBoot()
	helloworld.Register()
	system.Boot()
	system.ListenCLI()
}
