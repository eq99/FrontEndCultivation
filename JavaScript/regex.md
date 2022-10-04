# 正则表达式详解

## 正则表达式对象

JavaScript 正则表达式对象采用双斜线来声明，例如：

```js
const re = /ab+c/;
```

这种形式在JS加载的时候就解析正则表达式，性能比较好。也可以直接使用构造函数：

```js
const re = new RegExp('ab+c');
```

这种形式可以在运行时构造正则表达式，如果编码时不知道正则表达式的具体内容，或者需要从网络或用户获取时，可以采用这种方法。

【注意】`RegExp`构造函数模式参数时字符串，所以再某些情况下要对字符进项双重转义。所有元字符都必须双重转义，如字面量模式为`/\[bc\]at/`，那么等价的字符串为`"\\[bc\\]at"`。

正则表表示还可以添加标记 (flags) 配置解析模式，常用的标记：

- `g`：表示全局模式,即模式将被应用到所有字符串，而非在发现第一个匹配项时立即停止。
- `i`：表示不区分大小写模式。
- `m`：表示多行模式，即在到达一行文本末尾时还在继续查找下一行中是否存在于模式匹配的项。

```js
const re = /ab+c/gi;

const re = new RegExp('ab+c', 'i');
```

## 正则表达式字符集

- `a`：匹配一个字母 a。
- `a|b`：`|` 表示或，匹配任意一个字符
- `[xyz],[a-c]`：字符集合，或的另一种优雅写法，`[abc]` 表示 `a|b|c`，`-` 表示范围，`0-9` 表示 `[0123456789]`。
- `[^xyz]`表示 `[xyz]` 的补集。

- `.` ：匹配一个换行符(`\n`, `\r`, `\u2028` or `\u2029`)之外的其他字符。（1）如果在字符集里(`[]`)，则当做字面量。（2）`m` 标签不会匹配多行，如果要匹配任意字符请用`[^]`。（3）`s`标签会使允许匹配换行符。

- `\d`：匹配一个数字，等价于 `[0-9]`。
- `\D`：匹配一个非数字，等价于 `[^0-9]`。
- `\w`：等价于 `[A-Za-z0-9_]`，一般编程语言的变量名采用这个字符集。
- `\W`：上面的补集，`[^A-Za-z0-9_]`。
- `\s` ：匹配空白字符，等价于`[ \f\n\r\t\v\u00a0\u1680\u2000-\u200a\u2028\u2029\u202f\u205f\u3000\ufeff]`

- `\S`：上面的补集：`[^ \f\n\r\t\v\u00a0\u1680\u2000-\u200a\u2028\u2029\u202f\u205f\u3000\ufeff]`

- `\t`：匹配一个制表符
- `\r`：匹配一个回车 carriage return
- `\n`：匹配一个换行 linefeed
- `\v`：匹配一个列向制表符 vertical tab
- `\f`：匹配一个换页符 form-feed
- `[\b]`：匹配一个空格 backspace
- `\0`：空白字符
- `/\cX`: `X` 表示 `A–Z`，例如：`/\cM\cJ/` 匹配 "\r\n"

- `\xhh`：匹配两个十六进制表示的字符
- `\xhhhh`: 匹配一个 utf-16 字符
- `\u{hhhh} or \u{hhhhh}`：只有当 `u` 标签设置了才有效，匹配相应Unicode 对应的字符。
- `\p{UnicodeProperty}`, `\P{UnicodeProperty}`：匹配特定的Unicode字符集。

> 注：本节整理自参考文章 2。

## 正则表达式断言

断言包括边界检查，例如单词或者行的起始或结束位置；还包括配向前看，向后看条件，或设置相关条件。

