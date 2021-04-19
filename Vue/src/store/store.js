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
    GET_PICTURES(state, result) {
      
      state.pictures= result.data.hits.map(url=> url.largeImageURL)
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
    async getPictures(context, result) {
      context.commit("GET_PICTURES", result)
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