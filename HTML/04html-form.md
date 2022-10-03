# HTML 标签之表单

本讲知识点：

- 理解表单提交数据的原理。
- 掌握常见的 `input` 类型：`text, password, file, number, email, radio, date, checkbox`。
- 掌握 `textarea, select`。

表单的作用：HTML 表单用于接收不同类型的用户输入，用户提交表单时向服务器传输数据，从而实现用户与Web服务器的交互。

一个简单的表单：

```html
  <form action="/signup" method="POST">
    <div>
      <label for="name">昵称</label>
      <input type="text" id="name" name="name" placeholder="请输入昵称" />
    </div>
    <div>
      <label for="password">密码</label>
      <input type="password" id="password" name="password" placeholder="请输入密码" />
    </div>
    <button type="submit">注册</button>
  </form>
```

一个表单由一个 `form` 元素组织，其中包含若干个输入标签，如：`input, textarea, select`，一个提交按钮 `button`。

## `input` 元素

`input` 元素通过 `type` 属性设置输入类型，不同的 type 有不同的功能。

例如密码输入框：

```html
<input type="password">
```

在输入时会隐藏密码。

例如文件上传：

```html
<input type="file">
```

点击它会弹出文件上传的窗口。更多 `input` 类型请参考 [MDN input type](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input)。

`input` 元素除了之前见到过的 `type, id, name` 属性外，还有

-  `value` ：输入框的值。
- `disabled`：输入框是否被禁用。
- `required`：该字段是否是必填的。

更所的属性及其含义参考 [MDN input attributes](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attributes)。

## `label` 元素

`label` 元素用于给`input` 元素添加提示信息，当点击 label 元素时，会聚焦到相应输入框：

```html
<div>
    <label for="cheese">Do you like cheese?</label>
    <input type="checkbox" name="cheese" id="cheese">
</div>

<div>
    <label for="peas">Do you like peas?</label>
    <input type="file" name="peas" id="peas">
</div>
```

使用方法是把 `label` 的 `for` 属性指向 `input` 元素的 `id` 。最经典的使用场景有两个：

1）用作单选按钮的提示，[在线示例](https://www.w3schools.com/tags/tag_label.asp)：

```html
<form>
  <input type="radio" id="html" name="fav_language" value="HTML">
  <label for="html">HTML</label><br>
  <input type="radio" id="css" name="fav_language" value="CSS">
  <label for="css">CSS</label><br>
  <input type="radio" id="javascript" name="fav_language" value="JavaScript">
  <label for="javascript">JavaScript</label><br><br>
  <input type="submit" value="Submit">
</form>
```

2）文件上传：

```html
<form>
  <input type="file" name="file" id="file" class="inputfile" hidden />
  <label for="file" class="upload">
    📤点击或拖拽上传
  </label>
</form>
```

```css
.upload {
  border: 2px dashed #aaa;
  display: inline-block;
  padding: 5rem;
  cursor: pointer;
}
```

## `input` 元素样式

`input` 元素默认是 `inline-block` 元素，因此可以设置宽高，内外边距。浏览器的输入框一般带有一个黑框框，点击的时候黑框框还会加粗，可以使用如下样式设置或者去除黑框框：

```css
/*ios 系统*/
appearance:none;
-webkit-appearance:none;
-moz-appearance: none; 
-o-appearance: none;

/*设置或去除边框*/
border: none;
box-shadow:none;

/*去除或设置边线*/
outline: none;
```

典型的案例：

```html
<body>
  <form class="form">
    <label for="nickname" class="required">昵称</label>
    <input type="text" name="nickname" id="nickname" class="form-control" required />
  </form>
</body>
```

```css
.form-control {
  outline-style: none;
  background-color: transparent; /*透明*/
  border: 1px solid #aaa;
  border-radius: 3px;
  padding: 13px 14px;
  width: 620px;
  font-size: 1.2rem;
  font-family: "Microsoft soft";
}
```

`input` 带有许多伪类选择器，使用这些伪类或伪元素可以实现许多有趣的功能。

```css
/*聚焦调整边框颜色*/
.form-control:focus {
  border-color: #66afe9;
  outline: 0;
  -webkit-box-shadow: inset 0 1px 1px rgba(0,0,0,.075),0 0 8px rgba(102,175,233,.6);
  box-shadow: inset 0 1px 1px rgba(0,0,0,.075),0 0 8px rgba(102,175,233,.6)
}
```

## `textarea`  元素

`textarea` 元素用于输入大段文字：

```html
<label for="story">Tell us your story:</label>
<textarea id="story" name="story" rows="5" cols="33">
It was a dark and stormy night...
</textarea> 
```

他有两个属性：`rows` 与 `cols` 设置行数与列数。

## `select` 下拉选择菜单

```html
<label for="pet-select">Choose a pet:</label>
<select name="pets" id="pet-select">
    <option value="">--Please choose an option--</option>
    <option value="dog">Dog</option>
    <option value="cat">Cat</option>
    <option value="hamster">Hamster</option>
    <option value="parrot">Parrot</option>
    <option value="spider">Spider</option>
    <option value="goldfish">Goldfish</option>
</select>
```