- `^`：匹配输入的起始位置，如果设置了多行标志`m`，则在换行符之后立即匹配。例如：` /^A/` 能匹配 `An A` 里的第一个 `A`，但不能匹配第二个 `A`。**注意**：当`^` 在 `[]` 里面时，表示补集。
- `$`：匹配输入的结束位置，如果设置了多行标志`m`，则在换行符之前立即匹配。例如 `/t$/` 匹配 `eat` 里的 `t` 但不匹配 `winter` 里的 `t`。
- `\b`：匹配单词边界。就是单词第一个字母或最后一个字母和空格之间哪里，这个模式仅表示位置，它的长度为零。例如：（1）`/\bm/` 匹配 `moon` 的 `m`；（2）`/oo\b/` 不匹配 `moon` 中的`oo`，因为不是单词边界；（3）`/oo\b/` 匹配 `moon` 中的`oon`。（4）`/\w\b\w/` 不匹配任何东西，因为两个紧邻的单词之间没有边界。**注意**：匹配空格请使用`[\b]`。
- `\B`：匹配非单词边界。这个地方的前后两个种类必须一样，要么都是单词，要么都不是单词。例如：`/\Bon/` 匹配 `at noon` 的 `on`，`/ye\B/` 匹配 `possibly yesterday` 中的 `ye`。
- `x(?=y)`：向前看条件，匹配 `x` 当且仅当 `x` 后面跟着 `y`。例如：`/Jack(?=Sprat)/` 匹配 `JackSprat` 中的`Jack`，但不匹配 `Jacksprat`。
- `x(?!y)`：向前看条件取反，匹配`x` 当且仅当 `x` 后面**不是** `y`。例如：`/\d+(?!\.)/` 表示不跟着小数点的字符，`/\d+(?!\.)/.exec('3.141')` 匹配 "141" 但不匹配"3"。
- `(?<=y)x`：向后看条件，匹配 `x` 当且仅当 `x` 前面是 `y`
- `(?<!y)x`：向后看条件取反，匹配 `x` 当且仅当 `x` 前面不是 `y`。

> 注：本节参考文章: 3

## 分组与引用

引用分组的目的是对重复出现的文本进行匹配，**注意，不是出现重复的模式，而是出现重复的文本(匹配结果)**。定义分组的三种方法：

- `(x)`：把括号内的规则看做一个分组。正则表达式引擎从左向右扫描，遇到一个可捕获分组，就会给分配一个组号，规则是从1开始，依次加1，需要注意的是，有一个隐含的全局分组，它的组号是0，也就是整个正则表达式。
- `(?<name>x)`：含义与上面类似。除了可以通过组号引用结果之外，还可以使用`name` 来引用匹配结果。
- `(?:x)`：无捕获分组，无法设置组名，也不会分配组号。

引用有两种方式，编号与组名：

- `\n`：通过组号引用分组，`n` 代表组号。例如 `/apple(,)\sorange\1/` 匹配 `"apple, orange, cherry, peacle"` 中的 `"apple, orange,"`。
- `\k<name>`：通过组名引用分组。例如 `/(?<title>\w+), yes \k<title>/` 匹配 `"Do you copy? Sir, yes Sir!"` 中的 `"Sir, yes Sir"` 。

【**注意**】由于正则表达式的解析是有顺序的，从正则表达式的开头向后解析，引用分组的编号和名称，必须是前面已经存在的；如果在当前位置引用的编号和名称不存在，那么模式解析就会报错。

【**注意**】括号括起来的不一定是分组，还有可能是断言。

【例1】匹配两个重复单词

正则表达式：`\b(\w+)\b\s+\1\b`

> 注：本节参考文章：4, 5

## 正则表达式量词

正则表达式量词表示重复匹配个数。

- `x{n,m}`：匹配个数大于等于 $n$，小于等于 $m$，其中 $n\ge 0, m\ge n$ 。例如：`/a{1,3}/` 不会匹配 `cndy`，在 `candy` 中匹配 `a`，在 `caandy` 中匹配 `aa`，在 `caaaaaaadny` 中匹配开头的三个 `a`。

- `x{n,}`：重复次数大于等于 $n$。例如：`/a{2,}/` 不匹配 `candy`，在`caandy` 中匹配 `aa`，在 `caaaaaandy` 中匹配 `aaaaaa`。
- `x{n}`：刚好 `n` 个，即 $m=n$。例如：`/a{1}/` 不匹配 `cndy`，也不匹配 `caandy`，只匹配 `candy` 中的 `a`。
  - `x*`：等价于 `x{0,}`。例如 `/bo*/` 匹配 `A ghost booooed` 中的 `A ghost booooed`，匹配 `A bird warbled` 中的 `b`，但不匹配 `A bird warbled`。
- `x+`，等价于 `x{1,}`
- `x?`：等价于`x{0,1}`，一般用作可选的字符。`a?` 匹配 `cndy`，也会匹配 `candy` 中第一个 `a`。

【**特别说明**】在量词之后再次使用 `?`，会使匹配变成**惰性**的。

```js
x*?
x+?
x??
x{n}?
x{n,}?
x{n,m}?
```

