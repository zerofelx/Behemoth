<template>
    <div>
        <h1>{{ Hour }}:{{ Minute }}</h1>
    </div>
</template>

<script>
import api from "../api.js";

export default {
    name: 'Hora',
    data() {
        return {
            Hour: 0,
            Minute: 0,
            Second: 0,
            data: {}
        }
    },
    created() {
        this.PrintHour()
        setInterval(() => {
            console.log(this.Second)
            this.PrintHour()
        }, 60000);
    },
    methods: {
        PrintHour() {
            return api.GetHour().then(res => {
                this.data = res
                this.Hour = this.data.Hora
                this.Minute = this.data.Minuto
                this.Second = parseInt(this.data.Segundo + "000")
            })
        }
    }
}
</script>