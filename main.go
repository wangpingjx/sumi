package main

import (
     "goblin"
     _ "sumi/routers"
)

func main() {
    goblin.Run(":9090")
}
