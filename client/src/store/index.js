import Vue from "vue";
import Vuex from "vuex";
import axios from "../plugins/axios";

Vue.use(Vuex);

const loggedIn = Boolean(localStorage.getItem("loggedIn"));

export default new Vuex.Store({
  state: {
    auth: {
      loggedIn
    },
    posts: []
  },
  mutations: {
    addPosts(state, posts) {
      state.posts = state.posts.concat(posts);
    },
    loggedIn(state) {
      Vue.set(state.auth, "loggedIn", true);
    },
    loggedOut(state) {
      Vue.set(state.auth, "loggedIn", false);
    }
  },
  actions: {
    async loadPosts(context) {
      const { data: posts } = await axios.get("/posts", {
        params: {
          limit: 10
        }
      });
      context.commit("addPosts", posts);
    },
    async signUp(context, user) {
      try {
        await axios.post("/auth/signUp", user);
      } catch (error) {
        throw error;
      }
      // The "true" here can be any string. It just needs to be truthy.
      localStorage.setItem("loggedIn", "true");
      context.commit("loggedIn");
    },
    async login(context, credentials) {
      try {
        await axios.post("/auth/login", credentials);
      } catch (error) {
        throw error;
      }
      // The "true" here can be any string. It just needs to be truthy.
      localStorage.setItem("loggedIn", "true");
      context.commit("loggedIn");
    },
    async logout(context) {
      try {
        await axios.post("/auth/logout");
      } catch (error) {
        throw error;
      }
      localStorage.removeItem("loggedIn");
      context.commit("loggedOut");
    }
  },
  modules: {}
});
