# Golang Programming Style

![96](http://upload.jianshu.io/users/upload_avatars/2463211/80337da84cfb.jpg?imageMogr2/auto-orient/strip|imageView2/1/w/96/h/96)

 

[_张晓龙_](http://www.jianshu.com/u/1381dc29fed9)

 

**已关注

2016.12.19 23:56* 字数 8039 阅读 531评论 3喜欢 11

## 前言

本规范是针对 *Golang* 语言的编码规范，目的是为了统一项目的编码风格，提高软件源程序的可读性、可靠性和可重用性，提高软件源程序的质量和可维护性，减少软件维护成本。
本规范适用于部门所有产品的软件源程序，同时考虑到不同产品和项目的实际开发特性，本规范分成规则性和建议性两种：对于规则性规范，要求所有软件开发人员严格执行；对于建议性规范，各项目编程人员可以根据实际情况选择执行。本规范的示例都以*Golang* 语言描述。
本规范的内容包括：开发环境、包设计、布局、注释、命名、基本元素设计、函数设计、错误和异常设计、整洁测试等。
本规范自生效日起，对以后新编写的和修改的代码有约束力。
对本规范中所用的术语解释如下：
**原则**：编程时应该坚持的指导思想。
**规则**：编程时必须遵守的约定。
**建议**：编程时必须加以考虑的约定。
**说明**：对此规则或建议的必要的解释。
**正例**：对此规则或建议给出的正确例子。
**反例**：对此规则或建议给出的反面例子。

golang.jpg

## 开发环境

**【规则1-1】为了防止代码出现可移植性问题和兼容性问题，团队使用的操作系统、编译器类型、版本保持一致性。**

**【规则1-2】团队统一使用相同的IDE，并使用统一的代码模板，保持代码风格的一致性。**

**说明：**系统中所有的代码看起来就好像是由单独一个值得胜任的人编写的。

**【规则1-3】团队统一配置 IDE 的 TAB 为4个空格。**

## 包设计

**【原则2-1】包设计要满足单一职责原则。**

**说明：** 这是SRP(Single Reponsibility Priciple) 在包(package)设计时的一个具体运用，我们要将包设计的非常内聚，包间的 API 比较少（类似于class中的public方法）。

**【原则2-2】包内标识符遵守最小可见性原则。**
**说明：** 如果一个标识符（interface名、类型名、变量名或函数名）在语义上仅在包内可见，则它的命名不要用大写开头。

**【规则2-1】测试文件放在实现文件的同级目录下，便于Golang工具的使用。**

**说明：** 虽然测试文件和实现文件的代码在同一个包内，但是测试用例的设计尽量站在包用户的角度去考虑，一般只测试包外可见的函数。

**【规则2-2】包间禁止共享全局变量。**

**说明：常量除外

**【规则2-3】 不允许一个目录下有多个包。**

**【规则2-4】 import包时不允许使用点(.)操作。**

**正例：**

```
import (
    "fmt"
    "os"
)

func main() {
    for _, value := range os.Args {
        fmt.Println(value)
    }
}

```

**反例：**

```
import (
    . "fmt"
    "os"
)

func main() {
    for _, value := range os.Args {
        Println(value)
    }
}

```

**说明：测试文件中的测试框架除外

**【规则2-5】 import包时不允许使用别名。**

**正例：**

```
import (
    "fmt"
    "os"
)

func main() {
    for _, value := range os.Args {
        fmt.Println(value)
    }
}

```

**反例：**

```
import (
    f "fmt"
    "os"
)

func main() {
    for _, value := range os.Args {
        f.Println(value)
    }
}

```

**【规则2-6】 import包时不允许使用下划线(_)操作。**

**说明：**下划线（_）操作的含义是：导入该包，但不导入整个包，而是执行该包中的init函数，因此无法通过包名来调用包中的其他函数。使用下划线（_）操作往往是为了注册包里的引擎，让外部可以方便地使用。

**正例：**

```
import (
    "fmt"
    "os"
)

func main() {
    for _, value := range os.Args {
        fmt.Println(value)
    }
}

```

**反例：**

```
import (
    "fmt"
    "os"
    _ "time"
)

func main() {
    for _, value := range os.Args {
        fmt.Println(value)
    }
}

```

**注：反例中的time包仅为冗余的包，并不是为了注册引擎，同时我们的产品代码中，也没有隐式注册引擎的需求，所以我们在import包时统一不允许使用下划线(_)操作。**

## 布局

**【规则3-1】import 导入包时统一使用小括号，包名要另起一行 **

**说明：import "C" 除外

**正例：**

```
import (
    "fmt"
    "reflect"
)

```

**反例：**

```
import "fmt"
import "reflect"

```

**【规则3-2】import 包时，路径分隔符一律使用Unix 风格，拒绝使用Windows 风格；即采用/ 而不是使用\ 分割路径。
**

**正例：**

```
import (
    "knitter-agent/domain/object/port-obj"
)

```

**反例：**

```
import (
    "knitter-agent\domain\object\port-obj"
)

```

**【规则3-3】import 包时以 $GOPATH 为基准使用绝对路径，不要从当前位置开始使用相对路径 **

**正例：**

```
import (
    "knitter-agent/domain/object/port-obj"
)

```

**反例：**

```
import (
    "../../object/port-obj"
)

```

**【规则3-4】包含空格在内，代码的行宽不应超过120列。 **
**说明：** 长行要在低优先级操作符处拆分成新行，拆分出的新行要进行适当的缩进，使排版整齐。

**【规则3-5】程序实体之间有且仅有一行空行区分。**
**说明：** 函数之间的空行，能够帮组我们快速定位函数的始末的准确位置；甚至在函数内部，将逻辑相关的代码放在一起也同样具有意义，它能够帮组我们更好地理解代码块的语义。超过一行的空行完全没有必要，部分粗心的程序员在处理这些细节时总存在着或多或少的问题，团队应该杜绝这样的情况发生。

**【规则3-6】每个文件末尾都应该有且仅有一行空行。**

*【规则3-7】一元操作符如“!”、“~”、“++”、“--”、 “”、 “&”（地址运算符）等前后不加空格； “［］”、“.”、“->”这类操作符前后不加空格。**

**【规则3-8】函数名之后不要留空格**
**说明：** 函数名后紧跟左括号‘(’，以与关键字区别。
**正例：**

```
func getOpenstackConf() string {
    return OpenstackConfStr
}

```

**反例：**

```
func getOpenstackConf () string {
    return OpenstackConfStr
}

```

**【规则3-9】在进行“==”或“!="比较时，将常量或常数放在“==”或“!="号的右边。**
**正例：**

```
if err != nil {
    glog.Errorf("bind-GetBusInfo:getEthtoolOutputFunc failed, err: %v", err)
    return "", err
}

if link.Type() == "device" {
    return link
}

```

**反例：**

```
if nil != err {
    glog.Errorf("bind-GetBusInfo:getEthtoolOutputFunc failed, err: %v", err)
    return "", err
}

if "device" == link.Type() {
    return link
}

```

**【规则3-10】 数组的初始化按照矩阵结构分行书写。**

**正例：**

```
numbers := [4][3]int {
    1, 1, 1,
    2, 4, 8,
    3, 9, 27,
    4, 16, 64,
}


```

**【建议3-1】 每写完一段代码，就使用gofmt工具格式化一下。**

## 注释

> 注释有助于理解代码，有效的注释是指在代码的功能、意图层次上进行注释，提供有用、额外的信息，而不是代码表面意义的简单重复。
> 注释的恰当用法是弥补我们在用代码表达意图时遭遇的失败。每次写注释，你都应该做个鬼脸，感受自己在表达能力上的失败。
> 写注释时，首先想到的应该是通过重构来提高表达力，不要太早放弃。

**【规则4-1】注释与所描述内容进行同样的缩进。**
**说明：** 可使程序排版整齐，并方便注释的阅读与理解。

**【规则4-2】避免垃圾注释。**
**说明：** 对于代码本身能够表达的意思，不必增加额外的注释。

*【规则4-3】注释符 “//” 或 "/\*” (“/”) 与注释内容之间用一个空格分隔。**

**【建议4-1】并非所有的函数都要配有函数头，短函数需要一个好名字而非太多描述。**

**【建议4-2】提倡代码自注释。**
**说明：** 能用函数或变量时就不要用注释，如果可以的话，应该创建一个描述与注释所言同一事物的函数或变量用于消除注释。

**【建议4-3】行注释和块注释都可行时，优先使用行注释。**

**【建议4-4】保证代码和注释的一致性。**

**说明：** 修改代码同时修改相应的注释，不再有用的注释要删除。

**【建议4-5】注释应与其描述的代码相近，对代码的注释应放在其上方或右方（对单条语句的注释）相邻位置，不可放在下面，如放于上方则需与其上面的代码用空行隔开，而且注释内容与与被注释的代码相同缩进。**

## 命名

**【规则5-1】命名要名副其实——要像给自己的baby 起名字一样谨慎来对待程序命名。**

**说明：** 变量、函数的命名告诉我们，它为什么会存在，它做什么事，应该怎么用。如果名称需要注释来补充，那就不算是名副其实。要像给自己的baby 起名字一样谨慎来对待程序命名。

**【规则5-2】目录名一律使用小写和中划线风格的名称。**

**正例：**
knitter-agent

**反例：**
knitteragent
knitter_agent
KnitterAgent
knitterAgent

**【规则5-3】 包名一律使用小写风格，通常为过滤掉中划线的目录名。**

**正例1：**
目录名：context
对应的包名：context

**正例2：**
目录名：port-obj
对应的包名：portobj

**【规则5-4】 开发文件命名一律使用小写和下划线风格的名称。**

**正例：**
knitter_virtual_machine.go

**反例：**
knitter-virtual-machine.go
KnitterVirtualMachine.go
knittervirtualmachine.go

**【规则5-5】标识符要采用英文单词或其组合，便于记忆和阅读，切忌使用汉语拼音来命名。**

**说明：** 标识符应当直观且可以拼读，可望文知义，避免使人产生误解。程序中的英文单词一
般不要太复杂，用词应当准确。

**【规则5-6】如果函数返回值的类型或变量的类型为bool，则名字前面加上is, has, may, can, should, need等词修饰会增强语意。**

**正例：**

```
var isDpdk bool

```

**反例：**

```
var dpdk bool

```

**【规则5-7】接口名、类型名、变量名和函数名统一使用驼峰命名法，首字母是否大写由包外可见性决定。**
**说明：** 应遵循最小可见性原则

**【规则5-8】避免在名称中携带类型信息。**

**正例：**

```
var num int
var ports []Port
var tenantNetworks map[string]Value

```

**反例：**

```
var iNum int
var portSlice []Port
var tenantNetworkMap map[string]Value

```

**【规则5-9】避免在名称中携带作用域的信息。**

**正例：**

```
var num int

```

**反例：**

```
var gNum int

```

**【规则5-10】 变量名的主体应当使用“名词”或者“形容词＋名词”。**

**【规则5-11】 函数名应当使用“动词”或者“动词＋名词”（动宾词组）。**

**【规则5-12】 系统中每个实体概念对应一个词。**

**说明：** 给每个抽象概念选一个词，并且在同一个系统中统一，以便符合SRP 原则。如在同一系统的代码中既有controller，还有manager 和driver，会令使用者困惑，应统一。

**【规则5-13】 不使用双关语命名变量。**

**说明：** 变量命名时应避免将同一单词用于不同目的，同一术语用于不同概念，应遵从“一词一义”规则。比如add在表达计算两个值的和的语义时，就不能再表达往一个数组切片插入一个元素的语义。

**【规则5-14】 常量名使用驼峰命名法，首字母是否大写由包外可见性决定。**
**说明：** 应遵循最小可见性原则

**【规则5-15】 团队使用统一的缩略语，并和业界常用的缩略语保持一致。**

**说明：** 较短的单词可通过去掉“元音”形成缩写，较长的单词可取单词的头几个字母形成缩写，一些单词有大家公认的缩写，常用单词的缩写必须统一。协议中的单词的缩写与协议保持一致。对于某个系统使用的专用缩写应该在某处注释中做统一说明。

**正例：** 如下单词的缩写能够被大家认可
temp 可缩写为：tmp
flag 可缩写为：flg
statistic 可缩写为：stat
increment 可缩写为：inc
message可缩写为：msg

规范的常用缩写如下：

**【规则5-16】 用正确的反义词组命名具有互斥意义的变量或相反动作的函数等。**

**正例：**

| 词组                    | 词组                   |
| --------------------- | -------------------- |
| add / remove          | insert / delete      |
| create / destroy      | begin / end          |
| first / last          | lock / unlock        |
| increment / decrement | push / pull          |
| open / close          | min / max            |
| old / new             | start / stop         |
| next / previous       | source / destination |
| show / hide           | send / receive       |
| attach / detach       | left / right         |
| up / down             | north / south        |

## 基本元素设计

### 变量与常量

**【规则6-1-1】 一个变量有且只有一个功能，并与其名称相一致，不能把一个变量用作多种用途。**

**说明：** 一个变量只用来表示一个特定功能，不能把一个变量用作多种用途，即同一变量取值不同时，其代表的意义也不同。除循环变量和收集计算结果的变量，在一个函数中，一个变量被赋值不应该超过一次。

**【规则6-1-2】 代码中不允许出现魔法数。**

**说明：** 魔法数，即拥有特殊意义，却又不能明确表现出这种意义的数字。用const来定义常数，并根据其意义为它命名，既提高了代码的可读性，又便于使用IDE 等工具进行查找修改。

**【规则6-1-3】 如果 struct 中的数据变量需要进行 json 序列化，则需要以大写字母开头，同时需要 json 重命名。**

**说明：** 结构体中的变量以大写字母开头，可以保证 json.Marshal 的时候数据持久化正确。如果结构体中的变量以小写字母开头，则使得 json.Marshal 的时候忽略该字段，使得该字段的值丢失，从而 json.Unmarshal 的时候将该变量的值置为默认值。由于结构体中的变量以大写字母开头， json 串中的字段 key 的字符串形式变成了以大写字母开始，这对于追求以 json 串全小写为美的我们来说，需要进行 json 重命名。

**正例：**

```
type Position struct {
    X int `json:"x"`
    Y int `json:"y"`
    Z int `json:"z"`
}

type Student struct {
    Name string `json:"name"`
    Sex string `json:"sex"`
    Age int `json:"age"`
    Posi Position `json:"position"`
}

```

**反例：**

```
type Position struct {
    X int
    Y int
    Z int
}

type Student struct {
    Name string
    Sex string
    Age int
    Posi Position
}

```

**【建议6-1-1】 变量应尽可能的满足短跨度和短存活时间。**

**说明：** 那些介于同一个变量多个引用点之间的代码可称为攻击窗口，我们用跨度来衡量一个变量的不同引用点之间的靠近程度，而变量的存活时间是一个变量存在期间所跨越的语句总数。跨度越短，则表明一个变量的不同引用点越靠近；存活时间越短，则表明一个变量经历的语句数越少。

我们追求的目标是短跨度和短存活时间，因为
（1）可以提高程序的可读性；
（2）可以减小变量的攻击窗口；
（3）可以减少变量的初始化错误；
（4）可以减少全局变量的使用；
（5）可以方便修改Bug；
（6）可以方便重构代码。

### 表达式和语句

**【规则6-2-1】 对于布尔类型的变量，应直接进行真假判断**

**正例：**

```
/* 设flag 是布尔类型的变量 */
if flag  /* 表示flag为真 */
if !flag  /* 表示flag为假 */

```

**反例：**

```
/* 设flag 是布尔类型的变量 */
if flag == true
if flag == 1
if flag == false
if flag == 0

```

**【规则6-2-2】 在条件判断语句中，当整型变量与0 比较时，不可模仿布尔变量的风格，应当将整型变量用“==”或“!=”直接与0比较。**

**正例：**

```
/* 设value是整型的变量 */
if value == 0
if value != 0


```

**反例：**

```
/* 设value是整型的变量 */
if value /* 会让人误解 value是布尔类型的变量 */
if !value


```

**【规则6-2-3】 逻辑表达式已经具有 true 或 false 语义，无需画蛇添足。**

**正例：**

```
return i == 3


```

**反例：**

```
if i == 3 {
    return true
} else {
    return false
}


```

**【建议6-2-1】 循环嵌套次数不大于3。**

**【建议6-2-2】 if 语句的嵌套层数不要大于3。**

**说明：** 适当调整和优化判断逻辑，能够有效地控制if语句的嵌套层次，这对于代码的走查、测试、变更维护都有很大的帮助。如果能减少大语句块的嵌套深度，对于减轻代码阅读时的理解负担很有好处。
条件式通常有两种呈现形式：第一种形式是所有分支都属于正常行为；第二种形式则是条件式提供的答案只有一种是正常行为，其他都是不常见的情况。
这两类条件式有不同的用途，这一点应该通过代码表现出来。如果两条分支都是正常行为，就应该使用形如if-else的条件式；如果某个条件极其罕见，就应该单独检查该条件，并在该条件为真时立刻从函数中返回，这样的单独检查常常被称为卫语句。
使用卫语句，能够有效的减少if语句嵌套层数。

**【建议6-2-3】 使用for循环时，优先使用range 关键字而不是显式下标递增控制。**

**正例：**

```
for i, v := range array {
    fmt.Printf("element %v of array is %v\n", i, v)
}

```

**反例：**

```
for i := 0; i < len(array); i++ {
    fmt.Printf("element %v of array is %v\n", i, array[i])
}

```

**【建议6-2-4】 对于 range 的返回值，如果只需要第二项，则把第一项置为下划线。**
**正例：**

```
sum := 0
for _, value := range array {
    sum += value
}

```

## 函数设计

### 函数实现

**【规则7-1-1】 函数命名要短小精悍和名副其实，避免误导。一般以它" 做什么" 来命名，而不是以它" 怎么做" 来命名。**

**说明：** 函数命名名副其实就是指通过只读函数的名称就可以知道函数的功能，而不需要注释来补充。
给函数命名的方法：通过对要完成的功能进行分解和抽象，将功能分解成一个个单一的短小的功能实现体，对实现体的功能采用一个恰当的描述性名称命名，形成函数名称。

**【规则7-1-2】 函数要短小，还要更短小。尽量控制在20行代码之内，包括空行和{}。**

**说明：** 有几个原因造成我喜欢短而命名良好的函数。首先，如果每个函数的粒度都很小，那么函数被复用的机会就更大；其次，如果函数都是细粒度，那么函数在修改时也会更容易些；再次，高层函数调用命名良好的短小函数，使高层函数读起来就像一系列解释。
一个函数多长才算合适？长度不是问题，关键在于函数名称和函数本体之间的语义距离。建议函数体的规模不能太大，20 行封顶最佳。

**【规则7-1-3】 函数应该做一件事，做好这件事，只做这一件事。**

**说明：** 判断一个函数是否只做了一件事，可以通过两种方法：
（1）函数只是做了该函数名下同一抽象层上的步骤，则函数只做了一件事；
（2）如果一个函数内部的实现还可以拆分出一个函数，则该函数违反只做一件事原则。

**【规则7-1-4】 函数的缩进层次不应该超过3层。**

**【规则7-1-5】 分隔指令与询问，不要设置多功能函数。**

**说明：** 函数要么做什么事，要么回答什么事，两者不可兼得。如某个函数既返回对象状态值，又修改对象状态值，则需要建立两个不同的函数，其中一个负责查询对象状态，另一个负责修改对象状态。

**【建议7-1-1】 为简单功能编写函数。**

**说明：** 虽然为仅用一两行就可完成的功能去编函数好象没有必要，但使用函数可使功能明确化，增加程序可读性，亦可方便维护、测试。

### 参数

**【规则7-2-1】 禁止定义多于3个参数的函数。**

**说明：** 函数参数设置最理想的参数个数是零，其次是一，再次是二，最后是三。参数不易
对付，它们有太多的概念性。另外从测试的角度看，参数更叫人为难。

**【规则7-2-2】 函数参数不能含有标识参数。**

**说明：** 标识参数丑陋不堪，函数往往根据它的多个取值而做多件事情，这与函数只做一件事原则违背。如果参数只是用于赋值，那么就不是标识参数，所以是否标识参数不是今通过形参来界定，而是看函数的实现是否因为函数的入参而做了多件事情。

**【规则7-2-3】当struct变量作为参数时，应传送struct的指针而不传送struct，并且不得修改struct中的元素，用作输出时除外。**

**说明：** 一个函数被调用的时候，形参会被一个个压入被调函数的堆栈中，在函数调用结束以后再弹出。一个结构所包含的变量往往比较多，直接以一个结构为参数，压栈出栈的内容就会太多，不但占用堆栈空间，而且影响代码执行效率。
如果使用结构的指针作为参数，因为指针的长度是固定不变的，结构的大小就不会影响代码执行的效率，也不会过多地占用堆栈空间。
如果传递的参数类型是 map、slice 和 channel 等引用类型，则不用传递指针，修改引用类型变量的初始地址除外（比如 json.Unmarshal）。

**【规则7-2-4】在API函数中对输入参数的正确性和有效性进行检查，在内部能保证的条件下其他函数不用再进行重复检查。**

**说明：** 很多程序错误是由非法参数引起的，我们应该充分理解并正确处理来防止此类错误，特别是指针参数地址非法判断和数组下标参数的边界判断，但是我们没有必要在多个函数中重复检查。

**【规则7-2-5】防止将函数的参数作为工作变量。**

**说明：** 将函数的参数作为工作变量，有可能错误地改变入参的内容，所以很危险。对于必须要改变的出参，最好也先使用局部变量，最后再将该局部变量赋值给该出参。

**【规则7-2-6】如果参数列表中若干个相邻的参数类型相同，则可以在参数列表中省略前面变量的类型声明。**

**正例：**

```
func Add(a, b int)(int, error) {
    // ...
}

```

**【规则7-2-7】当 channel 作为函数参数时，根据最小权限原则，使用单向 channel。**

**说明：** 从设计的角度考虑，所有的代码应该都遵循“最小权限原则”。

**正例：**在函数Parse中ch不会被改写

```
func Parse(ch <-chan int) {
    for value := range ch {
        fmt.Println("Parsing value", value)
    }
}

```

### 返回值

**【规则7-3-1】 返回值的个数不要大于3。**

## 错误和异常设计

### 错误设计

**【规则8-1-1】 错误值统一分组定义，而不是跟着感觉走。**

**说明：** 很多人写代码时，到处return errors.New(value)，而错误value在表达同一个含义时也可能形式不同，比如“记录不存在”的错误value可能为：

1. "record is not existed."
2. "record is not exist!"
3. "###record is not existed！！！"

这使得相同的错误value撒在一大片代码里，当上层函数要对特定错误value进行统一处理时，需要漫游所有下层代码，以保证错误value统一，不幸的是有时会有漏网之鱼，而且这种方式严重阻碍了错误value的重构。
于是，我们可以参考C/C++的错误码定义文件，在Golang的每个包中增加一个错误对象定义文件，对于共性的错误对象定义，则放在公共的目录中。

**正例：**

```
// file error object
var （
    ErrEof = errors.New("EOF")
    ErrClosedPipe = errors.New("io: read/write on closed pipe")
    ErrShortBuffer = errors.New("short buffer")
    ErrShortWrite = errors.New("short write")
）

// port error object
var (
    ErrAddPortToOvsBriFailed = errors.New("add port to ovs bri failed")
    ErrPortNotEnteredPromiscuousMod = errors.New("port not entered promiscuous mod")
    ...
)

```

**【规则8-1-2】 失败的原因只有一个时，不使用error。**

**正例：**

```
func (self *AgentContext) IsValidHostType(hostType string) bool {
    return hostType == "virtual_machine" || hostType == "bare_metal"
}

```

**反例：**

```
func (self *AgentContext) CheckHostType(hostType string) error {
    switch hostType {
    case "virtual_machine":
        return nil
    case "bare_metal":
        return nil
    }
    return ErrInvalidHostType
}

```

**【规则8-1-3】 没有失败原因时，不使用error。**

**说明：** error在Golang中是如此的流行，以至于很多人设计函数时不管三七二十一都使用error，即使没有一个失败原因，而该函数的调用者无疑是无奈的。

**正例：**
函数设计：

```
func (self *CniParam) setTenantId() {
    self.TenantId = self.PodNs
}

```

函数调用：

```
self.setTenantId()

```

**反例：**
函数设计：

```
func (self *CniParam) setTenantId() error {
    self.TenantId = self.PodNs
    return nil
}

```

函数调用：

```
err := self.setTenantId()
if err != nil {
    // log
    // free resource
    return ERR_SET_ID_FAILED
}

```

**【规则8-1-4】 error/bool应放在返回值类型列表的最后。**

**正例：**

```
resp, err := http.Get(url)
if err != nil {
    return nill, err
}

value, ok := cache.Lookup(key) 
if !ok {
    // ...cache[key] does not exist… 
}

```

**【规则8-1-5】 错误逐层传递时，层层都加日志。**

**【规则8-1-6】 错误处理巧用defer。**

**正例：**

```
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ErrCreateResource1Failed
    }
    defer func() {
        if err != nil {
            destroyResource1()
        }
    }()
    err = createResource2()
    if err != nil {
        return ErrCreateResource2Failed
    }
    defer func() {
        if err != nil {
            destroyResource2()
        }
    }()

    err = createResource3()
    if err != nil {
        return ErrCreateResource3Failed
    }
    defer func() {
        if err != nil {
            destroyResource3()
        }
    }()

    err = createResource4()
    if err != nil {
        return ErrCreateResourc4Failed
    }
    return nil
}

```

**反例：**

```
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ErrCreateResource1Failed
    }
    err = createResource2()
    if err != nil {
        destroyResource1()
        return ErrCreateResource2Failed
    }

    err = createResource3()
    if err != nil {
        destroyResource1()
        destroyResource2()
        return ErrCreateResource3Failed
    }

    err = createResource4()
    if err != nil {
        destroyResource1()
        destroyResource2()
        destroyResource3()
        return ErrCreateResource4Failed
    }
    return nil
}

```

**【规则8-1-7】 当尝试几次可以避免失败时，不要立即返回错误。**

**说明：** 如果错误的发生是偶然性的，或由不可预知的问题导致。一个明智的选择是重新尝试失败的操作，有时第二次或第三次尝试时会成功。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。比如我们平时上网时，尝试请求某个URL，有时第一次没有响应，当我们再次刷新时，就有了惊喜。

**【规则8-1-8】 当上层函数不关心错误时，不返回error。**

**说明：** 对于一些资源清理相关的函数（destroy/delete/clear），如果子函数出错，打印日志即可，而无需将错误进一步反馈到上层函数，因为一般情况下，上层函数是不关心执行结果的，或者即使关心也无能为力，于是我们建议将相关函数设计为不返回error。

### 异常设计

**【规则8-2-1】 在程序开发阶段，坚持速错，让程序异常崩溃。**

**说明：** 所谓速错简单来讲就是“让它挂”，只有挂了你才会第一时间知道错误。在早期开发以及任何发布阶段之前，最简单的同时也可能是最好的方法是调用panic函数来中断程序的执行以强制发生错误，使得该错误不会被忽略，因而能够被尽快修复。

**【规则8-2-2】 在程序部署后，应恢复异常避免程序终止。**

**说明：** 在Golang中，虽然有类似Erlang进程的Goroutine，但需要强调的是Erlang的挂，只是Erlang进程的异常退出，不会导致整个Erlang节点退出，所以它挂的影响层面比较低，而Goroutine如果panic了，并且没有recover，那么整个Golang进程（类似Erlang节点）就会异常退出。所以，一旦Golang程序部署后，在任何情况下发生的异常都不应该导致程序异常退出，我们在上层函数FuncA中加一个延迟执行的recover调用来达到这个目的，并且是否进行recover需要根据环境变量或配置文件来定，默认需要recover。

**正例：**

```
func FuncA() (err error) {
    defer func() {
        if permittedRecover {
            if p := recover(); p != nil {
                fmt.Println("panic recover! p:", p)
                str, ok := p.(string)
                if ok {
                    err = errors.New(str)
                } else {
                    err = errors.New("panic")
                }
                debug.PrintStack()
            }
        }
    }()
    ...
}

```

**注：有时需要在延迟函数中释放资源，比如该携程在异常前进行了read channel操作，由于异常的发生使得该携程没有完成write channel操作，这会使得该channel后续的操作阻塞，所以必须在延迟函数中根据标志位进行write channel操作，以便操作始终都是闭合的。**

**【规则8-2-3】 对于不应该出现的分支，使用异常处理。**

**说明：** 当某些不应该发生的场景发生时，我们就应该调用panic函数来触发异常。

**正例：**

```
switch s := suit(drawCard()); s {
    case "Spades":
    // ...
    case "Hearts":
    // ...
    case "Diamonds":
    // ...
    case "Clubs":
    // ...
    default:
        panic(fmt.Sprintf("invalid suit %v", s))
}

```

**【规则8-2-4】 针对入参不应该有问题的函数，使用异常设计。**

**说明：** 入参不应该有问题一般指的是硬编码，而不是API的外部输入。当调用者明确知道输入不会引起函数错误时，要求调用者检查这个错误是不必要和累赘的。我们应该假设函数的输入一直合法，当调用者输入了不应该出现的输入时，就触发panic异常。

**正例：** 库函数MustCompile的实现

```
func MustCompile(str string) *Regexp {
    regexp, error := Compile(str)
    if error != nil {
        panic(`regexp: Compile(` + quote(str) + `): ` + error.Error())
    }
    return regexp
}

```

## 整洁测试

**【规则9-1】 不要为了测试而对产品代码进行侵入性修改。**

**说明：** 禁止为了测试而在产品代码中增加条件分支或函数变量。

**【建议9-1】 测试用例中不应该存在复杂的循环和条件控制语句。**

**说明：** 测试用例对可读性的要求非常高，如果出现大量的循环、条件控制语句，将大大地损害了用例的可读性。一般地，测试用例应该是由若干条陈述句所组成，越简单越好。

**【建议9-2】 测试代码和生产代码一样重要。**

**说明：** 测试代码不是二等公民，它需要被思考、被设计和被照料，它该像产品代码一般保持整洁。

**【建议9-3】 整洁的测试有三个要素：可读性，可读性和可读性。**

**说明：** 在测试代码中，可读性甚至比生产代码还重要。生产代码的正确性由测试代码来保证，而测试代码的正确性只能由自己的可读性来保证，让错误无处藏身。

**【建议9-4】 测试应该是黑盒的。**

**说明：** 避免根据代码编写测试。

** XP

© 著作权归作者所有

[举报文章]()

![96](http://upload.jianshu.io/users/upload_avatars/2463211/80337da84cfb.jpg?imageMogr2/auto-orient/strip|imageView2/1/w/96/h/96)

**已关注

_张晓龙_

写了 81772 字，被 302 人关注，获得了 319 个喜欢

程序员，开源软件爱好者

[**喜欢]()

 

[11]()

[**]() [**](javascript:void((function(s,d,e,r,l,p,t,z,c){var%20f='http://v.t.sina.com.cn/share/share.php?appkey=1881139527',u=z||d.location,p=['&url=',e(u),'&title=',e(t||d.title),'&source=',e(r),'&sourceUrl=',e(l),'&content=',c||'gb2312','&pic=',e(p||'')].join('');function%20a(){if(!window.open([f,p].join(''),'mb',['toolbar=0,status=0,resizable=1,width=440,height=430,left=',(s.width-440)/2,',top=',(s.height-430)/2].join('')))u.href=[f,p].join('');};if(/Firefox/.test(navigator.userAgent))setTimeout(a,0);else%20a();})(screen,document,encodeURIComponent,'','','http://cwb.assets.jianshu.io/notes/images/7790730/weibo/image_0351746a1000.png',%20'%E6%8E%A8%E8%8D%90%20_%E5%BC%A0%E6%99%93%E9%BE%99_%20%E7%9A%84%E6%96%87%E7%AB%A0%E3%80%8AGolang%20Programming%20Style%E3%80%8B%EF%BC%88%20%E5%88%86%E4%BA%AB%E8%87%AA%20@%E7%AE%80%E4%B9%A6%20%EF%BC%89','http://www.jianshu.com/p/4540bb8fc9b5?utm_campaign=maleskine&utm_content=note&utm_medium=reader_share&utm_source=weibo','%E9%A1%B5%E9%9D%A2%E7%BC%96%E7%A0%81gb2312|utf-8%E9%BB%98%E8%AE%A4gb2312'));) [**](http://cwb.assets.jianshu.io/notes/images/7790730/weibo/image_0351746a1000.png) [更多分享](javascript:void(0);)

- [**]()
- [**]()
- [**](javascript:void(0);)

被以下专题收入，发现更多相似内容

[**收入我的专题]()

![img](http://upload.jianshu.io/collections/images/30374/%E6%9C%AA%E6%A0%87%E9%A2%98-1.jpg?imageMogr2/auto-orient/strip|imageView2/1/w/64/h/64)Go语言

![img](http://upload.jianshu.io/collections/images/541882/golangweb-logo.png?imageMogr2/auto-orient/strip|imageView2/1/w/64/h/64)Golang语言社区