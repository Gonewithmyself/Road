# 概览

![etcd-server](./etcd-server.png)

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