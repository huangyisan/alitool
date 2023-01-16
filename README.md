# alitool

## 实现功能
1. 域名服务
   * 域名查询
     - [x] 查询域名是否在该账号
        - alitool list domain -a AccountName -i baidu.com
       
     - [x] 列出指定阿里云账号下的全部域名
       - alitool list domain -a AccountName
       
     - [x] 反查域名隶属哪个阿里云账号
       - alitool list domain -i baidu.com -r
       
   * 域名过期查询
     - [x] 查询指定时间内域名过期,全部账号,能输出所属账号.
       - alitool check  domain -A -e 100
       
     - [x] 定点查询某个域名过期时间, 能输出所属账号.
       - alitool check  domain -d baidu.com
       
     - [x] 查询指定时间内,指定账号中,即将过期的域名
       - alitool check  domain -a AccountName -e 100
       
     - [ ] 域名ssl证书查询
       - 证书ssl查询
2. dns服务
   - [x] 查询域名dns是否在该账号
     - alitool list dns -a AccountName -i baidu.com

   - [x] 列出指定账号下的所有dns
     - alitool list dns -a AccountName

   - [x] 给定域名,反查存在哪个账号
     - alitool list dns -i baidu.com -r
     
3. 域名对应机器查询
   - [ ] cdn查询, dns解析查询

4. 财务服务
    - [ ] 查询余额
      * 指定账号,查询余额
      * 查询全体账号余额,
    - [ ] 查询指定月份消费
    - [ ] 查询指定服务消费

5. 全站加速
    - [ ] 域名配置查询
    - [ ] ssl证书替换

## 其他
- [x] 列出当前配置的阿里云账号
   * go run main.go list account
- [x] 列出region
   * go run main.go list region
- [ ] 如果ram user没有权限,则跳过, 并返回错误.
- [ ] 将提示等输出标准化
- [ ] acme account.json信息回写(可不做)
- [ ] acme 注册失败,删除accounts相关路径

## ACME流程
    //config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
    //config.CADirURL = "https://acme-v02.api.letsencrypt.org/directory"
**任何执行都需要先判断用户是否存在**
- [x] acme用户创建 (user放入yaml配置中)
  * alitool create acme 
  * alitool create acme -t (acme测试接口)
- [x] 证书签发 (读取yaml中acme用户)
  * alitool create acme -d test.com -p cloudflare

