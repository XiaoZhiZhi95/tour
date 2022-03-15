package main

import (
	"flag"
	"log"
)

func main() {
	/*
	var name string
	var fz string
	//将命令行参数与变量进行绑定
	flag.StringVar(&fz, "fz", "lili", "best in the world")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	log.Printf("name: %s, fz = %s", name, fz)
	 */

	flag.Parse()

	arg := flag.Args()
	if len(arg) <= 0 {
		log.Println("there is no arg, exit success")
		return
	}

	var name string
	switch arg[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "default go", "help info")
		_ = goCmd.Parse(arg[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(arg[1:])
	}

	log.Println("name = ", name)
}
