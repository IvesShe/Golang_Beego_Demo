# Golang_Beego
學習之後，製作成筆記，方便日後使用時復習
<left class="half">
    <img src="./images/0_XoSKJTcPn78ADdUu.png" width="400"/>
</left>

# 下載GO安裝包
[GoLang](https://golang.org/dl/)

# 測試
```shell
go version
```
- ![image](./images/2020-08-01202623.png)

```shell
go env
```
- ![image](./images/2020-08-01202823.png)

# 安裝VSCode外掛
- ![image](./images/2020-08-01202948.png)

# 安裝Beego腳手架
[Beego官網](https://beego.me/quickstart)
```shell
go get -u github.com/astaxie/beego
```
- ![image](./images/2020-08-02103900.png)

```shell
go get -u github.com/beego/bee
```
- ![image](./images/2020-08-02104110.png)

安裝上若有卡住不動，可參考
[Golang Beego中没法下载第三方包解决办法](http://bbs.itying.com/topic/5ed08edee7c0790f8475e276)
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
- ![image](./images/2020-08-02103737.png)

輸入bee，確認安裝完成
```shell
bee
```
- ![image](./images/2020-08-02104217.png)

新建項目
```shell
bee new Golang_Beego_Demo
```
- ![image](./images/20200802134003.png)

運行項目
```shell
bee run
```
- ![image](./images/20200802134629.png)
- ![image](./images/20200802134735.png)

資料夾結構
- ![image](./images/20200802191658.png)

# 執行畫面示意
- ![image](./images/20200802200159.png)
- ![image](./images/20200802200135.png)