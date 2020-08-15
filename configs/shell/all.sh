#!/bin/bash

# apt source list
mv /etc/apt/source.list /etc/apt/source.list.bak
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted" >> /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic universe" >> /etc/apt/sources.list
echo "deb http://mirrors.aliyun.com/ubuntu/ bionic multiverse" >> /etc/apt/sources.list
apt update


# ssh
apt install -y openssh-server
mkdir -p /root/.ssh/
echo 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDGQ04uNjHJ5ZWRGDXj1Aq8sMT7INyrqV5fC2hQwZMHAGMzBFXhOuwIQFIyvcxgakv/8R/th9z3yD8M/h176NZYw3NHQ1pg35Xg4ZYgSxlN+AO1kFavW9Y2pU7l5DeMoFS6Rat+E11bOeiHeOY9hOStE/P01Vy4hJKq+Vom6shfZn3H6EyPf62z1jZdv5kA9tG5h1w9jCPkRQPLpfLVTnTLEgrs7UdL7wfuH9k2Z46ka/di2i0p1u1dx465iLd2teFEDy1/IsKd64i+U3F6FlpoNXVNF+BfJdp11b9wFYKHbfvu2hvMpNGYBuQ9ZmBEsMMGe1sYBO4PJ0z1BiROWmuZ guoqiang' >> /root/.ssh/authorized_keys
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
