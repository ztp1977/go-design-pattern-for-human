# go-design-pattern-for-human

### links (reference)

+ [design patterns for humans](https://github.com/design-patterns-for-humans/vala#-abstract-factory)
+ [golang design pattern](https://github.com/godsarmy/golang_design_pattern)

### Usages

```bash
cd $GOPATH/src/
git clone https://github.com/ztp1977/go-design-pattern-for-human.git
cd go-design-pattern-for-human
go run -p=Bridge
```

### Think About

+ 有些设计模式有些臃肿， 不用反而效果好
+ 像singleton, flyweight, builder, iterator等模式， 其实是程序的写法
+ 通过包含构造体的指针， 可以实现假继承，前提是数据结构要一样。 要习惯没有继承的编程
+ 面向接口编程，有必要的话先定义接口

### TODO

+ 添加go独有的模式， goroutine etc,...
+ 自反式的编程， 每个动作能自动验证结果
+ 容易测试的代码，为了测试方便， 接口预留