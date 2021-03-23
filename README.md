# riri

被啰嗦到网络请求烦到了, 内心难受

> 人生一片黑暗的社畜又替最为厌恶的人背上了无力偿还的债务, 让本就在贫困线边缘挣扎的自闭five更加绝望
>
> 在不确定的未来删号重练发生前, 妄图挤些代码, 自欺欺人

参考平时用的`Postman`, 链式调用的网络请求

# 莫名其妙

目前只能`struct`套一层`interface` 可惜不能无限套娃

暂时不想传变量, 说不定后续会改成`struct`套娃

使用之前瞎写的[错误追踪](https://github.com/M-Quadra/kazaana)来输出错误信息

使用[gin](github.com/gin-gonic/gin)作为测试

为了方便区分环境, 补充了`riri.Group`用来生成URL, 后续可能会做进一步集成

# 花里胡哨

```
result, kerr := GET(url).Params.Set(map[string]string{
    "1": "2",
}).Result()
```

.Result() 直接返回`[]byte`(body)与`kazaana.Error`(保存调用栈的错误回调)

若要更接近原生返回值, 可以使用`.Do()`

# 拖泥带水

随缘更新, 由于思路比较混乱, 方法不会咋地稳定, 基本属于瞎胡闹, 精神状态好了可能会补上说明

- interface

`interface`的缺点估计就在于没法非常完美地区分方法, 但使用`struct`套娃又会引入除必要方法外多余的结构体, 后续发展方向还未考虑好

- Group

对Group的进一步设想如下

```
tsGroup := riri.Group(...)
resData, kerr := tsGroup.Path("/11").Path("/22").GET.Result()
```

极大概率由于拖延与懒癌用不到

- 关于错误

是否应该根据http状态码直接抛出错误? 还是手动处理结果?

- 解析模版

有许多简单请求后的处理行为完全一致, 是否应该想办法优化一下这类可能出现的重复代码?

- http/2

应该默认开启`http/2`, 还是自定义`Client`?