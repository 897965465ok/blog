
import HelloWorld from './HelloWorld';
export default function Event() {

    const handleClick = () => {
        console.log('点击事件');
    }
    return (
        // 把函数传过去
        <HelloWorld event={() => { handleClick() }}>
        </HelloWorld>
    )
}

