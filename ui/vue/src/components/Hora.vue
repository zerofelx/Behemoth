<template>
    <div>
        <h1>4 de la xd</h1>
        <h1>{{ Hour }}</h1>
    </div>
</template>

<script>
export default {
    name: 'Hora',
    data() {
        return {
            Hour: 0,
            Minute: 0,
            data: {}
        }
    },
    // created() {
    //     this.PrintHour()
    //     setInterval(() => {
    //         this.PrintHour()
    //     }, 60000);
    // },
    methods: {
        PrintHour() {
            return GetHour().then(res => {
                this.data = res
                this.Hour = this.data.Hour
                this.Minute = this.data.Minute
            })
        }
    }
}
</script>