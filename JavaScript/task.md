macrotask 和 microtask 表示异步任务的两种分类。

在挂起任务时，`JS 引擎`会将所有任务按照类别分到这两个队列中，首先在 `macrotask` 的队列（这个队列也被叫做 `task queue`）中取出第一个任务，执行完毕后取出 `microtask` 队列中的所有任务顺序执行；之后再取 `macrotask` 任务，周而复始，直至两个队列的任务都取完。

浏览器环境下的异步任务分为宏任务(macroTask) 和微任务(microTask)：

宏任务(macroTask)：`script` 中代码、`setTimeout`、`setInterval`、`I/O`、`UI render`；

微任务(microTask)： `Promise`、`Object.observe`、`MutationObserver`。



作者：慕时_木雨凡
链接：https://www.jianshu.com/p/75107522813f
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

作者：慕时_木雨凡
链接：https://www.jianshu.com/p/75107522813f
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。



```javascript
console.log('1')
setTimeout(() => {
  console.log('2')
})
new Promise((resolve, rejects) => {
  console.log('3')
  resolve()
}).then(() => {
  console.log('4')
})
console.log(5)复制代码
```

诸位可以先给出来一个自己的答案，运行一下结果，看看是否与自己想的一致。


作者：Reed
链接：https://juejin.cn/post/6844903814508773383
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。