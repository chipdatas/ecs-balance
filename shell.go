package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func SetShell() {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前路径:", err)
		return
	}

	command := filepath.Join(currentPath, "ali")

	shell := os.Getenv("SHELL")
	var shellConfigFile string

	switch {
	case shell == "/bin/bash":
		if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".bashrc")); err == nil {
			shellConfigFile = filepath.Join(os.Getenv("HOME"), ".bashrc")
		} else {
			shellConfigFile = filepath.Join(os.Getenv("HOME"), ".bash_profile")
		}
	case shell == "/bin/zsh":
		shellConfigFile = filepath.Join(os.Getenv("HOME"), ".zshrc")
	case shell == "/usr/bin/fish":
		shellConfigFile = filepath.Join(os.Getenv("HOME"), ".config/fish/config.fish")
	default:
		fmt.Println("不支持的 shell 类型:", shell)
		fmt.Println("请手动添加以下命令到shell的配置文件:")
		fmt.Println(command)
		return
	}

	file, err := os.OpenFile(shellConfigFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("无法打开 shell 配置文件:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("\n# 自动添加的命令\n%s\n", command)); err != nil {
		fmt.Println("无法写入 shell 配置文件:", err)
		fmt.Println("请手动添加以下命令到", shellConfigFile)
		fmt.Println(command)
		return
	}

	fmt.Printf("已将命令添加到 %s\n", shellConfigFile)

	if shell == "/bin/zsh" {
		cmd := exec.Command("zsh", "-c", "source ~/.zshrc")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("无法执行 source 命令,配置将在下次访问 SSH时生效", err)
			return
		}
		fmt.Println("已执行 source ~/.zshrc")
	} else if shell == "/bin/bash" {
		cmd := exec.Command("bash", "-c", "source ~/.bashrc")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("无法执行 source 命令,配置将在下次访问 SSH时生效", err)
			return
		}
		fmt.Println("已执行 source ~/.bashrc")
	} else if shell == "/usr/bin/fish" {
		cmd := exec.Command("fish", "-c", "source ~/.config/fish/config.fish")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("无法执行 source 命令,配置将在下次访问 SSH时生效", err)
			return
		}
		fmt.Println("已执行 source ~/.config/fish/config.fish")
	} else {
		fmt.Println("请手动执行 source 命令以使更改立即生效，否则配置将在下次访问 SSH时生效")
	}
	fmt.Println("已执行 source ~/.zshrc")
	fmt.Println("所有任务执行完成，退出...")
}
