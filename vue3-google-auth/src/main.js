import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import gAuthPlugin from 'vue3-google-oauth2';
const app = createApp(App)
import { env } from "./dev.env";

app.use(gAuthPlugin, {
	clientId: env.GAuthClientId,
	scope: 'email',
	prompt: 'consent',
	fetch_basic_profile: false
})

app.mount('#app')
