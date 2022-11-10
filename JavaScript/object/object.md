# JavaScript 对象详解

JavaScript 共有八种数据类型，其中有七种基础数据类型（primitive types），还有今天要学习的对象数据类型（Objects）。对象数据类型与其他编程语言里的类，结构体类似，是一种复杂数据类型。

## 初识对象

### 对象的声明

使用构造函数:

```js
let user = new Object();
let car = Object();
```

不像基础数据类型，如果直接调用构造函数创建的是基本数据类型，存储在栈中，如果使用 `new` 那么创建的是相应的对象，存储在堆中。对象数据类型创建的结果都存储在堆中，返回的是对象的指针，或者说**引用**。

使用对象字面量：

```js
const user = {};

const car = {
  name: "DongFeng",
  getName: function () {
    return this.name;
  }
};

console.log(car.getName());
```

【**技巧**】对象的声明可以使用快捷方式：

```js
const addr = 'HuNan';
const age = 18;

let user1 = {
  addr: addr,
  age: age,
};

// 更简单
let user2 = {
  addr,
  age,
};
```

### 对象的属性

对象内部的键值对： `key: value`  叫做对象的属性(properties)，在其他语言里叫做成员变量或者结构体字段，冒号前面的叫属性名，后面的叫属性值，如果属性值得类型是函数，该属性可以称作对象的一个**方法(method)**。属性名只能是两种类型：字符串和 Symbol。属性值可以是任何八种数据类型。

【**知识点**】访问属性。可以通过点运算符 `.` 或中括号 `[]` 访问属性，请看示例：

```js
let id = Symbol("id");
let user = {
  "age of user": 18,
  _name: "Michael",
  $: 23,
  [id]: 42,
  2: 2333
};

let dollar = "$";

console.log(user._name); //"Michael" 使用点运算符
console.log(user["_name"]); //"Michael" 也可以使用 []
console.log(user[id]); //42 Symbol 类型
console.log(user.id); //undefined 不能用点运算符访问 Symbol 属性
console.log(user["age of user"]); //18 只要是字符串就能做属性名
console.log(user[1 + 1]); //2333 额额，算术表达式也行？
console.log(user[dollar]); //23 []可以使用变量访问属性
console.log(user.dollar); //undefined .操作符不能使用变量
```

【**思考**】`.` 操作符与 中括号`[]` 有什么区别呢？？

点操作只能只能用来获取**标识符形式**（不是变量）的属性，可以是纯数字或合法的变量名（`/^[a-zA-z_$][a-zA-Z0-9$_]*$/`），而中括号内部可以是一个表达式，该表达式的求值结果能转换成字符串或 Symbol 类型就行[^1]。

【**知识点**】修改或添加属性。可以通过赋值操作修改或添加属性：

```js
let user = {
  age: 18,
};
console.log(user); //{age: 18}

user.name = 'Jackson';
user['age'] += 1;
console.log(user); //(2) {name: "Jackson", age: 19}
```

【**提示**】声明对象的时候也可是使用 `[]` 添加表达式。

【**知识点**】属性存在性测试。

你可以使用 `undefinded` 测试：

```js
let user = {};

console.log(user.noSuchProperty === undefined);//true
```

专门的属性存在性测试算符 `in`:

```js
let user = {
  'age of user': 18,
};

let age = 'age of user';

console.log('age of user' in user); //true
console.log(age in user); //true
console.log('baga' in user); //false

//删掉之后
delete user[age];
console.log(age in user); //false

// undefined 不够专业
user.subject = undefined;
console.log(user.subject === undefined); //true
console.log('subject' in user); //true
```

【**知识点**】遍历属性名。可以使用 `for..in` 语法遍历对象的属性：

```js
let id = Symbol();

let user = {
  'age of user': 18,
  name: 'Jack',
  [id]: 'Can you see me?',
  [1 + 1]: '1+1=2',
};

for (let key in user) {
  console.log(key);
}
//2
//age of user
//name
```

【**技巧**】因为 `in` 用来测试属性存在性，所以也用 `in` 遍历属性。

> 本节参考文章[^2]

【知识点】解构语法`...`。解构语法用在等号右边会把对象展开，用在左边会搜集余下的属性：

```js
let user = {
  name: 'Jack',
  bio: 'Sleep, sleep',
  food: ['Chicken', 'Beef', 'apple'],
};

// 把 user 里的属性取出来，放到 user 2 里面
let user2 = {
  ...user,
};

console.log(user, user2);

// 把 user 的 name 单独拿出来，其他的放到 others 里面
let { name, ...others } = user;
console.log(name, others);
```

## 对象与引用

JS 创建对象时，返回的是对象的引用，因此对象变量进行赋值操作或函数传递参数时，传递的是引用：

