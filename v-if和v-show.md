v-if和v-show:

v-if条件为false时，压根不会有对应的元素在dom中，

v-show为false时，仅仅将元素的display属性设置为none而已

+ 当需要在显示与隐藏之间频繁切换时，使用v-show
+ 当只有一次切换时，使用v-if



组件的key属性

在使用v-for时，应给对应的组件或元素添加一个：key属性

为什么：跟vue的虚拟DOM的Diff算法有关

