# JavaScript 条件比较与类型转换详解

JavaScript 中用于条件判断（分支流）的方法主要有如下几种：

- `if ... else if ... else`，经典的 `if-else`

```js
let lang = prompt("你想学什么语言?");

if (lang == "JavaScript") {
  alert("前端");
} else if (lang == "Java") {
  alert("后端");
} else {
  alert("不知道什么端");
}
```

【要点】不像 Go 或 Rust，JavaScript `if(...)` 后面的表达式必须要加`()`，条件体只有一条语句时，可以省略`{}`。

- `switch(...){}`

```js
const expr = "Papayas";
switch (expr) {
  case "Oranges":
    console.log("Oranges are $0.59 a pound.");
    break;
  case "Mangoes":
  case "Papayas":
    console.log("Mangoes and papayas are $2.79 a pound.");
    // expected output: "Mangoes and papayas are $2.79 a pound."
    break;
  default:
    console.log(`Sorry, we are out of ${expr}.`);
}
```

【要点】

1）`switch` 语句会根据表达式的值跳转到相应的 `case` 中，如果不加 `break`，那么 JavaScript 就会从这里开始，依次执行下面的 `case` 。这与某些语言的行为不一样，例如 Go 语言执行一条 case 会自动推出。

2）case 后面 `{}` 可加可不加，如果不加，那么 case 作用于等于 `switch` 的，如果加了就是独立的作用域：

```js
const action = "say_hello";
switch (action) {
  case "say_hello": {
    const message = "hello";
    console.log(message);
    break;
  }
  case "say_hi": {
    const message = "hi";
    console.log(message);
    break;
  }
  default: {
    console.log("Empty action received.");
  }
}
```

如果把大括号去掉会报错，因为不能在同一个作用域用  `const` 声明同一个变量两次。

【案例】使用 `switch` 模拟 `if..else if ..else`:

```js
switch (true) {
  case 'fetch' in globalThis:
    // Fetch a resource with fetch
    break;
  case 'XMLHttpRequest' in globalThis:
    // Fetch a resource with XMLHttpRequest
    break;
  default:
    // Fetch a resource with some custom AJAX logic
    break;
}
```

等价于：

```js
if ('fetch' in globalThis) {
  // Fetch a resource with fetch
} else if ('XMLHttpRequest' in globalThis) {
  // Fetch a resource with XMLHttpRequest
} else {
  // Fetch a resource with some custom AJAX logic
}
```
【练习】请问如下程序打印的结果是什么？

```js
let x = 0;
switch (++x) {
  case 0:
    ++x;
  case 1:
    ++x;
  case 2:
    ++x;
}
console.log(x)
```

- 三元运算符 `?:`，先看一个例子：

```js
let x = 2;
let a = x > 0 ? 1 : -1;
console.log(a);//true
```

【要点】

1）三元运算符会对 `?` 前面的表达式求值，如果为真，就返回`:` 前面表达式的值，否则返回后面的。

2）返回的值可以赋给变量，三元运算符就好像函数一样。

- `try {...} catch(err){...} finally{...}` ，JavaScript 中错误捕获控制流。它的执行规则是：

1）先计算 `try` 里面的语句，如果抛出异常，就直接进入 `catch` 里的语句，如果没有异常，`catch` 就不会执行。

2）如果写了 `finally`，它必定执行。

具体细节请看如下例题：

【例题】请问以下代码的输出结果？

```js
let i = 100;
function foo() {
  bbb: try {
    console.log("try try try");
    return i++;
  } finally {
    console.log("finally");
    break bbb;
  }
  console.log("Will I run?");
  return i;
}
console.log(foo());
```

【解】要把这个题目作对不容易，涉及到的知识点[^1]：

- `try...catch...finally` 的用法
- `break` 的用法

一、`try...catch...finally` 的用法
1、try catch 还有 finally 的基本概念

1）try 块一共有三个关键字 try catch 还有 finally；
2）finally 语句无论 try 和 catch 执行结果如何都会执行；（本题考到的知识点）
3）catch 是捕获到 try 语句块里的错误才会执行；
注意： catch 和 finally 语句都是可选的，但你在使用 try 语句时必须至少使用一个（也就是try必须搭配 catch 或者 finally ）。

2、try catch 还有finally代码块中有return时 的执行情况（本题考到的知识点）

如果 try 语句没有使用 finally，如果 try 中有 return 语句，那么函数 try...catch 语句之外的 return 就不执行了：

```js
function testFinally() {
  var num = 10;
  try {
    return num + 1; //return这里的值11
  } catch (err) {
    console.log(err);
  }
  return num + 5; //无效，没执行到这，try...catch执行之后就跳出了
  //除非try...catch...finaly语句中没有return
}
console.log(testFinally()); //11
```

如果 try 语句后面有 finally，try 中的 return 语句不会跳出函数，因为一定要进入finally 语句，函数 try...finally 语句之外的 return 就不执行了：
```js
function testFinally() {
  var num = 10;
  try {
    return (num += 1); //return右边的语句正常执行，计算得num=11
  } finally {
    //有finally，try中的return不会跳出函数，因为一定要进入finally语句
    return num + 20; //11+20=31
  }
  return num + 5; //无效，没执行到这，try...finally执行之后就跳出了
}
console.log(testFinally()); // 31
```

