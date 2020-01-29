let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");

function Evento(v) {
    if (v != '') {
        ws.send(JSON.stringify({
            event: "Chrome",
            variables: v
        }))
    } else {
        alert('Error')
    }
}

function Login(user, password, ToggleShow, Loading) {
    console.log('Iniciando Login')
    if (user != '' && password != '') {
        var v = user + '|' + password
        ws.send(JSON.stringify({
            event: "Login",
            variables: v
        }))
    } else {
        Loading()
        ToggleShow("Los campos de usuario y contraseña están vacíos!")
    }
}

export default {
    Evento,
    Login
}
