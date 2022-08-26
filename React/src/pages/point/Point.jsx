import React from 'react';

export default class HelloWorld extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            styleData: { color: 'red', 'fontSize': "16px" },
            isHead: true,
            className: 'title'
        };

        // 为了在回调中使用 `this`，这个绑定是必不可少的
        // this.handleClick = this.handleClick.bind(this);
    }
    // 放入类的原型
    handleClick() {
        // console.log(this.state)
        console.log(this.state)
    }

    render() {
       this.handleClick()
       // 这里还是函数作用域
        return <>
            <button onClick={
                this.handleClick()
            }>按钮</button>

            {/* <button onClick={console.log(this)}>按钮2</button> */}
        </>
    }
}