（可看完后面的break知识点，再来看这个例子） 如果 try 语句后面有 finally，try 中就算有break 用来跳出语句块，也不管用，只要有 finally，不管 try 和 catch 语句中执行什么，一定会进入finally语句：
```js
function testFinally() {
  var num = 10;
  aaa: try {
    break aaa; //有break按理说应该要跳出 try...finally... 这个语句块了
    //但是不管用，因为后面有finally，一定要进入finally语句
  } finally {
    return num + 20; //10+20=30
  }
  return num + 100; //失效，没执行到这
}
console.log(testFinally()); // 30
```
【重点】try...catch...finally 语句中，只要有 finally 存在，不管 try catch 里面有什么，它都必须执行。

二、break 的用法😄

1、break语句用于跳出 switch 语句或者循环语句（ for 、for..in、while、do...while） 

语法：`break;`

1）当 break 语句用于 switch 语句时，会跳出 switch 代码块，终止执行 switch 代码。
2）当 break 语句用于循环语句时，会跳出整个循环，不再执行后续剩余的循环。
3）注意 break 与 continue 的区别：continue会直接进入下一轮循环。

2、break 语句也可用于标签引用，用于跳出标签指向的代码块。（本题考到的知识点）

语法：`break labelName;`

在标签引用中使用 break 语句，用于跳出标签代码块：
```js
var cars = ["BMW", "Volvo", "Saab", "Ford"];
var text = "";
list: {
  //list标签引用
  text += cars[0];
  text += cars[1];
  text += cars[2];
  break list; //在标签引用中使用 break 语句，用于跳出list代码块，不再执行list代码块里剩余的代码
  text += cars[3];
}
console.log(text); //BMWVolvoSaab
```
在标签引用中使用 break 语句，用于跳出嵌套循环：

```js
var text = "";
var i, j;

Loop1: for (i = 0; i < 3; i++) {
  // 第一个循环标签 "Loop1"
  Loop2: for (j = 10; j < 15; j++) {
    // 第二个循环标签 "Loop2"
    if (j == 12) {
      break Loop2; //跳出Loop2代码块
    }
    console.log(i, j);
  }
}
//i=0,j=10
//i=0,j=11
//i=1,j=10
//i=1,j=11
//1=2,j=10
//i=2,j=11
```

【参考答案】你做对了吗？

```js
"try try try"
"finally"
"Will I run?"
101
```

接下来我们学习 JavaScript 中的大小判断。

## 全等运算符 `===`
说明: 严格匹配,不会类型转换,必须要数据类型和值完全一致。

先判断类型，如果类型不一样，直接为 `false`;

1）对于基本数据类型: Number, BigInt, String, Boolean, Null 和 Undefined，两边的值要一致才相等

```js
console.log(null === null)            // true
console.log(undefined === undefined)  // true
// 注意: NaN: 不会等于任何数,包括它自己
console.log(NaN === NaN)              // false
```
【注意】同名的 Symbol 类型不相等，可以当做对象看待。

```js
console.log(Symbol("foo") === Symbol("foo"));//false
```

2）对于复杂数据类型(引用类型): Object,Array,Function 等：只有引用地址相同才是相等的（表示同一个对象）：

```js
let arr1 = [1, 2, 3];
let arr2 = arr1; //浅拷贝
console.log(arr1 === arr2); //true

console.log([] === []); //false
```

## 相等运算符 `==`

非严格匹配: 会类型转换，分前提条件一共有五种情况，接下来的代码以 `x == y` 为示例[^5]。

1）x 和 y 都是 null 或 undefined: 没有隐式类型转换，无条件返回 true：

```js
console.log ( null == undefined );     //true
console.log ( null == null );          //true
console.log ( undefined == undefined );//true
```
2）x 或 y 是 NaN : NaN与任何数字都不等，没有隐式类型转换，无条件返回 false

```js
console.log(NaN == NaN); //false
```
3）x 和 y 是 string，boolean，number

规则：有隐式类型转换，会将不是 number 类型的数据转成 number

```js
console.log ( 1 == true );    //true    (1) 1 == Number(true)
console.log ( 1 == "true" );  //false   (1) 1 == Number('true')
console.log ( 1 == !"true" ); //false  (1) 1 == !Boolean('true')  (2) 1 == !true  (3) 1 == false  (4)1 == Number(false)
console.log ( 0 == !"true" ); //true
console.log (true == 'true');  // false
```
4）x 或 y （其中一个）是复杂数据类型 : 会先获取复杂数据类型的原始值之后再左比较复杂数据类型的原始值： 
先调用`valueOf`方法，然后调用 `toString` 方法
valueOf: 一般默认返回自身
数组的 toString：默认会调用 join 方法拼接每个元素并且返回拼接后的字符串。

```js
console.log ( [].toString () );//空字符串
console.log ( {}.toString () );//[object Object]
```
注意:  空数组的toString()方法会得到空字符串，而空对象的toString()方法会得到字符串[object Object] （注意第一个小写o，第二个大写O哟）

