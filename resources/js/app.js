import './bootstrap.js';

import { createApp } from 'vue/dist/vue.esm-bundler';
import App from './App.vue';
import store from './store.js';
import Alpine from 'alpinejs';

window.Alpine = Alpine;
Alpine.start();
createApp(App).use(store).mount('#app')
 
