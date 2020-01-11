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

function Login(user, password) {
    console.log('Iniciando Login')
    if (user != '' && password != '') {
        var v = user + '|' + password
        ws.send(JSON.stringify({
            event: "Login",
            variables: v
        }))
    } else {
        alert('Error')
    }
}

export default {
    Evento,
    Login
}
