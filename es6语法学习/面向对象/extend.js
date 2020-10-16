//类的继承

const extend = require("extend");

// //1.父类
// class Father {
//     constructor() { }
//     money() {
//         console.log(1000);
//     }
// }
// //2.子类继承父类
// class Son extends Father {

// }
// //实例化一个子类

// var son = new Son();
// son.money();



class Father {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }
    sum() {
        console.log(this.x + this.y);
    }
}


class Son extends Father {

}
var son = new Son(100, 100);
son.sum();


