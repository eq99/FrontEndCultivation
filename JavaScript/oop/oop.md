# JavaScript 面向对象编程

## 封装

面向对象编程是一种程序设计思想，编程的关注点在对象而不是过程。对象主要有两大部分组成：一部分是状态，另一部分是行为。

什么是状态？状态用于描述物体的性质，通常体现为某些量化指标。例如：

1）描述一个人的状态：身高，体重，年龄等。

2）描述一辆车的状态：品牌，价位，油量等。

3）描述一个游戏角色：血量，蓝量，属性，技能等。

这些状态有的可以用数字精确的描述，但有的不能。通过生活实际我们可以切身的感受到状态不是一成不变的，而会按照某种规律不断变化着，**变化**所对应的概念就是行为，例如：

1）一个人的年龄会不断增长。

2）汽车的油量随着汽车的开动而减少。

3）游戏角色受到攻击血量会减少，释放技能会消耗蓝量。

状态在计算机内部的表现形式是**变量**，行为在计算机内部的表现形式是**函数**。



在设计程序的时候，有两种典型的思路：

1）游戏角色释放技能，然后减少蓝量，技能打到敌人，敌人减少血量，这种一步一步的设计风格称为过程式的设计。

2）把游戏角色看做一个整体，不用关心怎么减少蓝量，怎么计算血量，只需要关注“**释放技能**”这一行为即可。

有时候，关注行为比关注状态更加简洁高效，不管这个人是高还是矮，是胖还是瘦，只要他代码写得好，能说会道，那么给人的感受就是大佬。甚至有一类对象，它只有行为方式，没有状态，计算机专业属于叫做**接口**或者**抽象类**，例如：

1）能说会道的人：这句话并没有指定某个具体的人，而是描述了具有“**能说会道**”这种行为的人。

2）射手角色：这句话并没有说是哪个游戏角色，但我们知道这个游戏角色有一个行为：能射。

抽象类有一个优秀的性质：能与实体类轻松结合。就像是一部武功秘籍，大家都可以练，练出来都能打。



通过以上例子，可以总结出面向对象的一个特性：**封装**：不必关注对象的内部细节，只需要关注对象的行为。

### JavaScript 如何实现封装

JavaScript 实现封装的方式非常简单，内置的 Object 就是封装的体现：

```js
let car = {
  name: "Benz",
  oil: 120,
  run() {
    this.oil--;
    console.log(`我的${this.name}在奔驰, 当前油量：${this.oil}`);
  },
};

car.run();//我的Benz在奔驰, 当前油量：119
car.run();//我的Benz在奔驰, 当前油量：118
```

## 构造函数

看了上面的例子，有的同学可能表示不服，我骑的是摩托车，能不能给我造一辆摩托车？

```js
let moto = {
  name: "雅马哈",
  oil: 100,
  run() {
    this.oil--;
    console.log(`我的${this.name}在奔驰, 当前油量：${this.oil}`);
  },
};

moto.run(); //我的雅马哈在奔驰, 当前油量：99
moto.run(); //我的雅马哈在奔驰, 当前油量：98
```

没问题，我们造了一辆摩托车，这时候五菱宏光不乐意了，我秋名山神车居然排不上榜，你几个意思？仔细观察发现，这些车子除了名字不同，油量不同之外，没有其他区别，那就给大伙造一个车场，自己想开什么车就造什么车：

```js
function Factory(name, oil) {
  const car = {
    name,
    oil,
    run() {
      this.oil--;
      console.log(`我的${this.name}在奔驰, 当前油量：${this.oil}`);
    },
  };

  return car;
}

let car = Factory("五菱宏光", 180);
car.run();
```

工厂函数是不是很方便呀，利用工厂模式可以批量生产出对象。既然工厂函数这么方便，于是就演化成了一种程序设计模式：即工厂模式，或创建型模式。

### JavaScript 语言的构造模式。

创建对象是编程的家常便饭，任何语言都有自己创建对象的方法，C 语言是 malloc，C++是构造函数，Java 是 `new Class()`。JavaScript 非常简单：直接对着函数使用 new 操作符就可以构造对象：

