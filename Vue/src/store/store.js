import Vue from 'vue'
import Vuex from 'vuex'
// import createPersistedState from "vuex-persistedstate"
Vue.use(Vuex)
const store = new Vuex.Store({
  state: {
    articlers: [],
    pictures: [],
    tags: [],
    similar: [],
    recommen: {},
    token: "",
    router: [],
    oauthInfo: {}
  },
  getters: {
    token(state) {
      return state.token
    },
    pictures(state) {
      return state.pictures
    },
    oauth(state) {
      return state.oauth
    }
  },

  mutations: {
    CHANGE_ALL(state, articlers) {
      state.articlers = articlers
    },
    GET_PICTURES(state, result) {
      state.pictures = result.data.hits.map(url => url.largeImageURL)
    },
    WALLHAVEN(state, result) {
      state.pictures = result
    },
    SET_RECOMMEN(state, recommen) {
      state.recommen = { ...recommen }
    },
    SET_SIMILAR(state, similar) {
      state.similar = similar
    },
    SAVE_OAUTH(state, oauth) {
      state.oauthInfo = oauth
    }
  },
  actions: {
    async changeAll(context, data) {
      context.commit("CHANGE_ALL", data)
    },
    async getPictures(context, result) {
      context.commit("GET_PICTURES", result)
    },
    async setRecommen(context, recommen) {
      context.commit("SET_RECOMMEN", recommen)
    },
    async setSimilar(context, similar) {
      context.commit('SET_SIMILAR', similar)
    },
    async wallhaven(context, result) {
      context.commit('WALLHAVEN', result)
    },
    async saveOauth(context, oauth) {
      context.commit('SAVE_OAUTH', oauth)
    }

  },
  plugins: [window.createPersistedState({
    storage: window.sessionStorage
  })]

})
export default store