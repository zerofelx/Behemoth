<template>
  <main>
    <hora />
  </main>
</template>

<script>
import Hora from './components/Hora.vue'
export default {
  name: 'App',
  components: {
    Hora
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Montserrat:400,600&display=swap');
  body {
    margin: 0;
    font-family: 'Montserrat'
  }

  .topic {
    color: #3a2c38;
  }
  a {
      text-decoration: none;
      color: rgb(139, 36, 36);
      outline: none;
  }
  a:hover {
      color: rgb(175, 44, 44);
  }
  p {
      margin: 0;
  }
  h1, h2, h3, h4, h5 {
      margin: 0;
  }
</style>