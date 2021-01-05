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

# 花里胡哨

```
result, kerr := GET(url).Params.Set(map[string]string{
    "1": "2",
}).Result()
```

.Result() 直接返回`[]byte`(body)与`kazaana.Error`(保存调用栈的错误回调)

若要更接近原生返回值, 可以使用`.Do()`

# 磨磨唧唧

随缘更新

现有方法处于瞎胡闹阶段, 精神状态好了可能会补上说明
