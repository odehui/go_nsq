1.安装 NSQ  brew install nsq

2.验证安装版本 nsqd --version

3.使用 NSQ

4.启动服务发现节点（lookupd）: nsqlookupd &

5.启动消息节点（nsqd），连接到 lookupd : nsqd --lookupd-tcp-address=127.0.0.1:4160 &

6.启动管理界面（nsqadmin）: nsqadmin --lookupd-http-address=127.0.0.1:4161 &

7.访问 http://127.0.0.1:4171 查看
