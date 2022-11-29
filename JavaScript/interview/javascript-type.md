# JavaScript 面试题通关一：类型

## JavaScript 有哪些数据类型，它们的区别？

JavaScript 共有八种数据类型：Boolean, Number, String, Null, Undefined, BigInt, Symbol, Object。

1）按存储方式来看，其中前面七种是基本数据类型，存储在栈上；Object 是引用数据类型，存储在堆上，在栈中存储指针。

2）按 ES 标准来看，BigInt 与 Symbol 是 ES6 新增的数据类型。BigInt 用来存储大整数，Symbol 用于解决全局属性名冲突的问题。

🔑 记忆方法：7+1。

## JavaScript 数据类型检查的方式有哪些?

JavaScript 数据类型检查的方式有 4 种：

第一种：使用 typeof 检查：

```js
console.log(typeof 2);               // number
console.log(typeof true);            // boolean
console.log(typeof 'str');           // string
console.log(typeof undefined);       // undefined
console.log(typeof null);            // object
console.log(typeof 1n);              // bigint
console.log(typeof Symbol());        // symbol
console.log(typeof {});              // object
console.log(typeof []);              // object    
console.log(typeof function(){});    // function
```

🔑 记住：typeof 的结果是**小写字符串**，只有**八**种结果。

第二种：使用 `Object.prototype.toString.call()` 检查：

```js
function getType(obj) {
  return Object.prototype.toString.call(obj);
}

function Person() {}

console.log(getType(2)); // [object Number]
console.log(getType(true)); //[object Boolean]
console.log(getType("str")); //[object String]
console.log(getType([])); //[object Array]
console.log(getType(function () {})); //[object Function]
console.log(getType({})); //[object Object]
console.log(getType(undefined)); //[object Undefined]
console.log(getType(null)); //[object Null]
console.log(getType(1n)); //[object BigInt]
console.log(getType(Symbol())); //[object Symbol]
console.log(getType(new Person())); //[object Object]
```

💡 提示： `Object.prototype.toString.call()` 能更加细致的检查类型，它的返回值是**字符串**，第二个单词**首字母大写**。

第三种：使用 **instanceof** 检查原型:

```js
console.log(2 instanceof Number);                    // false
console.log(true instanceof Boolean);                // false 
console.log('str' instanceof String);                // false 
 
console.log([] instanceof Array);                    // true
console.log(function(){} instanceof Function);       // true
console.log({} instanceof Object);                   // true
```

🔑 记住：**instanceof** 只能用于**对象**，返回**布尔值**。

第四种：通过构造器判断类型：

```js
console.log((2).constructor === Number); // true
console.log((true).constructor === Boolean); // true
console.log(('str').constructor === String); // true
console.log(([]).constructor === Array); // true
console.log((function() {}).constructor === Function); // true
console.log(({}).constructor === Object); // true
```

⚠️ 构造器可能会改变：

```js
function Fn() {}

Fn.prototype = new Array();

const f = new Fn();

console.log(f.constructor === Fn); // false
console.log(f.constructor === Array); // true
console.log(Array.isArray(f)); // false
```

## 判断数组类型的方式有哪些？

数组是对象，不能用 typeof 检查类型，主要有三个思路：

一、通过原型做判断：

```js
const a = [];
const b = {};
console.log(a instanceof Array, b instanceof Array); // true false
console.log(Array.prototype.isPrototypeOf(a), Array.prototype.isPrototypeOf(b));
// true false
console.log(a.__proto__ === Array.prototype, b.__proto__ === Array.prototype); 
// true false
```

二、通过 Object.prototype.toString.call() 判断：

```js
const a = [];
const b = {};
function getType(obj) {
  return Object.prototype.toString.call(obj);
}
console.log(getType(a), getType(b));  //[object Array] [object Object]
```

三、使用 ES6 的 Array.isArrray()：

```js
const a = [];
const b = {};
console.log(Array.isArray(a), Array.isArray(b));  //true, false
```

## null 和 undefined 区别

先说共性：null 和 undefined 都是基础类型，他们的值分别为 null, undefined。比较的时候，两个等号返回 true，三个等号返回 false。

区别一、从变量初始化角度：

undefinded 代表未定义，变量声明了但未初始化是未定义。

null 代表空对象，一般用作某些对象变量的初始化值。

区别二、从语言层面来讲：

undefinded 是 window 下的一个属性，不是保留字，甚至可以用作变量名，安全获取方式应该是：`void 0`

null 是 JavaScript 语言关键字，使用 typeof 检查会返回 'object'

🔑 技巧：说区别的题可以先解释每个条目的含义，然后说共性，最后说区别。

## typeof null 的结果是什么，为什么？

typeof null 的结果是 "object"。

早期的 JavaScript 数据都采用 32 位存储，并且使用最低的1~3位作为类型标记：

```js
000: object   - 当前存储的数据指向一个对象。
  1: int      - 当前存储的数据是一个 31 位的有符号整数。
010: double   - 当前存储的数据指向一个双精度的浮点数。
100: string   - 当前存储的数据指向一个字符串。
110: boolean  - 当前存储的数据是布尔值。
```

其中 undefinded 的位模式是全1，而 null 的是全 0 （表示空指针）。

因此 null 的低三位与 Object 低三位相同，所以它的类型标志为 Object。

## 为什么0.1+0.2 ! == 0.3，如何让其相等 ？

```js
let a = 0.1;
let b = 0.2;
console.log(a + b);//0.30000000000000004
```

因为 JavaScript 的 Number 是 64 位浮点数，小数运算可能存在精度丢失问题，所以它们的值不相等。

让它们相等有两种思路：

思路一：放大 10 倍之后相加再缩小 10 倍。

思路二：封装浮点数相等的函数：

```js
function feq(a, b) {
  return Math.abs(a - b) < Number.EPSILON;
}

console.log(feq(0.1 + 0.2, 0.3));//true
```

补充：

```js
  0 01111111011  1001100110011001100110011001100110011001100110011010     // 0.1 位模式
```
计算机内部计算过程：
```js
  0 01111111100  0.1100110011001100110011001100110011001100110011001101  // 转为大指数
+ 0 01111111100  1.1001100110011001100110011001100110011001100110011010
= 0 01111111100 10.0110011001100110011001100110011001100110011001100111
= 0 01111111101  1.00110011001100110011001100110011001100110011001100111 // 进一
= 0 01111111101  1.0011001100110011001100110011001100110011001100110100
  0 01111111101  1.0011001100110011001100110011001100110011001100110011  // 0.3
```

## 其他值到字符串的转换规则？

JavaScript 共有八种类型，其他七种类型转化为字符串类型的规则是：

1）boolean 类型：直接转化为 "true" 和 "false"

2）number 类型：调用静态函数 `toString(10)`，即转化为相应的十进制格式。

3）null 与 undefined : 直接转化为 'null' 与 'undefined'。

4）bigint 类型：调用静态函数 `toString(10)`，即转化为相应的十进制格式。

5）symbol 类型：隐式类型转换会报错，只允许使用 显示转换。

6）object：转化时会调用 [`@@toPrimitive`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/toPrimitive) 指定的方法转化为基本类型之后再进行转换，具体过程是先调用 `toString()` 方法，如果不能得到基本类型就再调用 `valueOf()`方法，如果还不能转为基本类型就抛出 TypeError。

## 其他值到数字值的转换规则？

其他七种类型转换为数值的规则：

1）boolean: true 为 1，false 为 0.

2）string：调用 Number(xxx) 转换，空字符串为0，转化失败是 NaN

3）null：0

4）undefined: NaN

5）bigint：直接转成相应的 number

6）symbol：不能直接转化，报错

7）object: 转化时会调用 [`@@toPrimitive`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/toPrimitive) 指定的方法转化为基本类型之后再进行转换，具体过程是先调用 `valueOf()` 方法，如果不能得到基本类型就再调用 `toString()`方法，如果还不能转为基本类型就抛出 TypeError。

## 其他值到布尔类型的值的转换规则？

只需要记住如下8个值转为布尔值为 false：0, +0, -0, 0n, NaN, "", undefinded, null，其他一律为 true，[点击查看原因](https://262.ecma-international.org/#sec-toboolean)。

⚠️ 如何理解如下类型转换：

```js
console.log([] == false); //true
console.log(Boolean([])); //true
console.log(Number([]));  //0
```

## 什么是 JavaScript 中的包装类型？

JavaScript 中的基本类型没有属性与方法，使用某些属性或方法时会转成相应的对象类型：

```js
const a = "abc";
// 相当于 const a = new String("abc");
a.length; // 3
a.toUpperCase(); // "ABC"
```

将基本类型转成相应包装类型方法：

```js
const a = Object('abc');
```

将包装类型换成基本类型：

```js
const a = Object('abc').valueOf();
```

## `+` 操作符类型转换规则？

主要有四个规则：

1）转化为数值类型：boolean, number, undefined, null, bigint 之间运算。

2）有字符串参与：boolean, number, undefined, null, bigint, string 之间的运算转化为字符串拼接。

3）有 object 参与，转化为基本类型之后按上述两条规则进行。

4）有 symbol 参与直接报错。

## == 操作符的隐式类型转换规则是什么？

分为四个步骤：

1）判断类型是否相同，如果相同则不需要进行类型转换。

2）如果是基本类型之间的比较 (string, number, undefined, null, boolean, bigint, symbol) ：

a）如果是 undefinded == null，返回 true；

b）有 symbol：直接报错；

b）转成 number 再作比较。

3）如果包含对象：

a）只包含对象：不存在类型转换，只比较指针是否相同

b）把对象转成相应的的基本类型，再按照规则

 2）作比较。

## `<`  等的比较运算符转化规则

主要分为两个步骤：

1）基本类型之间：

a）只包含字符串：按字符串比较规则。

b）包含 symbol：直接报错

c）转化为 number 作比较。

2）包含对象，把对象转换成基本类型再按相应规则进行比较。

🔑 技巧：把类型分为两类：基本类型与对象类型，只需记住特殊情况（string, symbol），这样思路就比较清晰。

## 总结与说明

面试题整理自网络，如有疏漏与不足请在评论区留言。

建议照着题目，不看解析自己整理思路，回答一遍。



> ♥ 我是前端工程师：你的甜心森。非常感谢大家的点赞与关注，欢迎大家参与讨论或协作。
>
> ★ 本文[开源](https://github.com/xiayulu/FrontEndCultivation)，采用 [CC BY-SA 4.0 协议](http://creativecommons.org/licenses/by-sa/4.0/)，转载请注明出处：[前端工程师的自我修养](https://github.com/xiayulu/FrontEndCultivation). GitHub.com@xiayulu.
>
> ★ 创作合作或招聘信息请发私信或邮件：zuiaiqiansen@163.com，注明主题：创作合作或**招聘前端工程师**。

