

## Go语言Web框架ron-web


https://geektutu.com/post/gee.html



在设计一个框架之前，需要回答框架核心解决了什么问题。



`net/http`提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。一些Web开发中简单的需求并不支持，需要手工实现。

- 动态路由：例如`hello/:name`，`hello/*`这类的规则。
- 鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现。
- 模板：没有统一简化的HTML机制。
- …

当我们离开框架，**使用基础库时，需要频繁手工处理的地方**，就是框架的价值所在。

- 路由(Routing)：将请求映射到函数，支持动态路由。例如`'/hello/:name`。
- 模板(Templates)：使用内置模板引擎提供模板渲染机制。
- 工具集(Utilites)：提供对 cookies，headers 等处理机制。
- 插件(Plugin)：Bottle本身功能有限，但提供了插件机制。可以选择安装到全局，也可以只针对某几个路由生效。
- …



### 前置知识

#### 标准库启动Web服务



#### 实现http.Handler接口



### 上下文

将路由(router)独立

- 设计`上下文(Context)`，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。







### 前缀树路由



动态路由，即一条路由规则可以匹配某一类型而非某一条固定的路由。

动态路由有很多种实现方式，支持的规则、性能等有很大的差异。例如开源的路由实现`gorouter`支持在路由规则中嵌入正则表达式，例如`/p/[0-9A-Za-z]+`，即路径中的参数仅匹配数字和字母；另一个开源实现`httprouter`就不支持正则表达式。著名的Web开源框架`gin` 在早期的版本，并没有实现自己的路由，而是直接使用了`httprouter`，后来不知道什么原因，放弃了`httprouter`，自己实现了一个版本。

实现动态路由最常用的数据结构，被称为**==前缀树(Trie树)==**：每一个节点的所有的子节点都拥有相同的前缀。





```
curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
```



### 分组控制

#### 分组的意义

分组控制(Group Control)是 Web 框架应提供的基础功能之一。所谓分组，是指路由的分组。如果没有路由分组，我们需要针对每一个路由进行控制。但是真实的业务场景中，往往某一组路由需要相似的处理。例如：

- 以`/post`开头的路由匿名可访问。
- 以`/admin`开头的路由需要鉴权。
- 以`/api`开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权。



#### 分组嵌套



### 中间件

==中间件(middlewares)==就是**非业务的技术类组件**。Web框架本身不可能去理解所有的业务，因而不可能实现所有的功能。因此，框架需要有一个**插口**，允许**用户自己定义功能**，嵌入到框架中，仿佛这个功能是框架原生支持的一样。因此，对中间件而言，需要考虑2个比较关键的点：

- **插入点**在哪？使用框架的人并不关心底层逻辑的具体实现，如果插入点太底层，中间件逻辑就会非常复杂。如果插入点离用户太近，那和用户直接定义一组函数，每次在 Handler 中手工调用没有多大的优势了。
- 中间件的**输入**是什么？中间件的输入，决定了扩展能力。暴露的参数太少，用户发挥空间有限。



### 模板(HTML Template)

#### 服务端渲染

前后端分离的开发模式，即 Web 后端提供 RESTful 接口，返回结构化的数据(通常为 JSON 或者 XML)。前端使用 AJAX 技术请求到所需的数据，利用 JavaScript 进行渲染。Vue/React 等前端框架持续火热，这种开发模式前后端解耦，优势非常突出。

- 后端专心解决资源利用，并发，数据库等问题，只需要考虑数据如何生成；前端童鞋专注于界面设计实现，只需要考虑拿到数据后如何渲染即可。

- 前后端分离另外一个优势。因为后端只关注于数据，接口返回值是结构化的，与前端解耦。同一套后端服务能够同时支撑小程序、移动APP、PC端 Web 页面，以及对外提供的接口。随着前端工程化的不断地发展，Webpack，gulp 等工具层出不穷，前端技术越来越自成体系了。

前后分离的一大问题，页面是在客户端渲染的，比如浏览器，这**对于爬虫并不友好**。Google 爬虫已经能够爬取渲染后的网页，但是短期内爬取服务端直接渲染的 HTML 页面仍是主流。

#### 静态文件(Serve Static Files)



#### HTML 模板渲染



### 错误恢复



## 分布式缓存ron-cache

> 商业世界里，现金为王；架构世界里，缓存为王。

### 为什么

直接使用键值对（`map`）缓存有什么问题呢？

1. 内存不够了怎么办？

那就随机删掉几条数据好了。随机删掉好呢？还是按照时间顺序好呢？或者是有没有其他更好的淘汰策略呢？不同数据的访问频率是不一样的，优先删除访问频率低的数据是不是更好呢？数据的访问频率可能随着时间变化，那优先删除最近最少访问的数据可能是一个更好的选择。

需要实现一个**==合理的淘汰策略==**。

2. 并发写入冲突了怎么办？

对缓存的访问，一般不可能是串行的。map 是没有并发保护的，应对并发的场景，修改操作(包括新增，更新和删除)需要加锁。

3. 单机性能不够怎么办？

单台计算机的资源是有限的，计算、存储等都是有限的。随着业务量和访问量的增加，单台机器很容易遇到瓶颈。如果利用多台计算机的资源，并行处理提高性能就要缓存应用能够支持分布式，这称为**水平扩展(scale horizontally)**。与水平扩展相对应的是**垂直扩展(scale vertically)**，即通过增加单个节点的计算、存储、带宽等，来提高系统的性能，硬件的成本和性能并非呈线性关系，大部分情况下，**分布式系统**是一个更优的选择。

### 是什么

设计一个分布式缓存系统，需要考虑**资源控制、淘汰策略、并发、分布式节点通信**等各个方面的问题。而且，针对不同的应用场景，还需要在不同的特性之间权衡，例如，是否需要支持缓存更新？还是假定缓存在淘汰之前是不允许改变的。不同的权衡对应着不同的实现。



###  LRU缓存淘汰策略

