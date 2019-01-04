### 混合密码系统
    理念：用对称加密提高加密速度，用公钥加密配送密钥
    混合密码组成机制：
    ☛ 用对称密码加密消息
    ☛ 通过伪随机数生成器生成对称密码中使用的会话密钥
    ☛ 用公钥密码加密会话密钥
    ☛ 从混合密码系统外部赋予公钥密码加密时使用的密钥
    要求：要确保公钥密码与对称密码具有同等的强度
 ![](https://ww1.sinaimg.cn/large/007iUjdily1fyaqsn7ba3j30jj0n5add.jpg)
 ![](https://ww1.sinaimg.cn/large/007iUjdily1fyaqteohzpj30jk0lvjuq.jpg)