**惰性(lazy)**与**贪婪(greedy)** 是正则表达式匹配的两种方式，贪婪模式使得匹配的结果长度尽可能的长，而惰性（非贪婪）模式使得匹配的结果尽可能的短。相当于把区间 `{n,m}` 变成 `{n,n}`

【例1】`5?` 能匹配匹配 0个或1个 `5`，`5??` 只能0个字符了。

## 与正则表示相关的方法

### `RegExp.prototype.test()`

接受一个字符串参数，在模式与该参数匹配的情况下返回 `true`，否则返回 `false`。

```js
const str = 'hello';
const re1 = /E?/i;
const re2 = /e{2}/;

console.log(re1.test(str));//true
console.log(re2.test(str));//false
```

### `RegExp.prototype.exec()`

该方法是专门为捕获组而设计的，其接受一个参数，即要应用模式的字符串，然后返回包含第一个匹配项信息的数组；或者在没有匹配项的情况下返回`null`。返回的数组虽然是`Array`的实例，但是包含两个额外的属性：`index`和`input`。其中`index`表示匹配项在字符串中的位置，而`input`表示应用字符串表达式的字符串。

```js
const text = 'mom and dad and baby';
const pattern = /mom( and dad( and baby)?)?/gi;
const matches = pattern.exec(text);
console.log(matches);
console.log(matches.index); //0
console.log(matches.input); //mom and dad and baby
console.log(matches[0]); //mom and dad and baby 默认全局分组 0
console.log(matches[1]); //and dad and baby 组 1
console.log(matches[2]); //and baby 组 2
```

对于`exec()`方法而言，即使在模式中设置了全局标志`g`，它每次也只是返回一个匹配项。在不设置全局标志的情况下，在同一个字符串上多次调用`exec()`方法将始终返回第一个匹配项的信息。而在设置全局标志的情况下，每次调用`exec()`则都会在字符串中继续查找新匹配项，如下例子：

```js
const text = 'cat, bat, sat, fat';
const pattern1 = /.at/;

let matches1 = pattern1.exec(text);
console.log(matches1.index); //0
console.log(matches1[0]); //cat
console.log(pattern1.lastIndex); //0

matches1 = pattern1.exec(text);
console.log(matches1.index); //0
console.log(matches1[0]); //cat
console.log(pattern1.lastIndex); //0

const pattern2 = /.at/g;

let matches2 = pattern2.exec(text);
console.log(matches2.index); //0
console.log(matches2[0]); //cat
console.log(pattern2.lastIndex); //3

matches2 = pattern2.exec(text);
console.log(matches2.index); //5
console.log(matches2[0]); //bat
console.log(pattern2.lastIndex); //8
```

### `String.prototype.search()`

在字符串搜索符合正则的内容，搜索到就返回出现的位置（从0开始，如果匹配的不只是一个字母，那只会返回第一个字母的位置）， 如果搜索失败就返回 -1：

```js
const str = 'abcba';
const re = /B/i;
const re2 = /B/;
console.log(str.search(re)); // 1
console.log(str.search(re2)); // -1
```

### `String.prototype.match()` 

使用方法：`str.match(regexp)`

参数 `regexp` 是一个正则表达式对象，或者任何实现了 [`Symbol.match`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/match)  方法的对象。如果参数不符合以上两个条件，那么会尝试使用隐式转换 `new RegExp(regexp)` 成正则表达式对象。如果不传参数，默认参数是`/(?:)/`。

返回值是一个匹配的字符串数组，如果没有匹配的结果则返回 `null`。该结果取决于参数`regexp` 没有有设置标志 `g`。

- 如果没有 `g`，会返回第一个匹配到的结果以及相关的分组结果，这与 `exec` 的结果相同。
- 如果设置了 `g`，会返回所有的结果，不包括分组。

```js
const str = 'For more information, see Chapter 3.4.5.1';
const re = /see (chapter \d+(\.\d)*)/i;
const found = str.match(re);

console.log(found);

// logs [ 'see Chapter 3.4.5.1',
//        'Chapter 3.4.5.1',
//        '.1',
//        index: 22,
//        input: 'For more information, see Chapter 3.4.5.1' ]

// 'see Chapter 3.4.5.1' is the whole match.
// 'Chapter 3.4.5.1' was captured by '(chapter \d+(\.\d)*)'.
// '.1' was the last value captured by '(\.\d)'.
// The 'index' property (22) is the zero-based index of the whole match.
// The 'input' property is the original string that was parsed.
```

