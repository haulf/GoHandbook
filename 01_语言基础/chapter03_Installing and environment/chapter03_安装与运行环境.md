## 3 安装与运行环境

### 3.1 平台与架构

Go语言开发团队开发了适用于以下操作系统的编译器：

-         Linux

-         FreeBSD

-         Mac OS X（也称为Darwin）

目前有2个版本的编译器： Go原生编译器gc和非原生编译器gccgo，这两款编译器都是在类Unix系统下工作。其中，gc版本的编译器已经被移植到Windows平台上，并集成在主要发行版中，也可以通过安装MinGW从而在Windows平台下使用gcc编译器。这两个编译器都是以单通道的形式工作。

对于非常底层的纯Go语言代码或者包而言，在各个操作系统平台上的可移植性是非常强的，只需要将源码拷贝到相应平台上进行编译即可，或者可以使用交叉编译来构建目标平台的应用程序。但如果打算使用cgo或者类似文件监控系统的软件，就需要根据实际情况进行相应地修改了。

(1) Go原生编译器gc

主要基于Ken Thompson先前在Plan 9操作系统上使用的C工具链。这款编译器使用非分代、无压缩和并行的方式进行编译，它的编译速度要比gccgo更快，产生更好的本地代码，但编译后的程序不能够使用gcc进行链接。

 (2) gccgo编译器

一款相对于gc而言更加传统的编译器，使用GCC作为后端。GCC是一款非常流行的GNU编译器，它能够构建基于众多处理器架构的应用程序。编译速度相对gc较慢，但产生的本地代码运行要稍微快一点。它同时也提供一些与C语言之间的互操作性。

 (3) 文件扩展名与包（package）

Go语言源文件的扩展名很显然就是 .go。Go语言的标准库包文件在被安装后就是使用这种格式的文件。 

### 3.2 Go环境变量

Go开发环境依赖于一些操作系统环境变量，最好在安装Go之间就已经设置好它们。这里列举几个最为重要的环境变量：

- `$GOROOT`表示Go在的电脑上的安装位置，它的值一般都是`$HOME/go`，当然，也可以安装在别的地方。

- `$GOARCH`表示目标机器的处理器架构，它的值可以是386、amd64或arm。

- `$GOOS`表示目标机器的操作系统，它的值可以是darwin、freebsd、linux或windows。

- `$GOBIN`表示编译器和链接器的安装位置，默认是`$GOROOT/bin`。

目标机器是指打算运行Go应用程序的机器。

Go编译器支持交叉编译，也就是说可以在一台机器上构建运行在具有不同操作系统和处理器架构上运行的应用程序，也就是说编写源代码的机器可以和目标机器有完全不同的特性。

为了区分本地机器和目标机器，可以使用`$GOHOSTOS`和`$GOHOSTARCH`设置目标机器的参数，这两个变量只有在进行交叉编译的时候才会用到，如果不进行显示设置，它们的值会和本地机器（`$GOOS`和`$GOARCH`）一样。

- `$GOPATH`默认采用和`$GOROOT`一样的值，但从Go1.1版本开始，必须修改为其它路径。它可以包含多个包含Go语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：src、pkg和bin，这三个目录分别用于存放源码文件、包文件和可执行文件。

- `$GOARM`专门针对基于arm架构的处理器，它的值可以是5或6，默认为6。

- `$GOMAXPROCS`用于设置应用程序可使用的处理器个数与核数。

### 3.3 在Linux上安装Go

1. 配置Go环境变量

在Linux系统下一般通过文件`$HOME/.bashrc` 配置自定义环境变量，根据不同的发行版也可能是文件`$HOME/.profile`。

`export GOROOT=$HOME/go`

为了确保相关文件在文件系统的任何地方都能被调用，还需要添加以下内容：

`export PATH=$PATH:$GOROOT/bin`

在开发Go项目时，还需要一个环境变量来保存的工作目录。

`export GOPATH=$HOME/Applications/Go`

`$GOPATH`可以包含多个工作目录，取决于个人情况。如果设置了多个工作目录，那么当在之后使用 go get（远程包安装命令）时远程包将会被安装在第一个目录下。

在完成这些设置后，需要在终端输入指令source .bashrc 以使这些环境变量生效。然后重启终端，输入 go env 和 env 来检查环境变量是否设置正确。

2. 安装C工具

Go 的工具链是用 C 语言编写的，因此在安装 Go 之前需要先安装相关的 C 工具。如果使用的是 Ubuntu 的话，可以在终端输入以下指令。

​       sudo apt-get install bison ed gawk gcc libc6-dev make

 可以在其它发行版上使用 RPM 之类的工具。

3. 获取Go源代码

