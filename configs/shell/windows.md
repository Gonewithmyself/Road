### 修改配色

 [参见](https://github.com/neilpa/cmd-colors-solarized)

### 安装 Powerline Fonts

[下载](<https://github.com/powerline/fonts>)后直接安装 *DejaVuSansMono* 这个字体

### 修改默认字体及代码页

打开注册表，定位到`HKEY_CURRENT_USER\Console\%SystemRoot%_System32_wsl.exe`。增加一个`DWORD`项，命名为`CodePage`，值设为十进制`65001`。再增加一个`字符串`项，命名为`FaceName`，值为`DejaVu Sans Mono for Powerline`。如果Linux子系统是从Windows应用商店安装的，`Console`项下面应该还会有个`C:_Program Files_WindowsApps_`开头的项，也需要进行相同的修改。



