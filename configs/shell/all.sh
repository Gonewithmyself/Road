#!/bin/bash

# apt source list
mv /etc/apt/source.list /etc/apt/source.list.bak
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted" >> /etc/apt/source.list
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic universe" >> /etc/apt/source.list
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic multiverse" >> /etc/apt/source.list
apt update


# ssh
apt install -y openssh-server
mkdir -p /root/.ssh/
echo '' >> /root/.ssh/authorized_keys
service ssh restart


# install zsh
apt install -y zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
# sh install.sh
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions


# apt install -y 
#softs=["make" "build-essential" "htop" "dos2unix" "net-tools"]
apt install -y "make" "build-essential" "htop" "dos2unix" "net-tools" 


# docker 
apt install -y docker.io docker-compose
echo '{"registry-mirrors":["https://reg-mirror.qiniu.com/"]}' >> /etc/docker/daemon.json