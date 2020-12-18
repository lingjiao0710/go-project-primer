用一个棋牌游戏服务器的例子来比较完整的展现Go语言并发编程的威力。

棋牌游戏通常由一组服务器协同以支持尽量多的同时在线玩家，但由于这种分布式设计除了增加了局域网通信，从模型上与单服务器设计是一致的，
或者说只相当于把多台服务器的计算能力合并成逻辑上的一台单一服务器，所以本示例中只考虑单服务器、单进程的设计方法。

---

**项目的详细需求**
* 登陆游戏
* 查看房间列表
* 创建房间
* 加入房间
* 进行游戏
* 房间内聊天
* 游戏完成，退出房间
* 退出登陆

>棋牌游戏的特点在于房间与房间之间具备良好的个理性，这也是最能够体现并行编程威力的地方。


**系统设计**
每个玩家对应的信息如下：
* 用户唯一ID
* 用户名，用于显示
* 玩家等级
* 经验值

**总体上，我们可以将该示例代码划分为以下子系统：**

* 玩家会话管理系统，用户管理每一位登陆的玩家，包括玩家信息和玩家状态
* 大厅管理
* 房间管理系统，创建、管理和销毁每一个房间
* 游戏会话管理系统，管理房间内的所有动作，包括游戏进程和房间内聊天
* 聊天管理系统，用户接收管理员的广播信息

**为了避免贴出太多源代码，这里我们只实现了最基础的会话管理系统和聊天管理系统。因为它们足以展示以下的技术问题：**
* goroutine生命周期管理
* goroutine之间的通信
* 共享资源访问控制

```bash
$ go run cgss.go
Casual Game Server Solution
A new session has been created successfully.
Commands:
login <username><level><exp>
logout <username>
send <message>
listplayer
quit(q)
help(h)
Command> login Tom 1 101
Command> login Jerry 2 321
Command> listplayer
1 : &{Tom 1 101 0 <nil>}
2 : &{Jerry 2 321 0 <nil>}
Command> send Hello everybody.
Tom received message: Hello everybody.
Jerry received message: Hello everybody.
Command> logout Tom
Command> listplayer
1 : &{Jerry 2 321 0 <nil>}
Command> send Hello the people online.
Jerry received message: Hello the people online.
Command> logout Jerry
Command> listplayer
Failed. No player online.
Command> q
$
```

