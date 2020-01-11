<template>
    <div :style="BGPath" id="LoginContainer">
        <div id="Login">
            <label for="Username">Nombre o Correo</label>
            <input name="Username" placeholder="Username" type="text" v-model="user"/>
            <label for="Password">Contraseña</label>
            <input name="Password" placeholder="Password" type="password" v-model="password"/>
            <input type="button" value="send" @click="UserLogin(user, password)">
        </div>
    </div>
</template>

<script>
import api from '../api.js'
import setb from '../setb.js'
import getb from '../getb.js'

export default {
    name: 'Login',
    data() {
        return {
            user: '',
            password: '',
            BGPath: 'background-image: url("http://localhost/sources/images/LoginBG.jpg");'
        }
    },
    methods: {
        UserLogin(u, p) {
            setb.Login(u, p)
            getb.Login(this.$router)
        }
    }
}
fetch('http://localhost/Logged').then(res => res.json())
</script>

<style scoped>
    #LoginContainer {
        width: 100%;
        height: 100vh;
        background-size: cover;
    }
    #Login * {
        outline: none;
    }
    #Login {
        width: 320px;
        height: 420px;
        background: rgba(0, 0, 0, 0.61);
        color: white;
        top: 50%;
        left: 50%;
        position: absolute;
        transform: translate(-50%, -50%);
        box-sizing: border-box;
        padding: 70px 30px;;
    }
    #Login input {
        width: 100%;
        margin: 0 0 20px 0;
    }
    #Login input[type="text"],
    #Login input[type="password"] {
        border: none;
        border-bottom: 0.5px solid white;
        background: transparent;
        color: white;
        padding: 5px 2px;
    }
    #Login input[type="button"] {
        border: none;
        outline: none;
        height: 40px;
        background: #b80f22;
        color: white;
        font-size: 18px;
        border-radius: 20px;
    }
    #Login label {
        margin: 0 0 2px 0;
        padding: 0 2px;
        font-weight: 600;
        font-size: 20px;
        display: block;
    }
    #Login input[type="button"]:hover {
        cursor: pointer;
        background: #bb2738;
    }
    #Login input[type="button"]:active {
        background: #b80f22;
    }
</style>