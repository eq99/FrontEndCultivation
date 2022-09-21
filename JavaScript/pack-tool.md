# 1、简单说一下 webpack 的构建流程

webpack 的运行流程是一个串行的过程，从启动到结束会依次执行以下流程：

1. `初始化参数：`  从配置文件和shell 语句中读取与合并参数，得到最终的参数。
2. `开始编译：`  用上一步得到的参数初始化 Compiler 对象，加载所有配置的插件，执行 compiler 对象的 run 方法开始执行编译。
3. `确定入口：` 根据配置中的 entry 找出所有的入口文件。
4. `编译模块：` 从入口文件出发，调用所有配置的 Loader 对模块进行编译，找出该模块依赖的模块，再递归本步骤直到所有依赖文件都经过本步骤的处理。
5. 完成模块编译： 在经过第 4 步使用 Loader 编译完所有模块之后，得到每个模块被编译后的最终内容以及它们之间的依赖关系。
6. `输出资源：` 根据入口和模块之间的关系，组装成一个个包含多个模块的 chunk，再把每个 chunk 转换成一个单独的文件加入到输出列表，这步是可以修改输出内容的最后机会。
7. `输出完成：` 在确定输出内容之后，根据配置确定输出的路径和文件名，把文件内容写入到文件系统。

# Loader 和 Plugin 的区别

`Loader` 本质就是一个函数，在该函数对接收到的内容进行转换，返回转换后的结果。你可以理解为是一个“管道”，在外部接收到的内容通过这个“管道”进行转换，然后再将转换后的结果输出。因为 webpack 只认识 js，所以，你也可以将 Loader 称之为“翻译官”，对其他类型的资源进行转译的预处理工作。

`Plugin` 直译为插件，基于事件流框架 `Tapable`。Plugin 可以扩展 webpack 的功能，让 webpack 具有更多的灵活性。在 Webpack 运行的生命周期中会广播出许多事件，Plugin可以监听这些事件，在合适的时机通过 webpack 提供的 API 改变输出结果。其实 webpack 就是微内核架构的一个例子，本质上就是一个是很小的功能，它并没有携带任何业务的功能。

# 4、如何提高 webpack 的打包速度？

1. 多入口情况下，使用 `optimization.splitChunks `来提取公共代码。
2. 通过 `externals` 配置来提取常用库。
3. 利用 `DllPlugin` 和` DllReferencePlugin` 预编译资源模块，通过 `DllPlugin` 来对那些我们引用但是绝对不会修改的npm包来进行预编译，再通过`DllReferencePlugin` 将编译编译的模块加载进来。
4. 使用 `thread-loader` 实现多进程加速编译。
5. 使用 `terser-webpack-plugin` 对js进行代码压缩。
6. 优化 `resolve` 配置缩小范围。
7. 使用 `tree-shaking` 和 `Scope hoisting` 来剔除多余代码。

# 5、如何减少 webpack 打包体积？

1. 使用 `externals `配置来提取常用库。
2. 使用 `tree-sjaking` 和 `scope hoisting` 来剔除多余代码。
3. 使用 `optimize-css-assets-webpack-plugin` 压缩css。
4. 使用 `terser-webpack-plugin` 对 js 进行代码压缩。

# 6、webpack 有哪些常见的 loader？你用过哪些 loader？

- `cache-loader`：可以在一些性能开销较大的 Loader 之前添加，目的是将结果缓存到磁盘里。
- `file-loader`：把文件输出到一个文件夹中，在代码中通过相对 URL 去引用输出的文件（处理图片、字体、图标）。
- `url-loader`：与file-loader 类似，区别是用户可以设置一个阈值，大于阈值会交给 file-loader，小于阈值时返回文件 base64 形式编码（处理图片）。
- `image-loader`：加载并且压缩图片文件。
- `babel-loader`：把 ES6 转换成 ES5。
- `ts-loader`：将 typescript 转换成 JavaScript。
- `svg-inline-loader`：将压缩后的SVG内容注入代码中。
- `raw-loader`：加载文件原始内容（utf-8）。
- `sass-loader`：将 scss/sass 代码转换成 css。
- `css-loader`：加载 css，支持模块化、压缩、文件导入等特性。
- `less-loader`：将 less 代码转换成 css。
- `style-loader`：生成 style 标签，将 js 中的样式资源插入，并添加到 header 中生效。
- `postcss-loader`：扩展 css 语法，使用下一代 css，可以配合 autoprefixer 插件自动补齐 css3 前缀。
- `eslint-loader`：通过 eslint 检查 JavaScript 代码。
- `tslint-loader`：通过tslint 检查 typesc 代码。
- `vue-loader`：加载 vue.js 单文件组件。
- `awesome-typescript-loader`：将 typescript 转换成 JavaScript，性能优于 ts-loader。

# 7、webpack 有哪些常见的 plugin？你用过哪些 plugin？

- `ignore-plugin：`忽略部分文件。
- `html-webpack-plugin：`简化 html 文件创建。
- `web-webpack-plugin：`可以方便地为单页应用输出 html，比 html-webpack-plugin 好用。
- `terser-webpack-plugin：`支持压缩ES6。
- `optimize-css-assets-webpack-plugin：`压缩css代码。
- `mini-css-extract-plugin：`分离样式文件，css 提取为单独文件，支持按需加载。
- `werviceworker-webpack-plugin：`为网页应用追加离线缓存功能。
- `clean-webpack-plugin：`目录清理。
- `ModuleconcatenationPlugin：`开启 Scope Hoisting。
- `webpack-bundle-analyzer：`可视化 webpack 输出文件的体积（业务组件、依赖第三方模块）。
- `speed-measure-webpack-plugin：`可以看到每个 loader 和 plugin 执行耗时（这个打包耗时、plugin 和 loader 耗时）。 -` HotModuleReplacementPlugin：`模块热替换。



# vite

webpack会先打包，然后启动开发服务器，请求服务器时直接给予打包结果。 而vite是直接启动开发服务器，请求哪个模块再对该模块进行实时编译。 由于现代浏览器本身就支持ES Module，会自动向依赖的Module发出请求。vite充分利用这一点，将开发环境下的模块文件，就作为浏览器要执行的文件，而不是像webpack那样进行打包合并。 由于vite在启动的时候不需要打包，也就意味着不需要分析模块的依赖、不需要编译，因此启动速度非常快。当浏览器请求某个模块时，再根据需要对模块内容进行编译。这种按需动态编译的方式，极大的缩减了编译时间，项目越复杂、模块越多，vite的优势越明显。 在HMR方面，当改动了一个模块后，仅需让浏览器重新请求该模块即可，不像webpack那样需要把该模块的相关依赖模块全部编译一次，效率更高。 当需要打包到生产环境时，vite使用传统的rollup进行打包，因此，vite的主要优势在开发阶段。

# 参考文章 

1. 青峰10. [Webpack 高频面试题](https://juejin.cn/post/7066807280557096974).稀土掘金.
2. Delta🎵.[vite和webpack的区别](https://juejin.cn/post/6893699833425559559). 稀土掘金.