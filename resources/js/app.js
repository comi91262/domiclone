
/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

require('./bootstrap');

window.Vue = require('vue').default;
window.Vue.use(require('vuex'));


/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

Vue.component('entry', require('./components/Entry.vue').default);
Vue.component('front', require('./components/Front.vue').default);


Vue.component('card', require('./components/Card.vue').default);
Vue.component('balloon', require('./components/Balloon.vue').default);
Vue.component('selection', require('./components/Selection.vue').default);

Vue.component('hand', require('./components/Hand.vue').default);
Vue.component('supply', require('./components/Supply.vue').default);
Vue.component('playarea', require('./components/PlayArea.vue').default);
Vue.component('trash', require('./components/Trash.vue').default);
Vue.component('public', require('./components/Public.vue').default);


Vue.component('debug', require('./components/Debug.vue').default);

Vue.component('log', require('./components/Log.vue').default);



import phase from './store/modules/phase'
import log from './store/modules/log'
import action_phase from './store/modules/action_phase'
import buy_phase from './store/modules/buy_phase'
import clean_phase from './store/modules/clean_phase'
import player from './store/modules/player'
import Vuex from 'vuex'


const store = new Vuex.Store({
  modules: {
      phase,
      log,
      player,
      action_phase,
      buy_phase,
      clean_phase,
  },
  strict: true,
})

const app = new Vue({
    el: '#app',
    store,
});