```js
console.log ( [ 1, 2, 3 ].valueOf().toString());//‘1，2，3’
console.log ( [ 1, 2, 3 ] == "1,2,3" );//true  (1)[1,2,3].toString() == '1,2,3'  (2)'1,2,3' == '1,2,3'
console.log({} == '[object Object]');//true
```
5）x 和 y 都是复杂数据类型 :
规则：只比较地址，如果地址一致则返回 true，否则返回 false

```js
var arr1 = [10,20,30];
var arr2 = [10,20,30];
var arr3 = arr1;//将arr1的地址拷贝给arr3
console.log ( arr1 == arr2 );//虽然arr1与arr2中的数据是一样，但是它们两个不同的地址
console.log ( arr3 == arr1 );//true  两者地址是一样
console.log ( [] == [] );//false
console.log ( {} == {} );//false
```
![week-equal](img/week-equal.png)

【例题】判断如下结果：

```js
console.log([] == 0); //true 
// 分析:(1) [].valueOf().toString() == 0  (2) Number('') == 0  (3) false == 0  (4) 0 == 0
console.log(![] == 0); //true
// 分析: 逻辑非优先级高于关系运算符 ![] = false (空数组转布尔值得到true)

console.log({} == {}); //false
console.log([] == []); //false
// 分析：第五种情况，对象地址不一样

console.log([] == ![]); //true
// 分析：(1) [] == !Boolean([])   (2) [] == !true  (3)[] == false  (4) [].toString() == false  (5)'' == false   (6)Number('0') == Number(false)

console.log({} == !{}); //false
// 分析：(1){} == !{} (2){} == !true  (3){} == false  (4){}.toString() == false  (5)'[object Object]' == false  (6)Number('[object Object]') == false
```

【例题】请问变量 `a` 初始化为何值，能使其正确打印 1？

```js
var a = ???;
if (a == 1 && a == 2 && a == 3) {
  console.log(1);
}
```
参考答案：
```js
var a = {
  i: 1, // 声明一个属性i
  valueOf: function () {
    return a.i++; // 每调用一次，让对象a的i属性返回并且自增一次
  },
};

if (a == 1 && a == 2 && a == 3) {
  // 每一次运算时都会调用一次a的valueOf()方法
  console.log('1');
}
```

在 JavaScript 中，要尽量使用全等运算符（“===”），因为全等运算符在比较时都不会进行类型的转化，相对而言速度也会更快，因为在使用弱相等 `== ` 时，类型不一样会进行类型转换 。

图标版比较请阅读参考文章[^2]。

更多细节请参考文章[^3]。

## 大小于系列

字符串内部与数值类型内部可以自然的比较大小，除此之外其他类型就需要类型转换。

- 对象通过调用  [`@@toPrimitive()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/toPrimitive) 转化为基本类型 ([`valueOf()` ](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/valueOf)然后 [`toString()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/toString) )。
- 如果两个操作数都是字符串类型，按字符串比较。
- 其他的都转换成数值型，然后比较。
- NaN直接返回 false。

示例：

```js
console.log("5" < 3);          // false Number("5")===5
console.log("5" < 3);          // false Number("hello") is NaN
console.log(false < true);     // true  0 < 1
console.log(null < 1);         // true  0<1
console.log(undefined < 3);    // false NaN<3
console.log([] <= 0);		   // true []->""->0
```

【练习】请直接写出结果[^6]：

```js
[] == [];
[] === [];
{} == {};
{} === {};
{}+[];
[]+{};
```
更多面试题请阅读参考文章[^7]。

## 参考文章

[^1]: 笨笨只会灭火. [单选题](https://www.nowcoder.com/questionTerminal/a894fca888c7411894a448c4a08d60c0). 牛客. 有改动. 

[^2]: dorey. [JavaScript equality testing](https://dorey.github.io/JavaScript-Equality-Table/). blog. 
[^3]: [Equality (==)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Equality). MDN. 

[^4]: [Less than (<)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Less_than). MDN. 

[^5]:码飞CC. [JS比较运算符的匹配规则以及if()条件的判断结果](https://blog.csdn.net/cc18868876837/article/details/88867982). CSDN. 

[^6]:justjavac. [详解一下 javascript 中的比较](https://segmentfault.com/q/1010000000305997). segmentfault.com. 

[^7]: CUGGZ. [「2021」高频前端面试题汇总之JavaScript篇（上）](https://juejin.cn/post/6940945178899251230#heading-15). 稀土掘金.



> ♥ 我是前端工程师：你的甜心森。非常感谢大家的点赞与关注，欢迎大家参与讨论或协作。
>
> ★ 本文[开源](https://github.com/xiayulu/FrontEndCultivation)，采用 [CC BY-SA 4.0 协议](http://creativecommons.org/licenses/by-sa/4.0/)，转载请注明出处：[前端工程师的自我修养](https://github.com/xiayulu/FrontEndCultivation). GitHub.com@xiayulu.
>
> ★ 创作合作或招聘信息请发私信或邮件：zuiaiqiansen@163.com，注明主题：创作合作或**招聘前端工程师**。

