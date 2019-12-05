# antblog  
[示例地址](http://180.76.114.111/)

## 初衷
   作为一名十几年的程序员居然没有自己的网站，虽然早就想搞一个自己的网站，但是一直未能执行。这次下决心要做个自己的网站了，准备把自己学习到的东西用这种方式做记录和练习。

## 技术选型
选型需求：
1. 后端开发语言选择GO（也许是第六灵感或者观察到的事项，比如行业的字节跳动用Go替代PHP，以及很多其他的公司也在从PHP转Go，作为PHP程序员为了吃饭得考虑转了）
2. 开发网站首先需要个博客;选择一个开源的博客最合适，开源的博客最好简单易用，也方便自己后面改造；
3. 需要考虑服务端GO框架的性能&开发的效率&学习的难度&以及当前Go行业的选择；
4. 一直对前端不太熟悉，最近也看了typescript /javascript ，以及小程序开发，后面会考虑网站支持小程序/pc/web，后面甚至会支持nativeapp方式。也就是说希望自己能成为个全栈的工程师
5. 这个项目后面不仅仅是个blog，后面会支持各种形式的功能，实现的功能也会在blog上介绍。

根据上面的要求，在baidu上搜索各种资料，最终选择了wangsongyan开源的Wblog(https://github.com/wangsongyan/wblog)，谢谢wangsongyuan。
下面是Wblog的几个技术点，具体情况可以到github看wangsongyuan的项目介绍：   
1. web:[gin](https://github.com/gin-gonic/gin)
2. orm:[gorm](https://github.com/jinzhu/gorm)
3. database:[sqlite3](https://github.com/mattn/go-sqlite3)
4. ~~全文检索:[wukong](https://github.com/huichen/wukong)~~
5. 文件存储:~~[七牛云存储](https://www.qiniu.com/)~~[smms图床](https://sm.ms)
6. 配置文件 [go-yaml](https://github.com/go-yaml/yaml)

我这边初期会把database改成mysql，文件存储早期用redis。wblog的部分改动，参考https://www.cnblogs.com/bugmaking/p/9083311.html（上述文章可能是转载，原始的文章应该是http://www.bugclosed.com/post/14，不过cnblog的样式看起来舒服）

初步的规划

1: 数据库方面会选择redis + mysql + mongodb + hbase等，这些db分别支持不同情形的使用，这也是一个互联网公司项目常见的存储模型    
  
    redis提供性能方面的存储，
    mysql提供一般的业务存储，
    mongodb提供nosql方面的存储，
    hbase方面是支持大数据相关，比如后面对业务日志做分析实验。

2: 后端服务采用router(nginx) + inrouter(nginx) + go server方式
     
    router做外围接入层，支持安全策略（比如IP、UID白名单、限制策略、一般性限流、过载保护等），
    inrouter支持静态文件访问，内部的路由分发策略等
    go server提供真正的业务（采用gin框架，后端管理系统使用gorm，配置采用yaml，服务端配置推送到服务器采用zookeeper，服务守护采用Supervisor以及日志的采集推送监控等等）

3: 前端后面预计会采用typescript(主要是看起来跟OO很像，用起来可能稍微方便些，不过也是太庞杂了)；

4: 系统部署上，后面也会采用多云策略（百度云、腾讯云等）做稳定性考量，会针对系统稳定性做些尝试。 目前已购买百度云服务器（http://180.76.114.111/，域名antracewalking.fun[10年，审核中]）

5: 后面预计后面也会做产品生产相关的，比如代码使用github存储，项目工程管理、代码打包上线回滚系统，线上系统的扩缩容等，反正涉及到的东西都想尝试下，这个范围有点太大，可能难以实现。



## 项目结构
这个是当前结构，与wblog一致，后面预计会修改
```
-antblog
    |-conf 配置文件目录
    |-controllers 控制器目录
    |-helpders 公共方法目录
    |-models 数据库访问目录
    |-static 静态资源目录
        |-css css文件目录
        |-images 图片目录
        |-js js文件目录
        |-libs js类库
    |-system 系统配置文件加载目录
    |-tests 测试目录
    |-vendor 项目依赖其他开源项目目录
    |-views 模板文件目录
    |-main.go 程序执行入口
```
## TODO
- [ ] 系统日志
- [ ] 网站统计
- [x] 文章、页面访问统计
- [x] github登录发表评论
- [x] rss
- [x] 定时备份系统数据
- [x] 邮箱订阅功能

## 安装部署
本项目使用govendor管理依赖包，[govendor](https://github.com/kardianos/govendor)安装方法 [感觉不太好使]
```
go get -u github.com/kardianos/govendor
```

```
git clone https://github.com/antracewalking/antblog
cd antblog
govendor sync
go run main.go
```

简单的办法是下载代码后，go build(or go run main.go) 看看那些依赖包没有，就一个个下载吧

## 使用方法
### 使用说明
1. 修改conf.yaml，设置signup_enabled: true
2. 访问http://xxx.xxx/signup 注册管理员账号
3. 修改conf.yaml，设置signup_enabled: false

### 注意事项
1. 如果需求上传图片功能请自行申请七牛云存储空间，并修改配置文件填写
    - qiniu_accesskey
    - qiniu_secretkey
    - qiniu_fileserver 七牛访问地址
    - qiniu_bucket 空间名称
2. 如果需要github登录评论功能请自行注册[github oauthapp](https://github.com/settings/developers)，并修改配置文件填写
    - github_clientid
    - github_clientsecret
    - github_redirecturl
3. 如果需要使用邮件订阅功能，请自行填写
    - smtp_username
    - smtp_password
    - smtp_host,例如：smtp.163.com:25
4. Goland运行时，修改main.go getCurrentDirectory方法返回""

## 效果图

![file](screenshots/index.png)

![file](screenshots/blog.png)

![file](screenshots/admin.png)

## 捐赠
#如果项目对您有帮助，打赏个鸡腿吃呗！  
#<img src="https://raw.githubusercontent.com/antracewalking/antblog/master/screenshots/alipay.png" width = 40% height = 40% />
#<img src="https://raw.githubusercontent.com/antracewalking/antblog/master/screenshots/weixin.png" width = 40% height = 40% />
