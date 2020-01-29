<template>
    <div :style="BGPath" id="LoginContainer">
        <single-popup 
            :Message="LoginIncorrect"
            :Show="Show"
            @change-show="ToggleShow"
        >
        </single-popup>
        <Loading
            :Loading="Loading"
        ></Loading>
        <div id="Login">
            <label for="Username">Nombre o Correo</label>
            <input name="Username" placeholder="Username" type="text" v-model="user"/>
            <label for="Password">Contraseña</label>
            <input name="Password" placeholder="Password" type="password" v-model="password"/>
            <input type="button" value="Iniciar Sesión" @click="UserLogin(user, password)">
        </div>
    </div>
</template>

<script>
import api from '../api.js'
import setb from '../setb.js'
import getb from '../getb.js'
import SinglePopup from '../components/alerts/SinglePopup'
import Loading from '../components/alerts/Loading'

export default {
    name: 'Login',
    components: {
        SinglePopup,
        Loading
    },
    data() {
        return {
            user: '',
            password: '',
            BGPath: 'background-image: url("http://localhost/sources/images/LoginBG.jpg");',
            LoginIncorrect: "Error",
            Show: false,
            Loading: false
        }
    },
    methods: {
        UserLogin(u, p) {
            this.ToggleLoading()
            setb.Login(u, p, this.ToggleShow, this.ToggleLoading)
            getb.Login(this.$router, this.ToggleShow, this.ToggleLoading)
        },
        ToggleShow(message) {
            this.LoginIncorrect = message
            this.Show = !this.Show
        },
        ToggleLoading() {
            this.Loading = !this.Loading
        }

    }
}
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