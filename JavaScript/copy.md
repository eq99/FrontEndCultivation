# 浅拷贝与深拷贝

浅拷贝只复制指向某个对象的指针，而不复制对象本身，新旧对象还是共享同一块内存。但深拷贝会另外创造一个一模一样的对象，新对象跟原对象不共享内存，修改新对象不会改到原对象。

基础数据类型（Undefined、Null、Boolean、Number、String、Symbol、BigInt）的赋值操作都是深拷贝。而数据对象（Object, Array）的赋值操作是浅拷贝。

请看一个例子：

```js
function copy(src) {
  const dst = {};
  for (let prop in src) {
    if (src.hasOwnProperty(prop)) {
      dst[prop] = src[prop];
    }
  }
  return dst;
}

let obj1 = {
  name: "obj1",
  number: [1, [2, 3], [4, 5]]
};

let obj2 = copy(obj1);
console.log("obj1: before", JSON.stringify( obj1));
console.log("obj2: before", JSON.stringify( obj2));

obj2.name = "obj2";
obj2.number[1] = ["二", "三"];

console.log("obj1: after", JSON.stringify( obj1));
console.log("obj2: after", JSON.stringify( obj2));
```

打印结果：

```shell
"obj1: before" "{'name':'obj1','number':[1,[2,3],[4,5]]}"
"obj2: before" "{'name':'obj1','number':[1,[2,3],[4,5]]}"
"obj1: after" "{'name':'obj1','number':[1,['二','三'],[4,5]]}"
"obj2: after" "{'name':'obj2','number':[1,['二','三'],[4,5]]}"
```

可以看到新对象(obj2) 对 `name` 属性的修改只影响到了自己，但是 `number` 属性是两个不同的指针对象，但他们指向的数组依然是同一个，修改新对象(obj2.number)也会影响原对象(obj1.number)。

## 对象拷贝函数 `assign()` 

```js
let obj1 = {
  a: "aaa",
  b: {
    c: "ccc"
  }
};

let obj2 = { a: "aa", d: "dd" };
let obj3 = Object.assign(obj2, obj1);
console.log(obj3 === obj2); // true, 其实返回的是 obj2
console.log(obj3 === obj1); // false, obj1 与 obj2 不是同一个对象
console.log(obj3.a === obj1.a); // 简单数据：比较类型和值
console.log(obj3.b === obj1.b); // 对象：比较地址，同一个地址
console.log("obj1:", JSON.stringify(obj1));
// "obj1:" "{'a':'aaa','b':{'c':'ccc'}}"
console.log("obj3:", JSON.stringify(obj3));
// "obj3:" "{'a':'aaa','d':'dd','b':{'c':'ccc'}}"
```

大家可以看示例的结果和注释，`assign()` 函数会把第二个参数的属性赋值给到第一个对象上，这和第一个示例的拷贝有点类似。

## 展开运算符 `...` 

展开运算符 `...` 会拷贝所有的属性值到新的对象中，如果属性值是对象的话，拷贝的是地址。

```js
let a = {
    a: 'aaa',
    b: {
        a: 'aaa'
    }
}
let b = {
    ...a
};
console.log(b === a); // false
console.log(b.b === a.b); // true
```

## 如何进行深拷贝

我们发现只要是对象，就会进行浅拷贝，因此我们想到的是递归的对对象进行深拷贝。

```js
function deepClone(obj) {
  //可传入对象或数组
  if (obj == null) return obj; // null 或 undefined 直接返回

  // 区分对象是普通对象还是数组
  if (Object.prototype.toString.call(obj) === "[object Object]") {
    let target = {}; // 新对象
    for (let prop in obj) {
      if (obj.hasOwnProperty(prop)) {
        if (typeof obj[prop] === "object") target[prop] = deepClone(obj[prop]);
        // Object, 递归的进行深拷贝
        else target[prop] = obj[prop]; // 基础类型, 直接赋值
      }
    }
    return target; // 返回新的对象
  } else if (Array.isArray(obj)) {
    let arr = []; // 新数组
    obj.forEach((item, index) => {
      if (typeof item === "object") arr[index] = deepClone(item);
      else arr[index] = item;
    });
    return arr;
  }
}

const obj1 = {
  a: 1,
  b: {
    a: 2
  },
  c: [1, [2, 3], [4, 5]],
  d: null,
  e: undefined
};

const obj2 = deepClone(obj1);

console.log(obj1 === obj2);     //false
console.log(obj1.b === obj2.b); //false
console.log(obj1.c === obj2.c); //true

console.log("obj1:", JSON.stringify(obj1));
// "obj1:" "{'a':1,'b':{'a':2},'c':[1,[2,3],[4,5]],'d':null}"
console.log("obj2:", JSON.stringify(obj2));
// "obj1:" "{'a':1,'b':{'a':2},'c':[1,[2,3],[4,5]],'d':null}"
```

我们自己实现的深拷贝有很多情况没考虑到，例如函数对象，示例仅做学习参考。如果想在**生产环境**使用深拷贝，请使用第三方库 `lodash` 的 `cloneDeep()` 函数。

如果你的应用场景不是很复杂，可以考虑使用简单的深拷贝方案（零落成泥碾作尘，只有香如故，涅槃重生）  `JSON.parse(JSON.stringify(object))`：

```css
let obj1 = {
  age: 1,
  jobs: {
    first: "FE"
  }
};
let obj2 = JSON.parse(JSON.stringify(obj1));
console.log(obj1 === obj2); // false
console.log(obj1.jobs === obj2.jobs); // false
```

该方法的局限性：

- 会忽略 `undefined，symbol`
- 不能序列化函数
- 不能解决循环引用的对象 - 报错

# 参考文章

刘晓飞. [js深浅拷贝 - 详解](https://juejin.cn/post/7061960684916965407). 稀土掘金.