```js
const str = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
const regexp = /[A-E]/gi;
const matches = str.match(regexp);

console.log(matches);
// ['A', 'B', 'C', 'D', 'E', 'a', 'b', 'c', 'd', 'e']
```

> 注：本节参考文章：6, 7

### `String.prototype.replace()`

JavaScript 的替换函数非常强大。使用方法：`replace(pattern, replacement)`。

参数 `pattern` 可以是 普通字符串或者实现了  [`Symbol.match`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol/match)  方法的对象（通常是正则表达式）。

参数 `replacement` 可以是普通字符串或函数。

该方法不会改变原字符串，只是返回一个新的字符串。替换的时候只会替换一次，除非设置了`g`标志，或者使用  [`replaceAll()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replaceAll) 。

```js
const p = 'The quick brown fox jumps over the lazy dog. If the dog reacted, was it really lazy?';

console.log(p.replace('dog', 'monkey'));
// expected output: "The quick brown fox jumps over the lazy monkey. If the dog reacted, was it really lazy?"


const regex = /Dog/i;
console.log(p.replace(regex, 'ferret'));
// expected output: "The quick brown fox jumps over the lazy ferret. If the dog reacted, was it really lazy?"

// 空替换
"xxx".replace("", "_"); 
// "_xxx"
```

在替换的时候还可以使用 `$` 引用匹配的结果：

| **字符**  | **替换文本**                     |
| --------- | -------------------------------- |
| `$$`      | 插入一个`$`                      |
| `$&`      | 与 regexp 相匹配的子串           |
| *$`*      | 位于匹配子串左侧的文本           |
| `$'`      | 位于匹配子串右侧的文本           |
| `$n`      | 引用第 `n` 个分组，`n` 取值 1~99 |
| `$<name>` | 引用命名的组                     |

【例】大家可以自行测试更多的内容。

```js
const str3 = "这是一段原始文本,3c这要替换4d!";
const newStr = str3.replace(/([0-9])(?<name>[a-z])/g, "$<name>");
console.log(newStr);
```

函数式替换。函数式替换的回调函数声明如下：

```js
function replacer(match, p1, p2, /* …, */ pN, offset, string, groups) {
  return replacement;
}
```

参数 `match`表示匹配的字符串，即 `$&`

参数 `p1, p2, …, pN` 表示各个分组代表的字符串，相当于` $n`

参数 `offset` 代表匹配到的字符串起始下标。

参数 `string`代表原字符串。

参数`groups` 代表命名分组对象，键名是组名，值是匹配到的字符串(若未匹配到则是 undefined)。该参数存在需要至少有一个命名组。

【例】

```js
function replacer(match, p1, p2, p3, offset, string) {
  // p1 is non-digits, p2 digits, and p3 non-alphanumerics
  return [p1, p2, p3].join(' - ');
}
const newString = 'abc12345#$*%'.replace(/([^\d]*)(\d*)([^\w]*)/, replacer);
console.log(newString);  // abc - 12345 - #$*%
```

> 注：本节参考文章：8

## 参考文章 

1. 正则表达式测试工具. https://regex101.com/
2. [Character classes](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions/Character_Classes). MDN.
3. [Assertions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions/Assertions). MDN.
4. 悦光阴. [正则表达式 第三篇：分组和捕获](https://www.cnblogs.com/ljhdo/p/10678281.html). 博客园.
4. [Groups and backreferences](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions/Groups_and_Backreferences). MDN.
4. [String.prototype.match()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/match). MDN.
7. 风雨后见彩虹. [JS正则表达式详解 ](https://www.cnblogs.com/moqiutao/p/6513628.html). 博客园.
8. [String.prototype.replace()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replace#specifying_a_function_as_the_replacement). MDN.



★ 文章整理自网络，若有疏漏请在评论区指正。

★ 本文开源，转载只需注明出处：小星森. [JavaScript 正则表达式详解](https://zhuanlan.zhihu.com/p/570372405). 知乎.

★ 商业合作请发私信或邮件：zuiaiqiansen@163.com，并注明主题：商业合作。

★ 宣传做的好，产品差不了，宣传做的棒，产品No.1。我是前端小星森，让用户看到最伟大的产品。本人会一点点前端，如果贵公司有一流的产品或服务需要前端工程师展示，或需要一个前端工程师实现您的创业梦想，请发邮件：zuiaiqiansen@163.com，并注明主题：招聘前端工程师。
