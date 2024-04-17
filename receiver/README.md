## 功能
1. 接收和处理由device模块push到kafka消息队列的数据
2. 将数据写入到mysql数据库
3. 直接调用device模块的rpc接口,把数据推送到device模块

## 处理函数注册过程

来自device模块的数据称为消息，不同的消息由消息Code字段标记。在模块中，每个Code对应一个闭包，闭包返回函数处理对应的消息。所有闭包组织在一个map中，map的key为code。当消息需要处理的时候，由code 找到对应的闭包，闭包返回函数处理对应的消息。

如果需要增加一种消息类型，增加对应的处理函数：
1. 在logic目录下写一个闭包。
2. 在kqServer.RegisterHandler函数中把闭包添加到map中。

