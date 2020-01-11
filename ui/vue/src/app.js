import Vue from 'vue'
import App from './App'
import router from './router'

window.onload = function () {
  var topic = document.createElement("div");
  topic.id = 'vue-app';
  document.body.insertBefore(topic, document.body.childNodes[0])  

  new Vue({
    router,
    el: '#vue-app',
    components: { App },
    template: '<App/>'
  })
}
