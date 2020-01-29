<template>
<div v-show="Loading" id="PopupBG">
    <div id="Popup">
        <div id="content">
        <p>Cargando</p>
        </div>
    </div>
</div>
</template>

<script>
export default {
    name: 'loading',
    props: ['Loading']
}
</script>

<style  scoped>
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

    @keyframes fadeIn{
        from {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }
</style>