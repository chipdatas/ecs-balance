#!/bin/bash

# 指定下载 URL
URL="http://example.com/path/to/abcc"

# 检查 wget 是否存在
if ! command -v wget &> /dev/null; then
    echo "wget 未安装，正在安装 wget..."

    # 检查操作系统类型并安装 wget
    if [ -f /etc/debian_version ]; then
        # Debian/Ubuntu
        if [ "$(id -u)" -eq 0 ]; then
            apt-get update && apt-get install -y wget
        else
            sudo apt-get update && sudo apt-get install -y wget
        fi
    elif [ -f /etc/redhat-release ]; then
        # RHEL/CentOS
        if [ "$(id -u)" -eq 0 ]; then
            yum install -y wget
        else
            sudo yum install -y wget
        fi
    else
        echo "不支持的操作系统，请手动安装 wget"
        exit 1
    fi
else
    echo "wget 已安装"
fi

ALI_DIR="$HOME/.ali"
mkdir -p "$ALI_DIR"
cd "$ALI_DIR"
wget -O abcc "$URL"
chmod +x abcc
ali auto