# 每次访问服务器的 ssh 的时候，都马上获得 aliyun 服务器的流量情况

![WeChat26d10bddf3a4ba2ac27611695724e07c.jpg](WeChat26d10bddf3a4ba2ac27611695724e07c.jpg)

效果：
```
Last login: Wed Sep  4 03:04:54 2024 from 183.11.68.13
+-----------------------------------+-----------------+
|               指标                |       值        |
+-----------------------------------+-----------------+
| 公网带宽                          | 3372 Mbps       |
| 公网接收的字节数                  | 5 Mb            |
| 公网发送的字节数                  | 5 Mb            |
| 内网接收的字节数                  | 5 Mb            |
| 内网发送的字节数                  | 5 Mb            |
| 内网带宽                          | 3373 Mbps       |
| [突发实例]百分百CPU剩余运行分钟数 | 118.415000 分钟 |
+-----------------------------------+-----------------+

```






# 手动安装

1. 下载 源码里的 ali 到服务器上，
2. 给权限 `chmod +x ali`
3. 执行 `./ali `，首次访问会要求提示输入阿里云的信息，输入后会在同名目录下生成一个 `config.json` 文件，保存了阿里云的信息
4. 之后再执行 `./ali auto` 会尝试将这个程序加入到每次登录的时候执行，这样每次登录都会自动获取流量信息；如果不做这一步，也可以手动执行 `./ali` 来获取流量信息


# 自动安装
