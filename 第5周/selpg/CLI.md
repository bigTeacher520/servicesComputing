# 使用golang开发Linux命令行实用程序中的selpg--17343098-quanao

## 1.设计与实现

**该CLI程序设计的功能:**
+ 文件的读写
+ 从终端获取输入以及在终端输出

**设计模块:**

+ 命令行参数结构体：
```
type selpg struct {
	startPage      int
	endPage        int
	file         string
	pageLen        int
	pageType       bool 
	outFile string
}
```

+ 解析参数：
    + 使用pflag包对命令行输入参数进行解析
    + 获得flag参数后pflag.Parse()函数才能把参数解析出来
    + 使用pflag.Args()来获取未被标记的参数


```
func getInput(args *selpg) {
	pflag.IntVarP(&(args.startPage), "startPage", "s", -1, "start page")
	pflag.IntVarP(&(args.endPage), "endPage", "e", -1, "end page")
	pflag.IntVarP(&(args.pageLen), "pageLen", "l", 72, "the length of page")
	pflag.BoolVarP(&(args.pageType), "pageType", "f", false, "page type")
	pflag.StringVarP(&(args.outFile), "outFile", "d", "", "print destination")
	pflag.Parse()
	mid := pflag.Args()
	args.file = "" 
	if len(mid) > 0 {
		args.file = mid[0]
	} 
}
```

+ 检查参数：
    + 是否输入了起始页和结束页
    + 起始页大于1小于结束页以及不能溢出（MaxInt32）
    + 结束页大于起始页并且不能溢出（MaxInt32）
    + 遇到不合法则输出错误同时结束程序
```
func checkInput(args *selpg) {
	if (args.startPage == -1 || args.endPage == -1) ||  (args.startPage < 1 || args.startPage > math.MaxInt32)||(args.endPage < 1 || args.endPage > math.MaxInt32 || args.endPage < args.startPage)||(!args.pageType) && (args.pageLen < 1 || args.pageLen > math.MaxInt32){
		os.Stderr.Write([]byte("You shouid input valid argument\n"))
		os.Exit(0)
	}
}

```

+ 处理输入

```
func processInput(args *selpg) {
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)
	if args.file != "" {
		fileIn, err := os.Open(args.file)
		defer fileIn.Close()
		if err != nil {
			os.Stderr.Write([]byte("Open file error\n"))
			os.Exit(0)
		}
		reader = bufio.NewReader(fileIn)
	}
	if args.outFile == "" {
		outputCurrent(reader, args)
	} else {
		outputToFile(reader, args)
	}
}
```
## 2.测试
1. 导入包编译程序：
![avatar](./Selpg/src/sel_img/compile.png)

2. 输出指定文件的页：
>selpg -s1 -e1 selpg.go
![avatar](./Selpg/src/sel_img/testPrint.png)

3. 重定义输入：
>selpg -s1 -e1 < selpg.go
![avatar](./Selpg/src/sel_img/defineInput.png)

4. 重定义输出：
>selpg -s10 -e20 input_file >output_file
![avatar](./Selpg/src/sel_img/definePrint.png)

5. 管道操作符：
>other_command | selpg -s10 -e20
![avatar](./Selpg/src/sel_img/boolPrint.png)

6. 输出报错信息：
> selpg -s10 -e20 input_file 2>error_file
![avatar](./Selpg/src/sel_img/errorPrint.png)

7. selpg 完成工作，并且您希望对输出和错误都进行保存:
>selpg -s10 -e20 input_file >output_file 2>error_file
![avatar](./Selpg/src/sel_img/errorandprint.png)

8. >  selpg -s10 -e20 input_file | other_command
![avatar](./Selpg/src/sel_img/sort.png)

9. >selpg -s10 -e20 input_file 2>error_file | other_command
![avatar](./Selpg/src/sel_img/sortanderror.png)

10. >selpg -s10 -e20 -l66 input_file
![avatar](./Selpg/src/sel_img/selPrint.png)
