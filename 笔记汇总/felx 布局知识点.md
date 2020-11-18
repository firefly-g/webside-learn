felx 布局之-----父项属性

- flex-direction :  主轴方向（<u>元素根据主轴排列的</u>）

  > row                        默认值（从左到右）
  >
  > row-reverse         从右到左
  >
  > colum    				默认值（从上到下）
  >
  > colum-reverse     从下到上

- justify-content :  <u>主轴</u>上子元素的排列方式

  > flex-start                默认值（从头开始，主轴为X轴时，则为从左到右）
  >
  > flex-end   			  从尾部开始排列
  >
  > center    				在主轴 居中对齐（主轴为X轴时，则为水平居中）
  >
  > space-around       平分剩余空间
  >
  > space-between     <u>先两边贴边  再平分剩余空间（重要）</u>

- flex-wrap:子元素是否换行

- align-content:     侧轴上子元素的排列方式（多行--子项出现换行情况下）

  > flex-start                默认值（从头开始，侧轴为X轴时，则为从左到右）
  >
  > flex-end   			  从尾部开始排列
  >
  > center    				在主轴 居中对齐（主轴为X轴时，则为水平居中）
  >
  > space-around       平分剩余空间
  >
  > space-between     <u>先两边贴边  再平分剩余空间（重要）</u>
  >
  > stretch					设置子项元素高度平分父元素高度

- align-items：      侧轴上子元素的排列方式（单行）

  > flex-start
  >
  > flex-end
  >
  > center                     侧轴方向上的居中对齐
  >
  > stretch    				拉伸 但是子盒子不要给高度

- flex-flow:复合属性，相当于提示设置了flex-direction和flex-wrap属性





felx 布局之-----子项属性

-  flex								子项目占据的份数

- align-self           			 控制子项自己在侧轴的排列方式

- order								定义子项的排列顺序（前后顺序）数值越小越靠前 默认为0



https://github.com/firefly-g/daily-learn-record.git