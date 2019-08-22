package main

import "os/exec"

//生成子文件内容
func main() {
	cmd := exec.Command("go", "build", "app/main.go")
	_ = cmd.Start()
}
