let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");

function GetEvent() {
    ws.onmessage = (message) => {
        let obj = JSON.parse(message.data)

        console.log(obj.event);
        console.log(obj.variables);
        console.log(obj.AtrNameInFrontend)
    }
}

function Login(router, ToggleShow, FinishLoading) {
    ws.onmessage = (message) => {
        let obj = JSON.parse(message.data)

        if (obj.logged == true) {
            FinishLoading()
            router.push('/')
        } else {
            FinishLoading()
            ToggleShow("El usuario o la contraseña son inválidos")
        }
    }
}

export default {
    GetEvent,
    Login
}