import Vue from 'vue'
import App from "../App.vue"
import VueRouter from 'vue-router'
import userAuth from "../components/Userauth.vue"
import logIn from "../components/Login.vue";
import signUp from "../components/Signup.vue";
import mainApp from "../components/Mainapp.vue";
import store from "../store";

Vue.use(VueRouter)

const routes = [

  {
    path: '/',
    name: 'App',
    component: App
  },
  {
    path: '/user',
    name: "User",
    component: userAuth
  },
  {
    path: "/login",
    name: "Login",
    component: logIn
  },
  {
    path: "/signup",
    name: "SignUp",
    component: signUp
  },
  { path: '*' },
  {
    path: "/app",
    name: "MainApp",
    component: mainApp
  },

]

const router = new VueRouter({
  mode: 'history',
  routes
})

router.beforeEach((to, from, next) => {
  setTimeout(() => {
    if (to.path == "/") {
      if (store.getters.authStatus) {
        next("/app")
      } else {
        next("/user")
      }
    } else if (to.path == "/user" || to.path == "/login" || to.path == "/signup") {
      if (!store.getters.authStatus) {
        next()
      } else {
        next("/app")
      }
    } else if (to.path == "/app") {
      if (store.getters.authStatus) {
        next()
      } else {
        next("/user")
      }
    } else {
      next({ name: "not" })
    }
  }, 500)

})
export default router
