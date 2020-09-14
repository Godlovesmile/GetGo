# GetGo
go go go

# Go 1.11 以及之前版本的环境配置
## 一. 环境配置(Mac)
1. 终端, cd ~
2. 查看是否存在.bash_profile文件, ls -all
3. 如没有创建.bash_profile文件(有, 跳过)
   + 创建.bash_profile文件, touch .bash_profile
   + 编辑, open -e .bash_profile
   + 自定义.bash_profile文件内容
   ```
   export GOPATH=/Users/(电脑用户名)/Go
   export GOBIN=$GOPATH/bin
   export PATH=$PATH:$GOBIN
   ```
   + 编译.bansh_profile文件, source .bash_peofile
4. 查看Go环境变量: go env

## 二. 代码运行
1. 在go的环境中创建保存项目的文件夹

    mkdir -p $GOPATH/src
2. 将已知的项目路径指向到go环境的项目路径

   ```
   eg:
   cp -a apiserver_demos/demo01/(已知项目路径) $GOPATH/src/apiserver(go环境项目路径)
   ```    
3. 首次编译需要下载 vendor 包

   ```
   cd $GOPATH/src
   git clone https://github.com/lexkong/vendor
   ```
4. 进入 apiserver 目录编译源代码

   ```
   cd $GOPATH/src/apiserver
   gofmt -w .   
   go tool vet .
   go build -v .
   ```


# Go 1.12 开始已不再推荐使用GOPATH来构建应用了

## go mod 使用
1. go mod init 'name' -> 生成 go.mod
   cat go.mod -> 查看 go.mod 内容
   

## 常用 go 第三方库记录
1. [cobra](https://github.com/spf13/cobra)是一个命令行程序库

2. [viper](https://github.com/spf13/viper)是一个配置解决方案，拥有丰富的特性

3. pflag 是命令行参数解析包 pflag, pflag 包的设计目的就是替代标准库中的 flag 包，因此它具有更强大的功能并且与标准的兼容性更好