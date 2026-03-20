package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// ================== LS COMMAND ==================
	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("Output:", string(output))

	// ================== PIPES ==================
	// pipeReader, pipeWriter := io.Pipe()
	// cmd := exec.Command("grep", "foo")
	// cmd.Stdin = pipeReader

	// go func() {
	// 	defer pipeWriter.Close()
	// 	pipeWriter.Write([]byte("food is good\nbar\nbaz\n"))
	// }()
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("err starting command", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	// cmd := exec.Command("printenv", "SHELL")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	// ================== LONG PROCESS ==================
	// cmd := exec.Command("sleep", "60")
	// // Start command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("err starting command", err)
	// 	return
	// }

	// ------------------ KILL ------------------
	// // Kill the process
	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("err starting command", err)
	// 	return
	// }
	// fmt.Println("Process is killed")

	// ------------------ WAITING ------------------
	// Waiting
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("err waiting", err)
	// 	return
	// }
	// fmt.Println("Process is completed")

	// ================== WITH INPUT ==================
	// cmd := exec.Command("grep", "foo")
	// // Set input
	// cmd.Stdin = strings.NewReader("food is good\nbar\nbaz\n")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	return
	// }
	// fmt.Println("Output:", string(output))

	// ================== BASIC ==================
	// cmd := exec.Command("echo", "Hello World!")

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Err:", err)
	// 	return
	// }

	// fmt.Println("Output:", string(output))
}
