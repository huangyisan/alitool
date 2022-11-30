# alitool

## 实现功能
1. 域名服务
   * 域名过期查询
     * 指定时间内域名过期,全部账号,能输出所属账号.
       * go run main.go check  domain -A -e 100
     * 定点查询某个域名过期, 能输出所属账号.
       * go run main.go check  domain -d baidu.com
   * 域名ssl证书查询
     * 证书ssl查询
2. dns服务
   * 查询域名dns是否在该账号
     * go run main.go list dns -a accountName -i baidu.com
   * 列出指定账号下的所有dns
     * go run main.go list dns -a accountName
   * 给定域名,反查存在哪个账号
     * go run main.go list dns -i baidu.com
3. 财务服务
    * 查询余额
      * 指定账号,查询余额
      * 查询全体账号余额,
    * 查询指定月份消费
    * 查询指定服务消费

## 其他
1. 列出当前配置的阿里云账号
   * go run main.go list account
