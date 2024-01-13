#!/bin/bash
if [ "$EUID" -ne 0 ]; then
    echo "Run script as root, sudo -i."
    exit 1
fi
sudo apt update -y
sudo apt-get install pkg-config libssl-dev -y
wget -O /etc/logo.sh https://raw.githubusercontent.com/Azumi67/UDP2RAW_FEC/main/logo.sh
chmod +x /etc/logo.sh

architecture=$(uname -m)
if [ "$architecture" = "x86_64" ]; then
    if [ ! -f "go1.21.5.linux-amd64.tar.gz" ]; then
        wget https://github.com/Azumi67/UDP2RAW_FEC/releases/download/go/go1.21.5.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
    fi
elif [ "$architecture" = "aarch64" ]; then
    if [ ! -f "go1.21.5.linux-arm64.tar.gz" ]; then
        wget https://github.com/Azumi67/UDP2RAW_FEC/releases/download/go/go1.21.5.linux-arm64.tar.gz
        sudo tar -C /usr/local -xzf go1.21.5.linux-arm64.tar.gz
    fi
else
    echo "Unsupported arch: $architecture"
    exit 1
fi

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile

source ~/.bash_profile
go mod init mymodule
go install github.com/AlecAivazis/survey/v2
go mod tidy
go get github.com/AlecAivazis/survey/v2
go get github.com/fatih/color

if [ -f "http2.go" ]; then
    rm http2.go
    echo "deleted previous version!"
fi

wget https://raw.githubusercontent.com/Azumi67/Http2_reverseTunnel/main/http2.go

go run http2.go
