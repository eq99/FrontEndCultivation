# HTML 标签之表单

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

一个表单由一个 `form` 元素组织，其中包含若干个输入标签，例如：`input`，一个提交按钮 `button`。

## `input` 元素

`input` 元素通过 `type` 属性设置输入类型，不同的 type 有不同的功能。

例如：

```html
<input type="password">
```

在输入时会隐藏密码。

例如：

```html
<input type="file">
```

会弹出文件上传面板。

更多类型请参考：https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input

`input` 元素除了之前见到过的 `type, id, name` 属性外，还有

-  `value` ：输入框的值
- `disabled`：输入框是否被禁用

更所的属性及其含义参考：https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input#attributes

## `label` 元素

`label` 元素用于给`input` 元素添加提示信息，当点击 label 元素时，会激活相应输入框：

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

使用方法是把 `label` 的 `for` 属性指向 `input` 元素的 `id` 。

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

