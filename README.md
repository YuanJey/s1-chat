# s1-chat

# 设计思想
    handles：自定义实现后通过set函数注入到manage,来实现自己的业务


# 多数据中心
    clusterManage,负责处理跨数据中心的消息,当用户上线,server 通知clusterManage,clusterManage再把未同步的消息分发给server