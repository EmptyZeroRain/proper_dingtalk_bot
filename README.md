# dingtalk
项目根据https://github.com/CatchZeng/dingtalk
根据实际需求进行了修改

DingTalk(dingding) 是钉钉机器人的 go 实现。支持 Docker、Jenkinsfile、命令行模式，module 模式；支持加签安全设置，支持链式语法创建消息；支持文本、链接、Markdown、ActionCard、FeedCard 消息类型。

# 专属钉钉
支持所有专属钉钉，政务钉钉，私有化钉钉，修改internal/security/security.go中的const dingTalkOAPI 替换为机器人API接口地址。
