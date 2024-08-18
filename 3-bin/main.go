package main

import "cloud-json/file"

func main() {
	file.ReadFile("~/.config/nvim/init.lua")
}
