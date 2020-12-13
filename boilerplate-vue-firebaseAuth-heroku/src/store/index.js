import Vue from 'vue'
import Vuex from 'vuex'
import localforage from "localforage"
var indexDB = localforage.createInstance({
  name: "users"
})

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    isAuth: false,
  },
  getters: {
    authStatus: state => {
      return state.isAuth;
    }
  },
  mutations: {
    authStatus: state => {
      indexDB.getItem("user").then(data => {
        if (data) {
          state.isAuth = true
          console.log(state.isAuth)
        } else {
          state.isAuth = false
          console.log(state.isAuth)
        }
      }).catch(error => {
        console.log(error);
        state.isAuth = false
      })
    },
    changeStatus: state => {
      console.log("change")
      state.isAuth = !state.isAuth;
      console.log("state" + state.isAuth)
    }
  },
  actions: {
    getAuthStatus: context => {
      context.commit("authStatus")
    },
    changeAuthStatus: context => {
      context.commit("changeStatus")
    }
  },
  modules: {
  }
})
