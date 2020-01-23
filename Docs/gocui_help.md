### 关于g和v的颜色

- 对于`g`, `DefaultColor` = 0，它用在 `bg`和`fg`上效果不一样，`fg`是`white`的效果，`bg`是`透明`效果

- 如果使用非`DefaultColor`，在 `g`上，`bg`he`fg`分别定义了 ui背景和 view的线框颜色；在`view`上，分别定义了上面的**文字**的背景和前景色

### 关于 ColorNormal和 Color256

- Normal只能用[color256](./color256.png) 的前8个颜色，而 256模式可以使用全部, 一般来说256使用前16个（后8个是前8个的加亮色)，然后进阶一些再挑选其他的单个颜色