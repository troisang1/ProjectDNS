package main

import (
	"Tiamat/prisma"
	"Tiamat/process"
)

func main() {
	prisma.InitPrisma()
	process.HandleRequest()
}
