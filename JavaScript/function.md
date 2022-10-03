# 函数

## `bind, call, apply`

call 和 apply 都是为了改变某个函数运行时的上下文（context）而存在的，换句话说，就是为了改变函数体内部 this 的指向。对于 apply、call 二者而言，作用完全一样，只是接受参数的方式不太一样。

```js
const func = function(arg1, arg2) {
   // ...  
};

func.call(this, arg1, arg2);
func.apply(this, [arg1, arg2]);
```

### `bind`

如果将方法从对象中拿出来，然后调用，如果不做特殊处理，一般会丢失原来的对象（使得 this 指向 global 或window 对象），使用 `bind()` 方法能够很漂亮的解决这个问题： 

```js
let mywrite = document.write;
// mywrite("hello"); 
// 报错: Uncaught TypeError: Illegal invocation
// this 对象丢失

mywrite = mywrite.bind(document);
mywrite("hello");
```



# 参考文章

- 风雨后见彩虹. [JS中的call、apply、bind方法详解 ](https://www.cnblogs.com/moqiutao/p/7371988.html). 博客园.

