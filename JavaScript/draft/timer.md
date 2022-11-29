## 定时器

## setInterval

定时器开启方法：

```js
setInterval(code)
setInterval(code, delay)
	
setInterval(func)
setInterval(func, delay)
setInterval(func, delay, arg0)
setInterval(func, delay, arg0, arg1)
setInterval(func, delay, arg0, arg1, /* … ,*/ argN)
```
`setInterval()` 会返回一个标识符（其实是一个整数），表示定时器的“身份证”，作为 `clearInterval()` 的参数来清除定时器。

【例】给定时器传参数。

```js
const intervalID = setInterval(myCallback, 1000, "Parameter 1", "Parameter 2");

function myCallback(a, b) {
  console.log(a);
  console.log(b);
}
```

【例】清除定时器

```html
<style>
  .go {
    color: green;
  }

  .stop {
    color: red;
  }
</style>

<div id="my_box">
  <h3>Hello World</h3>
</div>
<button id="start">Start</button>
<button id="stop">Stop</button>

<script>
  // variable to store our intervalID
  let nIntervId;

  function changeColor() {
    // check if an interval has already been set up
    if (!nIntervId) {
      nIntervId = setInterval(flashText, 1000);
      //console.log(nIntervId, typeof nIntervId) // 6 "number"
    }
  }

  function flashText() {
    const oElem = document.getElementById("my_box");
    oElem.className = oElem.className === "go" ? "stop" : "go";
  }

  function stopTextColor() {
    clearInterval(nIntervId);
    // release our intervalID from the variable
    nIntervId = null;
  }
  document.getElementById("start").addEventListener("click", changeColor);
  document.getElementById("stop").addEventListener("click", stopTextColor);
</script>
```

【提示】定时器的设定是异步的，简单理解异步：就是把活丢给别人(工作线程)去干。`setInterval()` 就像给秘书打个电话：每隔一小时帮我检查一下邮箱。

## setTimeout

等待一定时间后执行函数：

```js
setTimeout(code)
setTimeout(code, delay)

setTimeout(functionRef)
setTimeout(functionRef, delay)
setTimeout(functionRef, delay, param1)
setTimeout(functionRef, delay, param1, param2)
setTimeout(functionRef, delay, param1, param2, /* … ,*/ paramN)
```

可以通过清除计时器取消：

```js
let timeoutID = setTimeout(code, delay)
clearTimeout(timeoutID)
```