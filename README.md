# go-cli
## 为什么
在上家公司，推动使用 GO 作为新的技术选型，
## 功能
|功能|说明|
|:-:|:-|
|内置命令|内置 `help` 、`version` 等命令，提供基础支持|
|参数绑定|支持将 `flag` 参数绑定到 `struct`|
|全局参数|通过管道模型支持添加自定义的全局参数，进行统一化处理|
## 示例

### 定义命令及参数

```
type (
	HelloCommandFlags struct {
		Name      string
		ClassRoom string `flag:"usage:学生的教师信息"`
	}

	HelloCommand struct {
		*cli.BaseCommand
	}
)
```
### 创建命令实例及配置默认参数

```
var (
	defaultHelloCommandFlags = &HelloCommandFlags{
		Name: "ace",
	}

	helloCommand = &HelloCommand{
		BaseCommand: &cli.BaseCommand{
			Meta: &cli.CommandMeta{
				Name:  "hello",
				Usage: "say hello to someone",
			},
			DefaultFlags: defaultHelloCommandFlags,
		},
	}
)
```

### 给 cli 应用添加命令

```
func main() {

	builder := cli.NewBuilder("hello").
		SetUsage("say hello to someone").
		AddCommand(helloCommand)

	app := builder.Build()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Print(err)
	}
}
```

## `flag` 结构体标签

基本试用方式如下：

```
type HelloCommandFlags struct {
	 	Name      string
	 	ClassRoom string `flag:"name:class-room,usage:学生的教师信息"`
	 }
```

支持的项：
||示例|说明|
|:-:|:-|:-|
|`name`|`flag:"name:id"`|定义字段对应的命令行参数，默认是字段名称|
|`aliases`|`flag:"aliases:id\|ID"`|定义一些别名，主要是为了简化参数的输入|
|`require`|`flag:"require"`|是否为必填参数|
|`usage`|`flag:"usage:help 提示信息"`|用于自动生成 help 命令时展示的提示信息|