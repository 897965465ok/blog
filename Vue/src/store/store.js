import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
const store = new Vuex.Store({
  state: {
    articlers: [],
    pictures: [],
    tags: [],
    similar: [],
    recommen: {},
    token: ""
  },
  getters: {
    token(state) {
      return state.token
    },
    pictures(state) {
      return state.pictures
    }
  },

  mutations: {
    CHANGE_ALL(state, articlers) {
      state.articlers = articlers
    },
    GET_PICTURES(state, url) {
      state.pictures.push(url)
    },
    SET_RECOMMEN(state, recommen) {
      state.recommen = { ...recommen }
    },
    SET_SIMILAR(state, similar) {
      state.similar = similar
    }
  },
  actions: {
    async changeAll(context, data) {
      context.commit("CHANGE_ALL", data)
    },
    async getPictures(context, url) {
      context.commit("GET_PICTURES", url)
    },
    async setRecommen(context, recommen) {
      context.commit("SET_RECOMMEN", recommen)
    },
    async setSimilar(context, similar) {
      context.commit('SET_SIMILAR', similar)
    }

  }

})
export default store