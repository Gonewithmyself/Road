# 概览

![etcd-server](./etcd-server.png)

​										                        摘自[codedump](https://www.codedump.info/post/20181125-etcd-server/)

1.1	serverClient中handler处理客户端请求(processInternalRaftRequestOnce)

1.2	raftNode.Propose 写入proposeC,注册wait等待结果

1.3	proposeC 读取消息进入stepLeader

1.4 raftLog.append写入unstable

1.5 bcastAppend广播给其他节点（readyC通知etcdserver, 再通过tansport发送）

2.1	serverPeers中handler收到其他节点回应, raftNode.Propose 写入recvC

2.2 收到超过半数节点的回应，通过readyC通知raftNot可以提交

2.3	r.storage.Save先写入wal模块

2.4	applyc写入appy消息，

2.5	applyAll中写入持久化存储

2.6	applyEntryNormal中s.w.Trigger触发回复客户端





# raft模块

### 概览（node.run）

不断接收外部消息（recvC, proposeC, confC），进入step函数根据当前角色处理对应消息

需要上层处理的消息追加到msgs字段，等待时机封装到Ready实例通知上层

通过advancec保证当前只有一个ready实例在处理中



#### 写操作处理流程

先append到unstable.ents，广播通知其他节点

计算prs.Match（递增排序，前一半分界点即为大多数节点已复制的index）作为commitIdx

ready实例封装commitIdx、unstable.ents、committedEnts给上层进行持久化

advance中再根据ready实例中的数据更新appliedIdx等数据

### node 结构

```
type node struct {
	propc      chan msgWithResult   // 本地客户端消息
	recvc      chan pb.Message	// peer节点消息
	confc      chan pb.ConfChangeV2	// 配置变更消息
	confstatec chan pb.ConfState	// 
	readyc     chan Ready	// 通知上层
	advancec   chan struct{} // 上层已处理完一个ready实例
	tickc      chan struct{} // 上层调用推进逻辑时钟
```

实现Node接口，供上层模块调用

### raftlog 结构

![raftlog](./raftlog.png)



firstIndex、lastIndex针对所有ents(包括storage和unstable)

unstable中snapshot只有follower从leader接受快照时才存在，snapshot和ents不会共存



# 存储

### WAL

独立的goroutine提前预分配文件，通过fallocate系统调用预分配默认64M大小的文件。

文件中的内容与网络包格式类似，header存放包体长度（64bit 高8位存放padding字节，低56位才是真实包长）。

每个文件第一帧均为metadata类型

写入时，将长度和包体写入buf，再写入文件，接着调用fsync刷入磁盘

### ApplierV3

使用boltDB作为存储层，key存的是版本，value详细信息

在内存中用B树维护treeIndex，保存reversion到keyIndex的映射

keyIndex中维护key的创建版本，最后修改版本，以及多个从创建到删除的轮回（generations）

操作boltDB中间有一层缓存bucketBuffer，提升读取效率

### watcher

分为两组，一组是已经追上当前版本、另一组未追上。

根据加入时版本号与当前版本号比较

- ​	加入synced(写事务提交时，根据changes进行通知)
- ​    加入unsynced（从boltdb取出minRev到curRev+1的所有reversions，转换成events进行通知）



### lessor

持久化到lease bucket的只有 lease的ID、ttl、remain ttl。

promote中根据remain ttl设置expiry

revokeExpiredLeases 找出过期lease，通过expiredC通知上层

checkpointScheduledLeases 找出未过期，但到了checkpointInterval的lease，更新remain ttl并写入raftlog



```
	leaseMap             map[LeaseID]*Lease	// id => lease 
	leaseExpiredNotifier *LeaseExpiredNotifier	//  定期检测过期lease
	leaseCheckpointHeap  LeaseQueue	// 定期更新remain ttl
	itemMap              map[LeaseItem]LeaseID // key ==> leaseID
```



# 网络层

监听两个端口（2379用于客户端用户通信、2380内部节点通信）

节点间通信使用stream、pipeline两种通道

前者使用长连接用于常规消息，pipeline使用短连接用于传输快照数据

### http

serveHttp 中接收对端连接，传入streamWriter用于写入消息

streamReader中dial对端连接，接收对端消息

​	

### gRPC-gateway

使用cmux复用同一个端口，同时gRPC和HTTP服务