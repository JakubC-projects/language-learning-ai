import { createApp } from 'vue'
import App from './App.vue'
import { ensureIsLoggedIn } from './auth.ts'
import "./style/index.css"


async function load() {
    console.log("loading")
    if(!await ensureIsLoggedIn()){
        return;
    }
    console.log("logged in")
    createApp(App).mount('#app')
}

load()
