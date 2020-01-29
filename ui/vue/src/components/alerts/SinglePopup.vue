<template>
<div v-show="Show" id="PopupBG" :class="Show ? 'show' : 'hide'">
    <div :class="Show ? 'show' : 'hide'" id="Popup">
        <div id="content">
            <h3>{{ Message }}</h3>
            <input type="button" value="Cerrar" @click="EmitShow()">
        </div>
    </div>
</div>
</template>

<script>
export default {
    name: "SinglePopup",
    props: ['Message', 'Show'],
    methods: {
        EmitShow() {
            this.$emit('change-show')
        }
    }
}
</script>

<style scoped>
    .hide {
        color: red;
    }
    #PopupBG {
        height: 100vh;
        width: 100%;
        background-color: rgba(0, 0, 0, 0.58);
        position: absolute;
        z-index: 2;
    }
    #PopupBG.show {
        animation-duration: .5s;
        animation-name: fadeIn
    }
    #Popup {
        min-height: 25vh;
        width: 50%;
        background-color: white;
        color: tomato;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        z-index: 3;
        padding: 10px 20px;
        text-align: center;
        display: flex;
        align-items: center;
    }
    #Popup.show {
        animation-duration: .5s;
        animation-name: Scale
    }
    #content{
        width: 100%;
        display: block;
    }
    #content input[type="button"] {
        margin-top: 40px;
        border: none;
        outline: none;
        height: 30px;
        min-width: 30px;
        border-radius: 20px;
        background: #b80f22;
        color: white;
        font-size: 18px;
    }
    #content input[type="button"]:hover {
        cursor: pointer;
        background: #bb2738;
    }
    #content input[type="button"]:active {
        background: #b80f22;
    }

    #PopupBG.show {
        opacity: 1;
    }

    @keyframes fadeIn{
        from {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }
    @keyframes Scale {
        from {
            transform: translate(-50%, -50%) scale(0.5);
        }
        to {
            transform: translate(-50%, -50%) scale(1);
        }
    }
</style>