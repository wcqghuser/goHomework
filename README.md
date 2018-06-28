 # GO练习题

 ## RUN
 
```shell
  git clone https://github.com/wcqghuser/goHomework.git
  cd goHomework
```
 
> 1.1 打印一个菱形

```
  go test rhomb/main_test.go rhomb/main.go  
  go run rhomb/main.go 7 3
``` 

> 1.2 输入一行字符,分别统计出其中英文字母、空格、数字和其它字符的个数

```shell
  go test statistical/main_test.go statistical/main.go  
  go run statistical/main.go "fdsfsfsd你好33  fdsdf"
```

> 1.3 构造一个 map 以学号为 key 存储学生信息(姓名、性别、学号、年龄、成绩),产生 10 个学生信息并存入 map,然后将所有学生信息取出,按成绩排序(由高到低)存入一 个切片中,然后按顺序输出学生信息
```shell
  go run studentInfo/main.go
```

> 2.1 定义一个底层类型为 int 类型的类型别名,实现调用这个新类型的某个方法, 这个类型的对应的对象值增加 200。例如:a:=200, 调用 a.Increase 后 a 的值变成 400
```shell
  go test increaseInt/main_test.go increaseInt/main.go
  go run increaseInt/main.go 200
```

> 2.2 使用面向对象的方式实现 zhangSan 开车去银行,然后拿着银行卡去银行取钱这样的 场景。编码实现中请加入接口技术(zhangSan 持有汽车的驾驶接口,而不是持有汽车对 象指针)
```shell
  go run oop/main.go
```

> 2.3 将 runtime.GOMAXPROCS 设置为 runtime.NumCPU, 然后启动 10 个 goroutine,按顺 序为每个 goroutine 分配一个编号(1~10),每个 goroutine 分别打印自己的编号,打印完 成之后 goroutine 退出,使用一定的方法保证打印结果是顺序的(1 2 3 4 5 6...10)。
```shell
  go run goroutine/main.go
```

> 2.4 实例化一个如下的类型,并使用反射的方式调用 SetScore 方法重新设置 Score 成员变 量的值,然后使用反射的方法打印对象的所有信息(成员变量的类型名称和值)。
```shell
  go run reflect/main.go
```