从官方页面或国内镜像下载Go的源码包到的计算机上，然后将解压后的目录go通过命令移动到 $GOROOT 所指向的位置。

4. 测试安装

写一个输出hello world的程序测试。

5. 验证安装版本

可以通过在终端输入指令 go version 来打印 Go 的版本信息。

6. 更新版本

可以在发布历史页面查看到最新的稳定版。Go的源代码有以下三个分支：

-         Go release：最新稳定版，实际开发最佳选择

-         Go weekly：包含最近更新的版本，一般每周更新一次

-         Go tip：永远保持最新的版本，相当于内测版



### 3.4 在Mac OS X上安装Go

如果想要在的Mac系统上安装Go，则必须使用Intel 64位处理器，Go不支持PowerPC处理器。在Mac系统下使用到的C工具链是Xcode的一部分，因此需要通过安装Xcode来完成这些工具的安装。并不需要安装完整的Xcode，而只需要安装它的命令行工具部分。在安装Go程序时会弹出提示安装Xcode的命令行工具，按照提示一步步安装。

安装完成后，在终端里输入go version，若能够查看到版本信息，则说明安装成功了。

接着需要配置环境变量。在自己的主目录下新建.bash_profile文件。在.bash_profile文件中输入：

```
#aihaofeng begin, add
export GOPATH=/Users/aihaofeng/Documents/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
#aihaofeng end
```

接着运行source命令让设置生效：

```go
source .bash_profile
```

这里设置的GOPATH为/Users/aihaofeng/Documents/go，也就是把工作目录设置在Documents下的go目录，所以需要在自己的Documents目录下创建名称为go目录。

如果Mac电脑上安装的是ohmyzsh，则并不使用bash，因此上面的配置不会生效。需要修改的是.zshrc文件。

```shell
➜  /Users/aihaofeng subl .zshrc

# aihaofeng begin, 20171117
# Go environment config.
export GOPATH="/Users/aihaofeng/go"
export GOBIN="/Users/aihaofeng/go/bin"
export PATH="$PATH:$GOBIN"
# aihaofeng end
```

GOROOT: 指的是Go的安装目录。

GOPATH：指工作目录。

GOBIN：go install时位置。

### 3.5 在Windows上安装Go

可以在下载页面页面下载到 Windows 系统下的一键安装包。在完成安装包的安装之后，只需要配置$GOPATH 这一个环境变量就可以开始使用 Go 语言进行开发了，其它的环境变量安装包均会进行自动设置。在默认情况下，Go将会被安装在目录c:\go 下，但如果在安装过程中修改安装目录，则可能需要手动修改所有的环境变量的值。

### 3.6 安装目录清单

Go 安装目录（$GOROOT）的文件夹结构应该如下所示：

​        README.md, AUTHORS, CONTRIBUTORS, LICENSE

-         /bin：包含可执行文件，如：编译器，Go工具。

-         /doc：包含示例程序，代码工具，本地文档等。

-         /lib：包含文档模版。

-         /misc：包含与支持Go编辑器有关的配置文件以及cgo的示例。

-         /os_arch：包含标准库的包的对象文件（.a）。

-         /src：包含源代码构建脚本和标准库的包的完整源代码。

-         /src/cmd：包含Go和C的编译器和命令行脚本。



### 3.7 Go运行时

尽管Go编译器产生的是本地可执行代码，这些代码仍旧运行在Go的runtime（这部分的代码可以在runtime包中找到）当中。这个runtime类似Java和.NET语言所用到的虚拟机，它负责管理包括内存分配、垃圾回收、栈处理、goroutine、channel、切片（slice）、map和反射（reflection）等等。

runtime主要由C语言编写（Go1.5开始自举），并且是每个Go包的最顶级包。可以在目录$GOROOT/src/runtime中找到相关内容。

垃圾回收器Go拥有简单却高效的标记-清除回收器。它的主要思想来源于IBM的可复用垃圾回收器，旨在打造一个高效、低延迟的并发回收器。目前gccgo还没有回收器，同时适用gc和gccgo的新回收器正在研发中。使用一门具有垃圾回收功能的编程语言不代表可以避免内存分配所带来的问题，分配和回收内容都是消耗CPU资源的一种行为。

Go的可执行文件都比相对应的源代码文件要大很多，这恰恰说明了Go的 runtime 嵌入到了每一个可执行文件当中。当然，在部署到数量巨大的集群时，较大的文件体积也是比较头疼的问题。但总体来说，Go的部署工作还是要比Java和Python轻松得多。因为Go不需要依赖任何其它文件，它只需要一个单独的静态文件，这样也不会像使用其它语言一样在各种不同版本的依赖文件之间混淆。
