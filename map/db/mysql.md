### 数据结构

16K一页，页之间形成双链表

页内存放header、data、tailer

data部分存放行记录（trx_id、roll_pointer）、目录项（key、页号），行记录间形成主键递增的单链表

使用最小、最大两个虚拟指针方便定位

页目录将1-8条记录分为一组，每组存放该组最大记录指针，方便二分查找

区extent，物理连续的64页（1M）。分为直属表空间的区、与段中的区两大部分，每个部分都有三个链表（full、not_full、free）

段segment, 零散页数组32+区的链表，分为存放非叶节点的索引段与存放叶子节点的数据段。每个索引都对应着2个段

修改一个页的数据，先在buffer pool中标记为脏页，

### 事务

begin/start transcation

aotuocommit

隐式提交， DDL、修改表结构，加锁等操作、

##### redo log  保证持久性

乐观插入  页中剩余空间足够 一条redo log

悲观插入 剩余空间不足，需要进行也分裂

Mini-Transaction mtr 访问页面的原子操作，一组redo log 组成个mtr，最后一条打上tagMLOG_MULTI_REC_END

lsn 日志序列号

checkpoint_lsn  重放的起点

flushed_to_disk_lsn 刷盘的lsn



##### undo log 用于回滚

事务中每个操作生成一条undo log，undo no递增。roll_pointer指向这条链表

##### mvcc

利用undolog形成的版本链，可以得到该记录的历史版本信息

##### readView

min_ids 当前活跃的事务id  介于min、max之间且不在该列表中时**可见**

min_trx_id  min_ids中最小值   小于该值时**可见**

max_trx_id 即将分配的下一个事务ID   大于等于该值**不可见**

creator_trx_id 创建该view时的事务ID  等于该值**可见**

##### 幻读

select 某记录是否存在，不存在，准备插入此记录，但执行 insert 时发现此记录已存在，无法插入

### 查询优化



### 锁

记录锁、表锁

X、S锁、间隙锁

Innodb行锁锁定的是索引不是记录，索引相同的记录都会被加锁