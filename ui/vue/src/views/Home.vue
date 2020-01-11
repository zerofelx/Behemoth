<template>
    <div>
        <h1>Soy el home!</h1>
        <input type="text" v-model="nombre">
        <input type="button" value="Abrir Google" @click="SendEvent()">
        <hora />
    </div>
</template>

<script>
import Hora from '../components/Hora';
import gevent from '../setb.js'

export default {
    name: 'Home',
    data() {
        return {
            nombre: ''
        }
    },
    components: {
        Hora
    },
    methods: {
        SendEvent() {
            gevent.Evento(this.nombre + "|" + 'testeo')
        }
    }
}
</script>