import Vue from 'vue'
import axios from 'axios'
import App from './App.vue'
import './plugins/element.js'
import router from './router'

Vue.config.productionTip = false;
Vue.prototype.$http = axios;

// eslint-disable-next-line
const app = new Vue({
    el: '#app',
    router: router,
    render: h => h(App),
});

