# JavaScript 回调函数：我一直在等你

回调函数（callback function）在 JavaScript 中无处不在，许多刚接触回调概念的同学可能会感到懵，为了理解回调的概念，我们先看几个使用场景。

## 从事件看回调函数

事件从字面意义上来说，就是发生了一些事情。站在不同的角度有不同的感受：从当局者来看：完了，要上新闻了；从路人来看：好耶，又可以吃瓜了；从管理者来看：麻了，又有闹事的，得赶紧把事情处置了。

![image-20221111101228315](image-20221111101228315.png)

计算机中存在各式各样的事件，不同于吃瓜事件，计算机中的事件大多是可预测的，并且计算机里的事件有特定的格式，称作事件对象，事件对象描述了事件的细节：例如事件类型，事件源等。

![image-20221111120257309](image-20221111120257309.png)

当处理系统中的事件时，可以这么干：

```js
function handleClick(event) {
  //处理点击事件的函数
}

function handleMousemove(event) {
  //处理鼠标移动事件的函数
}

var events = []; //事件队列

//系统事件循环
while (true) {
  // 处理每一个事件
  for (const event of events) {
    switch (event.type) {
      case "click":
        handleClick(event);
        break;
      case "mousemove":
        handleMousemove(event);
        break;
      default:
      //...
    }

    // 清空事件队列
    events.clear();
  }
}
```

现在只要有事件添加到事件队列里边，就能被系统监视循环处理掉。有的同学可能不理解**系统事件循环**，它就像值班室一样，每时每刻都有人值班，一旦有什么事情发生就能快速处理。



但是这么做有许多弊端：不能定制事件处理函数，事件的处理函数是提前写好的，不能根据具体的对象调用自己的处理函数。要是能根据不同事件调用自己的处理函数该多好啊。要完成这个任务，需要事件源事先保存好自己的事件处理函数，等到事件触发的时候可以供系统使用：

```js
btn.addEventListener("click", function (event) {
  //事件处理函数
});

//系统监视事件循环
while (true) {
  // 处理每一个事件
  for (const event of events) {
    const handlers = event.target.getHandlersByType(event.type);

    for (const handler of handlers) {
      handler(event);
    }
  }

  // 清空事件队列
  events.clear();
}
```

如果不考虑系统部分，那么整个程序就变成了：

```js
btn.addEventListener("click", function (event) {
  //事件处理函数
});
```

以上案例，回调函数只字未提，但你已经能感受到回调的思想了：先保存函数，等我要用的时候再回来调用。

如果用英语说就是：Save the function, the system will **callback** later.

## 从编程风格看回调函数

习惯与经验能帮助我们快速解决问题，但也可能禁锢我们的思维，导致定式思维。大多数同学接触的编程方式是命令式的方式：

【例】你会做紫菜蛋汤吗？

```js
const ingredients = "紫菜, 鸡蛋";
const prepared = handle(ingredients);
const cooked = cook(prepared);
const meal = serve(cooked);
```

命令式的思维就是把自己看一个厨师，把数据看成菜，先处理食材（打鸡蛋，泡紫菜），再烹饪食材，最后上桌。

假如我们要开一个饭店，但我们自己不会择菜，切菜，炒菜，请问要怎么办呢？这时候，我们的很容易想到的是请一个厨师：

```js
const ingredients = "紫菜, 鸡蛋";

function chief(){
  //...
}

const meal = cook(ingredients, chief);
```

同样也能吃到美味的紫菜蛋汤，这种程序设计风格比命令式的程序更加灵活，更加优雅。如果我想喝咖喱味的紫菜蛋汤，只需要请一个印度厨师，如果想吃原汁原味的紫菜蛋汤，只需要把妈妈叫过来就行了，这样是不是很方便呀。



这种编程风格称为函数式编程风格，熟悉函数式编程语言的同学对这种程序的设计方式习以为常，经典的函数式语言有 Lisp，haskell，OCaml。在函数式编程的世界里，**数据与运算是对等的**：数据可以用运算表示，运算可以当做数据传递。

从这个角度来看，回调函数就是一个作为另一个函数参数的函数而已。

【例】数组中的函数式编程。

```js
const fruits = ["apple", "pear", "grape"];
fruits.forEach(function (fruit) {
  console.log("eat: " + fruit);
});
```

函数式编程风格不需要关心每个数组元素是怎么取到的，只需要关心拿到每个数据之后需要怎么做就行了。

【例】请求数据。

```js
$.ajax({
  url: "api/user",
  success: function (data) {
    // 数据处理回调函数
  },
  error: function (err) {
    //错误处理回调函数
  },
});
```

我们不需要数据是怎么请求的，只需要知道怎么处理请求成功后的数据与请求失败之后错误即可。

### JavaScript 函数是对象

```js
function add(a, b) {
  return a + b;
}

console.log(add instanceof Object); //true
console.log(add instanceof Function); //true
```

JavaScript 函数既然是对象，那么它理所当然地能像普通对象那样被当做参数传递。

## 参考文章

> ♥ 我是前端工程师：你的甜心森。非常感谢大家的点赞与关注，欢迎大家参与讨论或协作。
>
> ★ 本文[开源](https://github.com/xiayulu/FrontEndCultivation)，采用 [CC BY-SA 4.0 协议](http://creativecommons.org/licenses/by-sa/4.0/)，转载请注明出处：[前端工程师的自我修养](https://github.com/xiayulu/FrontEndCultivation). GitHub.com@xiayulu.
>
> ★ 创作合作或招聘信息请发私信或邮件：zuiaiqiansen@163.com，注明主题：创作合作或**招聘前端工程师**。