```js
function Car() {
  let name = "Audi";
  console.log(`Drive ${name}`);
}

let car = new Car();//Drive Audi
console.log(car);//{}
```

上述例子表明，对函数使用 new 操作符可以返回一个对象，但是这个对象是空对象，怎么样才能像工厂函数那样创造出对象呢？

```js
function Car() {
  this.name = "Audi";
  this.oil = 120;
  this.run = function () {
    this.oil--;
    console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
  };
}

let car = new Car();
console.log(car); //{"name":"Audi","oil":120}
car.run();        //我的Audi在飞驰, 当前油量：119
```

成功了，我们造出了一个汽车对象。

### JavaScript 构造函数

上例的函数 Car() 称为构造函数，当 new 构造函数的时候发生了什么？

1）创建一个新的空对象，即在获取堆区的一块内存空间

2）改变 `this` 的指向，指向刚创建的新对象

3）调用构造函数体

4）返回对象：

- 如果没有 `return`语句，则返回创建好的新对象，即 `return this`。
- 如果显示调用 `return` 语句，如果返回基本类型，则被忽略，依然返回新对象；如果返回对象，则返回对象。

最后的返回值听起来有点绕，看个例子：

```js
function Car() {
  this.name = "Audi";
  this.oil = 120;
  this.run = function () {
    this.oil--;
    console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
  };
  
  return false; // 返回几本类型，与不返回值是一样的，都是新对象
}

let car = new Car();
console.log(car);
car.run();
```

返回对象：

```js
function Car() {
  this.name = "Audi";
  this.oil = 120;
  this.run = function () {
    this.oil--;
    console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
  };

  return {
    name: "Benz",
  };
}

let car = new Car();
console.log(car);//{"name":"Benz"}
```

最后返回值是对象，那么前面的过程就是白忙活了一场，所以**作为构造函数的函数，不要返回值**。

### 工厂函数存在的问题

工厂函数能够批量生产对象，但是这些对象之间毫无关联：

```js
function Car() {
  this.name = "Audi";
  this.oil = 120;
  this.run = function () {
    this.oil--;
    console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
  };
}

let car1 = new Car();
let car2 = new Car();
console.log(car1 === car2); //false
console.log(car1.run === car2.run); //false
```

car1 与 car2 的 run 函数是一模一样的，但是存在两份，要是能够实现共享该有多好啊？

## 原型

JavaScript 通过原型对象实现属性与方法的共享，所谓的原型对象是一个特殊对象：构造函数通过 `prototype` 属性获取原型对象，实例通过 `__proto__` 获取原型对象：

```js
function Car() {
  this.name = "Audi";
  this.oil = 120;
  this.run = function () {
    this.oil--;
    console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
  };
}

let car1 = new Car();
let car2 = new Car();
console.log(Car.prototype); //{}
console.log(car1.__proto__ === Car.prototype);//true
console.log(car2.__proto__ === Car.prototype);//true
```

