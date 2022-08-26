export default function HelloWorld(props) {

    const { event } = props

    const handleClick = () => {
        event()
    }
    return (
        <button onClick={handleClick}>hello world</button>
    );
}

