# sync.Map

```go
type Map struct {
   mu Mutex
    
   read atomic.Value // readOnly

   dirty map[interface{}]*entry

   misses int
}
```

冗余结构 read map 和dirty map

读取数据先从read map读取，无锁

读不到加锁从dirty map读

删除时如果read map存在key，则将value置空，避免后续读该key时加锁



# sync.Pool

```go
type localpool{

​	private interface{}

​	shared []interface{}

​	mutex

}
```

每个P对应一个localpool

Put: private有值则追加到shared

Get: 本地取不到，遍历其他P的shared偷一个

# sync.Mutex

```go
type Mutex struct {

​    state int32

​    sema  uint32

}
```

正常模式 

饥饿模式 有goroutine等锁时间大于1ms

# channel

循环队列

发送、接收下标

发送、接收阻塞队列

不应该依赖顺序（发送时，如果阻塞队列不空，直接把数据拷给阻塞goroutine）

# interface

```go
type iface struct {

​    tab  *itab

​    data unsafe.Pointer

}

type eface struct {
    _type *_type
    data  unsafe.Pointer
}

type itab struct {

​    inter  *interfacetype

​    _type  *_type

​    link   *itab

​    hash   uint32 // copy of _type.hash. Used for type switches.

​    bad    bool   // type does not implement interface

​    inhash bool   // has this itab been added to hash?

​    unused [2]byte

​    fun    [1]uintptr // variable sized

}

type interfacetype struct {
​    typ     _type

​    pkgpath name

​    mhdr    []imethod

}

type _type struct {
    // 类型大小
    size       uintptr
    ptrdata    uintptr
    // 类型的 hash 值
    hash       uint32
    // 类型的 flag，和反射相关
    tflag      tflag
    // 内存对齐相关
    align      uint8
    fieldalign uint8
    // 类型的编号，有bool, slice, struct 等等等等
    kind       uint8
    alg        *typeAlg
    // gc 相关
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}
```



- iface  保存接口约束、class、object
- eface 保存class 和 object

# defer

- 匿名

1. var temp = 要返回的值
2. 调用defer 函数
3. return temp

- 有名

1. 定义的返回变量 = xxx
2. 调用defer函数
3. return 

# 闭包

- 不引用局部变量= 普通函数
- 引用外部局部变量= 打包所引用局部变量地址（注意for循环闭包）

### 

# 调度

G、M、P对应runtime数据结构

### M

对应os thread

m0、普通线程、sysmon监控管理线程（强制GC、抢占调度）

从本地队列拿可运行G（1/64概率从全局队列拿）

拿不到从其他P偷一半G过来、其他P也没有再从全局队列拿

状态有： 

自旋（绑定P，没有可运行G）

休眠（未绑定P，所有P上已经绑定其他M）

阻塞（运行的G发起同步系统调用，让出P）

### P

逻辑核心，控制并发度，为M运行提供环境

本地队列

不共享的数据可以不加锁

### G

go关键字创建出轻量级线程，保存运行上下文

g0运行调度程序

创建后放入P的本地队列

本地队列满时把前一半移动到全局队列

移到全局队列的部分顺序会被打乱



# 内存管理

