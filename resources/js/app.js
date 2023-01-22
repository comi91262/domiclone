import './bootstrap.js';

import { createApp } from 'vue/dist/vue.esm-bundler';
import App from './App.vue';
import store from './store.js';

createApp(App).use(store).mount('#app')

