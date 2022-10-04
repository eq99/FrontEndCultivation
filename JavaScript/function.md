# 函数

## 函数的声明

【例】

```js
const f = function g() {
  return 23;
};
typeof g();
```

在 JS 里，声明函数只有 2 种方法：
第 1 种： function foo(){...} （函数声明）
第 2 种： var foo = function(){...} （等号后面必须是匿名函数，这句实质是函数表达式）

除此之外，类似于 var foo = function bar(){...} 这样的东西统一按 2 方法处理，即在函数外部无法通过 bar 访问到函数，因为这已经变成了一个表达式。

如果是typeof f，结果是function

如果是typeof f()，结果是number

如果是typeof g,结果是undefined.

如果是typeof g(),结果是ReferenceError，g is not defined

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

