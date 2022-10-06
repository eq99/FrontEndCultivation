# JavaScript 条件判断详解

## 两种等号的使用策略

在 JavaScript 中，要尽量使用全等运算符（“===”），因为全等运算符在比较时都不会进行类型的转化，相对而言速度也会更快。那么什么时候使用相等运算符（“==”）呢？以下两种情况供参考。

1）用于判断对象的属性是否存在:

```js
let obj = {};
if( obj.a == null ) {
    //这里相对于：obj.a === null || obj.a === undefined 的简写形式，jquery源码的推荐写法
}
```
2）用于判断函数的参数是否存在：
```js
function fn( a, b ) {
    if( a == null ) {
        //这里也相当于 a === null || a === undefined 的简写
    }
}
```
总结: 一般情况下我们尽量使用 `===` 来精确判断，在判断对象属性和函数参数是否存在时可以使用 `==`。

## 全等运算符 `===`
说明: 严格匹配,不会类型转换,必须要数据类型和值完全一致。

先判断类型,如果类型不一样，直接为 `false`;

1）对于基本数据类型(值类型): Number,String,Boolean,Null 和 Undefined: 两边的值要一致才相等
```js
console.log(null === null)            // true
console.log(undefined === undefined)  // true
// 注意: NaN: 不会等于任何数,包括它自己
console.log(NaN === NaN)              // false 
```
2）对于复杂数据类型(引用类型): Object,Array,Function等：两边的引用地址如果一致的话,是相等的：
```js
let arr1 = [1,2,3];
let arr2 = arr1;
console.log(arr1 === arr2)   // true
```

## 相等运算符 `==`
非严格匹配: 会类型转换，但是有前提条件一共有五种情况
(接下来的代码以 `x == y` 为示例)

1）x 和 y 都是 null 或 undefined: 没有隐式类型转换，无条件返回 true

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
4）x 或 y （一个）是复杂数据类型 : 会先获取复杂数据类型的原始值之后再左比较复杂数据类型的原始值： 
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
💯【面试题】判断如下结果：

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

💯【变态面试题】请问变量 `a` 初始化为何值，能使其正确打印 1？

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

### switch

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



## 三目运算符

```js
let y = false;
let x = y === true ? 'haha' : 'hoho';
console.log(x);
```

# 参考文章

- [JS比较运算符的匹配规则以及if()条件的判断结果](https://blog.csdn.net/cc18868876837/article/details/88867982)
- [详解一下 javascript 中的比较](https://segmentfault.com/q/1010000000305997)
- [「2021」高频前端面试题汇总之JavaScript篇（上）](https://juejin.cn/post/6940945178899251230#heading-15)