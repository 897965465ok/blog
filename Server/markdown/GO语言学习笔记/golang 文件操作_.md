## 文件操作的各种包

> os


> io


> io/ioutil


> bufio


## 文件操作需要先理解的概念

> os包的内置的常量flag   代表着以什么方式打开文件


```
    os.O_RDONLY // 只读模式打开文件
    os.O_WRONLY  // 只写模式打开文件
    os.O_RDWR   // 读写模式打开文件
    os.O_APPEND // 写操作时将数据附加到文件尾部
    os.O_CREATE // 如果不存在将创建一个新文件
    os.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    os.O_SYNC   // 打开文件用于同步I/O
    os.O_TRUNC  // 如果可能，打开时清空文件
```

> os包的内置常量ModePerm 代表文件的模式和权限位


```
 // 单字符是被String方法用于格式化的属性缩写。
    os.ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    os.ModeAppend                                     // a: 只能写入，且只能写入到末尾
    os.ModeExclusive                                  // l: 用于执行
    os.ModeTemporary                                  // T: 临时文件（非备份文件）
    os.ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    os.ModeDevice                                     // D: 设备
    os.ModeNamedPipe                                  // p: 命名管道（FIFO）
    os.ModeSocket                                     // S: Unix域socket
    os.ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
    os.ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
    os.ModeCharDevice                                 // c: 字符设备，需已设置
    os.ModeDevice
    os.ModeSticky                                     // t: 只有root/创建者能删除/移动文件
    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
```

## os 包读写文件的操作


> > ### 读文件



1. 打开文件

   - 填写你需要打开的路径
   - 填写flag
   - 填写ModePerm
2. 得到文件指针
3. 用Read方法读
4. 关闭文件

```
    // 以 os.O_RDWR(读写模式)权限为ModePerm(0777) 打开文件 拿到file 
    file, error := os.OpenFile("123.txt", os.O_RDWR, os.ModePerm)
    // defer 关闭文件
	defer file.Close()
	if error != nil {
		panic(error)
	}
    // 创建临时一个容器  类型是切片 长度,和容量视机器配置而定
	size := make([]byte, 1, 1)
    // 创建一个容器 将读取出的数据放入
	result := make([]byte, 1)

	for {
        // Read 方法是将读取的内容放入size变量
		lenght, err := file.Read(size)
        // 然后将size 放入result
		result = append(result, size[:lenght]...)

		fmt.Println(string(size))
        // 读取完毕的时候  err==io.EOF lenght==0
		if lenght == 0 || err == io.EOF {

			break 
           // 结束掉
		}

	}
	fmt.Println(string(result), len(result), cap(result))
```


> > ### 写文件



1. 打开文件

   - 填写你需要打开的路径
   - 填写flag
   - 填写ModePerm
2. 得到文件指针
3. 用Write方法写
4. 关闭文件

```
filePath := "./123.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	container := []byte{65, 66, 67, 68, 69, 70}
	sizeLen, err := file.Write(container)
	fmt.Printf("%d,%v", sizeLen, err)
```


> > ### 复制文件



```
// 复制文件
// 先读取到内存 边写边读
// writer reader  一个是写的对象，一个是读的对象
func CopyFile(entry, outFile string) (int, error) {
	file1, err := os.Open(entry)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	container := make([]byte, 1024, 1024)
	n := -1 //读取的数量
	total := 0
	for {
		n, err = file1.Read(container)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("出错了")
			return total, err
		}
		total += n
		file2.Write(container[:n])
	}
	return total, nil
}
```


> > ### io包复制文件



```
func IoCopy(entry, outFile string) (int, error) {
	file1, err := os.Open(entry)
	if err != nil {
		fmt.Println("错误")
		return 0, err
	}
	file2, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("错误")
		return 0, err
	}
	defer file1.Close()
	io.Copy(file1, file2)
	return 0, err
}
```


> > ### 断点续传



```
//  移动光标
func Seeker(entry, outu string) {
	readFile, _ := os.OpenFile(entry, os.O_RDWR, os.ModePerm)
	writeFile, _ := os.OpenFile(outu, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer readFile.Close()
	defer writeFile.Close()
	container := make([]byte, 1024*10, 1024*10)
	for {
		fileInfo, _ := writeFile.Stat()
        // 获取文件字节大小
		ofss := fileInfo.Size()
		// ofss 偏移量  io.SeekStart 重文件开头计算
        // 移动光标到偏移量
		readFile.Seek(ofss, io.SeekStart)
        // 然后读取
		len, err := readFile.Read(container)
		if len == 0 || err == io.EOF {
			fmt.Println("读取完毕")
			return
		}
		writeFile.Write(container)
	}

}
```


> > ### ioutil 复制操作这个比较简单直接填路径就行



```
	file, _ := ioutil.ReadFile("./1.txt")
	ioutil.WriteFile("./2.txt", file, os.ModePerm)
```


> > ### Reader 一些函数



```
// Reader
func NewReader(entry string) {
	file, _ := os.OpenFile(entry, os.O_RDWR, 0777)
	defer file.Close()
	// reader := bufio.NewReader(file)
	// container := make([]byte, 1024)
	// // 根据字节大小读取
	// reader.Read(container)
	// 根据行读取
	// reader.ReadLine()
	// // 根据符号读取
	// reader.ReadString('\t')
	// // 读取byte
	// reader.ReadByte()
	// //根据某个符号读取byte切片
	// reader.ReadBytes('\n')
    // 读取输入端内容
	// stdin := bufio.NewReader(os.Stdin)
	// stdin.ReadString('\n')
	writeFile := bufio.NewWriter(file)
	writeFile.WriteString("hellow")
	// 刷新缓冲区
	writeFile.Flush()

}
```