```js
let user = {
  name: 'Jack',
};

let userCopy = user; // 赋值操作只是增加对象的引用
userCopy.name = 'JackMan'; // 修改对象的任何一个引用会反映到对象本体上

console.log(user, userCopy); // {name: "JackMan"} {name: "JackMan"}  <---都改变了
console.log(user === userCopy); // 其实是同一个对象的指针
```

【重点】这种传递引用的方式叫做**浅拷贝**。

对比基本数据类型，他们进行赋值操作或函数传递参数时，会创建新对象，两个变量毫不相关：

```js
let name1 = 'Jack';
let name2 = name1;

name2 = 'JackMan';
console.log(name1, name2); //Jack JackMan
```

【重点】基本类型赋值操作创建的数据类型相互独立，这种方式也叫**深拷贝**。

【技巧】赋值和传参的小口诀：基本类型深拷贝，对象类型浅拷贝。

【知识点】深拷贝对象。如果要对对象进行深拷贝，有如下几种方法：

- 使用 `structuredClone`

```js
let sheep = {
  name: "dolly",
  foods: ["apple", "pear"]
};

let cloned = structuredClone(sheep);

console.log(sheep.name === cloned.name, sheep.foods === cloned.foods);
// true false
```

- 使用 `Json.parse(Josn.stringify(obj))`，相当于把对象碾碎成字节然后重构，涅槃重生了属于是。
- 使用第三方库函数 [_.cloneDeep(obj)](https://lodash.com/docs#cloneDeep)。

## 对象常用的方法

### [`Object.assign()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/assign)

把从一个或多个源对象的可迭代属性复制到目标对象中。

使用方法：`Object.assign(target, ...sources)`，示例：

```js
const target = { a: 1, b: 2 };
const source1 = { b: 4, c: 5 };
const source2 = { 'i am cat': 'cat' };

const returnedTarget = Object.assign(target, source1, source2);

console.log(target);
// {a: 1, b: 4, c: 5, i am cat: "cat"}

console.log(returnedTarget === target);//true
```

【注意】`Object.assign()` 对对象的复制是浅拷贝。

【注意】解构语法`...` 与 `Object.assign()` 类似，都可以轻松的拷贝属性，并且对象都是浅拷贝。

### [`Object.entries()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/entries)

返回一个数组，每个元素包含对象自己的可迭代字符串属性键值对`[key, value]`，使用方法 `Object.entries(obj)`。

【示例】

```js
const obj = {
  a: 'somestring',
  b: 42,
};

for (const [key, value] of Object.entries(obj)) {
  console.log(`${key}: ${value}`);
}
// "a: somestring"
// "b: 42"

// 或者：函数式迭代对象的属性
Object.entries(obj).forEach(([key, value]) => {
  console.log(`${key} ${value}`);
});
```

这与 `for..in` 语法类似

【示例】把一个对象转化成字典：

```js
const obj = {foo: "bar", baz: 42};
const mp = new Map(Object.entries(obj));
console.log(mp); // Map(2) {"foo" => "bar", "baz" => 42}
```

更多对象的操作方法请查阅参考文章 3[^3]

【例题】请设计一个函数 `flaten`，将对象扁平化，假设对象的值只涉及到普通对象和数字。

```js
// 原对象
const obj = {
  a: 123,
  b: {
    c: 1145,
    d: {
      e: 666
    }
  }
};

// 扁平化之后
let result  = {
  "a": 123,
  "b.c": 1145,
  "b.d.e": 666
};
```

参考答案：

```js
function flaten(obj) {
  const tmp = {};
  Object.entries(obj).forEach(([key, val]) => {
    if (typeof val === "object") {
      const inner = flaten(val);
      Object.entries(inner).forEach(([key2, val2]) => {
        tmp[key + "." + key2] = val2;
      });
    } else {
      tmp[key] = val;
    }
  });

  return tmp;
}
```

## 参考文章

[^1]: [JavaScript property access: dot notation vs. brackets?](https://stackoverflow.com/questions/4968406/javascript-property-access-dot-notation-vs-brackets). stackOverflow.
[^2]:[Objects](https://javascript.info/object). JavaScript.info. 
[^3]:[Object](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object). MDN.



> ♥ 我是前端工程师：你的甜心森。非常感谢大家的点赞与关注，欢迎大家参与讨论或协作。
>
> ★ 本文[开源](https://github.com/xiayulu/FrontEndCultivation)，采用 [CC BY-SA 4.0 协议](http://creativecommons.org/licenses/by-sa/4.0/)，转载请注明出处：[前端工程师的自我修养](https://github.com/xiayulu/FrontEndCultivation). GitHub.com@xiayulu.
>
> ★ 创作合作或招聘信息请发私信或邮件：zuiaiqiansen@163.com，注明主题：创作合作或**招聘前端工程师**。

