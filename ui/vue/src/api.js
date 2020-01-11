function GetHour() {
    return fetch('http://localhost/PrintHour')
        .then(res => res.json())
}

function GetUser() {
    return fetch('http://localhost/Logged')
}

export default {
    GetHour,
    GetUser
}