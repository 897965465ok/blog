
export function useWebSocket(handleMessage) {
    const ws = new WebSocket('ws://localhost:3800/ws')
    const init = () => {
        ws.addEventListener('open', handleOpen, false)
        ws.addEventListener('close', handleClose, false)
        ws.addEventListener('error', handleError, false)
        ws.addEventListener('message', handleMessage, false)
    }
    function handleOpen(e) {
        console.log('webSocket 打开', e);
        console.log(ws.readyState)
    }
    function handleClose(e) {
        console.log('webSocket 关闭', e);
        console.log(ws.readyState)
    }
    function handleError(e) {
        console.log('webSocket 错误', e);
        console.log(ws.readyState)
    }
    init()
    return ws
}

// class CreateWbSocket {
    
//     constructor() {

//         ws = new WebSocket('ws://localhost:3800/ws')
//     }
// }