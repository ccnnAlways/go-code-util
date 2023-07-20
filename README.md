# go-code-util



### uttputil

#### GET请求

方法：`func SendGet(u string, data url.Values) (body []byte, err error)`

参数：

* u：url，url后面不能跟参数，参数需要放到`data`中
* data，参数



#### Post请求

方法：`func SendPost(u string, data interface{}) (body []byte, err error)`

参数：

* u：url
* data，参数



#### 待验证

- [ ] post请求url中，带的参数能否被解析到
- [ ] get请求中，url带参数同时，data也有参数，能否被解析到