![prototype](https://pic.imgdb.cn/item/6381ded616f2c2beb1b251e3.jpg)

控制台打印出原型对象，它有两个属性：

一个是 `constructor`，它指向 `Car`，其实有点像双向绑定的意思。

另一个是 `[[Prototype]]`，它就是 `__proto__` 属性。

从上数例子可以看出原型对象就是构造函数与相关实例的公共区域。原型对象除了共享之外，实例能够直接使用原型对象上的方法或属性：

```js
function Car() {
  this.name = "Audi";
  this.oil = 120;
}

let car1 = new Car();
let car2 = new Car();

Car.prototype.run = function () {
  this.oil--;
  console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
};

car1.run();//我的Audi在飞驰, 当前油量：119
car1.run();//我的Audi在飞驰, 当前油量：118
car2.run();//我的Audi在飞驰, 当前油量：119
```

技巧：一般把方法定义在 `prototype` 上面，数据定义到 `this` 上面。

![原型关系](https://pic.imgdb.cn/item/6381d74b16f2c2beb1a61be3.png)

### 原型继承

细心的同学发现了一个问题：原型对象也有原型，现在我们来做一个有意思的事情：

```js
function Vehicle(name) {
  this.name = name;
}

Vehicle.prototype.move = function () {
  console.log(`${this.name} is moving...`);
};

function Car() {
  this.name = "Audi";
  this.oil = 120;
}

Car.prototype.run = function () {
  this.oil--;
  console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
};

Car.prototype.__proto__ = new Vehicle("Car");

let car = new Car();
car.run(); //我的Audi在飞驰, 当前油量：119
car.move(); //Audi is moving...
```

通过 `prototype` 和 `__proto__` 打通了上下游关系，这种关系可以用如下原型链形象的描述：

![原型链](https://pic.imgdb.cn/item/6381db0b16f2c2beb1ad4f2a.jpg)

当打通了原型链之后，所有下游的对象就会根据就近原则，一层层访问链上的属性和方法，真的是相当的方便。用计算机专业术语来说是**原型继承**。

### 多态 

上面一句简短的话：就近原则就实现了**多态性**。所谓的多态性有很多含义，在面向对象中的术语是**重载**，其中包括运算符重载与方法重载，所谓重载就是自定义父类的方法的行为。JavaScript 通过原型链实现了继承，只需要在最近的原型上声明同名的方法即可实现重载：

```js
function Vehicle(name) {
  this.name = name;
}

Vehicle.prototype.move = function () {
  console.log(`${this.name} is moving...`);
};

function Car() {
  this.name = "Audi";
  this.oil = 120;
}

Car.prototype.run = function () {
  this.oil--;
  console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
};

// 覆盖 move 方法
Car.prototype.move = function () {
  console.log(`国产${this.name} 就是强。`);
};

Car.prototype.__proto__ = new Vehicle("Car");

let car = new Car();
car.run(); //我的Audi在飞驰, 当前油量：119
car.move(); //国产Audi 就是强。
```

###  再看 new 

1）创建一个新的空对象，即在获取堆区的一块内存空间

2）将空对象的 `__proto__` 指向构造函数的 `prototype`

3）改变 `this` 的指向，指向刚创建的新对象

4）调用构造函数体

5）返回对象：

- 如果没有 `return`语句，则返回创建好的新对象，即 `return this`。
- 如果显示调用 `return` 语句，如果返回基本类型，则被忽略，依然返回新对象；如果返回对象，则返回对象。

### 小结

通过以上例子的讲解，我们已经学习到了 JavaScript 面向对象的三个思想：

1）封装：使用对象包裹属性和方法，隐藏属性，通过方法体现行为。

2）继承：通过原型链实现继承。

3）多态：通过就近原则覆盖同名属性或方法实现重载。

## ES6 Class

JavaScript 中的面向对象程序设计方法有些另类，与主流的面向对象语言格格不入，所以 ES6 引入了 Class 语法糖。上述例子可以用 ES6 改造一下：

```js
class Vehicle {
  constructor(name) {
    this.name = name;
  }

  move() {
    console.log(`${this.name} is moving...`);
  }
}

class Car extends Vehicle {
  constructor() {
    super();
    this.name = "Audi";
    this.oil = 120;
  }

  run() {
    this.oil--;
    console.log(`我的${this.name}在飞驰, 当前油量：${this.oil}`);
  }
}

let car = new Car();
car.run();  //我的Audi在飞驰, 当前油量：119
car.move(); //Audi is moving...
```

上面的代码与之前的版本是一模一样的，只是写起来更加优雅，与其他编程语言的编程风格类似。

更多有关类的操作，例如 private 与 public 属性，静态方法，可以阅读MDN文档[^1]。

## 参考文章

[^1]: MDN. [Classes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Classes)

> ♥ 我是前端工程师：你的甜心森。非常感谢大家的点赞与关注，欢迎大家参与讨论或协作。
>
> ★ 本文[开源](https://github.com/xiayulu/FrontEndCultivation)，采用 [CC BY-SA 4.0 协议](http://creativecommons.org/licenses/by-sa/4.0/)，转载请注明出处：[前端工程师的自我修养](https://github.com/xiayulu/FrontEndCultivation). GitHub.com@xiayulu.
>
> ★ 创作合作或招聘信息请发私信或邮件：zuiaiqiansen@163.com，注明主题：创作合作或**招聘前端工程师**。