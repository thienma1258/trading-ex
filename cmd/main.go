package main

import "trading/manager"

func main()  {
	manager.InitCommands()
	manager.Execute()
}