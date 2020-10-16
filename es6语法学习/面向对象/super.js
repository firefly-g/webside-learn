//super关键字：用于访问和调用对象父类上的函数。可以调用父类的构造函数和普通函数。
class Father {
    constructor(x, y) {
        this.x = x;
        this.y = y;
        this.sum()
    }
    sum() {
        console.log(this.x + this.y);
    }

    say() {
        return '我是爸爸';
    }
}

class Son extends Father {
    constructor(x, y) {
        super(x, y);       //super调用了父类中的构造函数
    }

    say() {
        //console.log('我是儿子');
        console.log(super.say());
    }
}
var son = new Son(100, 100);
// son.sum();
son.say();
console.log(son);
