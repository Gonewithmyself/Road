## 安装zsh

`apt install zsh -y`

`sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"`

`chsh /bin/zsh`

~/.bashrc中加入`bash -c zsh` 设置默认shell

### 安装插件

`git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting`

`git clone https://github.com/zsh-users/zsh-autosuggestions${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions`