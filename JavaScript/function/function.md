# JavaScript 函数快速入门

函数是程序设计的基本单元，是程序员的锤子与镰刀，如果没有函数，程序的世界将漆黑一片。

## 函数的定义

JavaScript 函数主要有两种的定义方式：

### 具名函数

```html
<input type="text" id="nickname" placeholder="请输入昵称">
<button onclick="alert('你的昵称：' + getNickname())">查看昵称</button>
<script>
  function getNickname() {
    const inputEle = document.getElementById("nickname");
    return inputEle.value || "你的甜心森";
  }
</script>
```

【要点解读】

1）具名函数的声明方式是直接在 `function ` 后面接函数名，例如 `getName`

2） 函数的主要作用是封装复杂的逻辑，函数调用者不需要关心内部细节，拿来即用就行。上述函数的例子封装了获取昵称的逻辑，在需要使用昵称的地方调用 `getNickname()` 就行。
3） JS 函数的命名采用小驼峰风格，并且以动宾短语为主：`function doSomething(){...}`。

### 匿名函数

```html
<input type="text" id="nickname" placeholder="请输入昵称">
<button id="display">查看昵称</button>
<script>
  function getNickname() {
    const inputEle = document.getElementById("nickname");
    return inputEle.value || "你的甜心森";
  }
  document.getElementById("display").onclick = function() {
    alert('你的昵称：' + getNickname());
  }
</script>
```

【要点解读】

1）匿名函数的特点是 `function` 没有函数名，匿名函数没有名字，看似无用，但是使用场景其实很多：

- 用作对象属性

```js
const user = {
  firstName: "Java",
  lastName: "Script",
  fullName: function () {
    return this.firstName + this.lastName;
  }
};

document.write(user.fullName());
```

- 用作回调函数：

```js
<input type="text" id="nickname" placeholder="请输入昵称">
<button id="display">查看昵称</button>
<script>
  function getNickname() {
    const inputEle = document.getElementById("nickname");
    return inputEle.value || "你的甜心森";
  }
  document.getElementById("display").addEventListener("click", function() {
    alert('你的昵称：' + getNickname());
  });
</script>
```

- 当然也可以通过赋值给变量，然后像普通函数那样使用：

```js
<input type="text" id="nickname" placeholder="请输入昵称">
<button id="display">查看昵称</button>
<script>
  const getNickname = function() {
    const inputEle = document.getElementById("nickname");
    return inputEle.value || "你的甜心森";
  }
  document.getElementById("display").addEventListener("click", function() {
    alert('你的昵称：' + getNickname());
  });
</script>
```

【❓ 思考】如果同时使用两种方式会发生什么？

```js
const f = function g() {
  return 23;
};

console.log(typeof g, typeof f);
console.log(f()); //23
console.log(g()); //ReferenceError: g is not defined
```

上述例子表明：如果把具名函数赋值给一个变量，那么具名函数失效，函数当做匿名函数处理。

## 变量提升

函数的的声明存在变量提升，但是匿名函数必须先声明后使用。

【例】正确：

```js
fun();
function fun() {
  console.log(1);
}
```

【例】错误：

```js
fn(); //TypeError: fn is not a function

var fn = function () {
  console.log(1);
}
```

## 函数的参数

函数参数(parameters, arguments)放在函数声明的括号里，如果有多个参数，使用逗号隔开即可：

```js
function validateEmail(email) {
  return email.includes('@');
}

const email = 'hello@world.com';
console.log(validateEmail(email));
```

【🔑技巧】声明函数时括号里的参数叫形参，调用函数时括号里的参数叫实参。

### 可选参数

调用函数时，实参的个数可以小于形参的个数：

```js
function haveLunch(main, meat, vegetable, fruit) {
  console.log('今天中午我吃了：', main, meat);
}

let main = '饺子';
let meat = '牛肉';
let vegetable = '菠菜';
let fruit = '苹果';
console.log(haveLunch(main, meat));
```

如果实参个数大于形参个数，函数不会报错，多余的实参会被丢弃：

```js
function haveLunch(main, meat) {
  console.log('今天中午我吃了：', main, meat);
}

let main = '饺子';
let meat = '牛肉';
let vegetable = '菠菜';
let fruit = '苹果';
console.log(haveLunch(main, meat, vegetable));
```

JS 不能像 Python 等语言语言那样指定相应参数的值，参数只与顺序有关：

```js
function haveLunch(main, meat) {
  console.log('今天中午我吃了：', '主食:', main, '肉食:', meat);
}

let main = '饺子';
let meat = '牛肉';
let vegetable = '菠菜';
let fruit = '苹果';
console.log(haveLunch(meat = vegetable, main = fruit));
// 今天中午我吃了：主食: 菠菜 肉食: 苹果
```

【🔑技巧】**一个萝卜一个坑，形参挖坑，实参填坑**。

### 默认参数

```js
// 生成 n 位验证码，默认四位
function genCaptcha(n = 4) {
  let result = '';
  for (let i = 0; i < n; i++) {
    result += Math.floor(Math.random() * 10);
  }
  return result;
}

console.log(genCaptcha());
console.log(genCaptcha(6));
```

可以通过给形参赋值指定它的默认值。

### arguments 对象[^1]

每一个函数都有一个隐形的对象：`arguments`，代表传进来来的参数。他有两个属性：`length` 和 迭代器，类似数组。

```js
function haveLunch(main, meat) {
  for (let i = 0; i < arguments.length; i++) {
    console.log(arguments[i]);
  }
}

let main = '饺子';
let meat = '牛肉';
let vegetable = '菠菜';
let fruit = '苹果';
haveLunch(main, meat, vegetable, fruit);
```

形参与 arguments 对象指向是一样的：

```js
function haveLunch(main, meat) {
  console.log(main === arguments[0]); // true
  console.log(meat === arguments[1]); // true
}

let main = '饺子';
let meat = '牛肉';
let vegetable = '菠菜';
let fruit = '苹果';
haveLunch([], {}, vegetable, fruit);
```

有了 arguments 对象，函数可以拥有不确定个数参数：

```js
function sum() {
  let result = 0;
  for (let i = 0; i < arguments.length; i++) {
    result += arguments[i];
  }
  return result;
}

console.log(sum(1, 2, 3, 4));//10
```

## 函数的返回值

函数使用 `return` 关键字返回值，`return` 语句特点：

- 返回一个表达式的值，例如变量，字面量，三目运算符等的值。

- 只能返回一个值，如果有多个值，请使用数组或对象。

```js
function getToken() {
  let accessToken = "accss";
  let freshToken = "fresh";

  return [accessToken, freshToken];
}

const [accessToken, freshToken] = getToken();
console.log(accessToken, freshToken);//"accss" "fresh"
```

- `return` 语句能够提前结束循环，与 `switch` 语句。

```js
function hasElement(arr, ele) {
  for (let val of arr) {
    if (val === ele) {
      return true;
    }
  }
  return false;
}

console.log(hasElement([1, 2, 3, 4, 5], 3));//true
```

## 参考文章

[^1]:MDN. [The arguments object](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/arguments)

