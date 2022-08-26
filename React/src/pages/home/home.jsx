import React from 'react';
class HelloWorld extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            time: null,
            timer: null
        }
    }

    componentDidMount() {
        let timer = setInterval(() => {
            // 函数的方式
            this.setState((state,props) => { 
                timer : timer
            })
            this.setState({ time: new Date().toLocaleTimeString() })
        }, 500)
        this.setState(() => { timer })
    }

    componentDidUpdate() {
        console.log("数据更新")
    }
    
    componentWillUnmount() {
        console.log("当组件从 DOM 中移除时会调用如下方法")
        clearInterval(this.state.timer)
        this.setState(() => { timer: null })
    }

    render() {
        return(<>
                <h1>{this.state.time}</h1>
                <div>{this.props.title}</div>
            </>)
    }
}
HelloWorld.defaultProps = {
    title: 'hello world',
    fuck: 'fuck thursday'
};
export default HelloWorld;
