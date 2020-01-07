function GetHour() {
    return fetch('http://localhost/hour')
        .then(res => res.json())
}

export {
    GetHour
}