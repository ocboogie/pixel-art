import Vue from "vue";
import Vuex from "vuex";
import axios from "../plugins/axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    me: null,
    posts: []
  },
  mutations: {
    addPosts(state, posts) {
      state.posts = state.posts.concat(posts);
    },
    modifyMe(state, meDiff) {
      state.me = { ...state.me, ...meDiff };
    },
    loggedIn(state, user) {
      state.me = user;
    },
    loggedOut(state) {
      state.me = null;
    }
  },
  actions: {
    async fetchMe(context) {
      const { data: user } = await axios.get("/me");
      context.commit("loggedIn", user);
    },
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
      context.dispatch("fetchMe");
    },
    async login(context, credentials) {
      try {
        await axios.post("/auth/login", credentials);
      } catch (error) {
        throw error;
      }
      context.dispatch("fetchMe");
    },
    async logout(context) {
      try {
        await axios.post("/auth/logout");
      } catch (error) {
        throw error;
      }
      context.commit("loggedOut");
    },
    async updateProfile(context, user) {
      const body = {};
      if (user.name != context.state.me) {
        body.name = user.name;
      }
      if (user.avatar != context.state.me.avatar) {
        body.avatar = user.avatar;
      }

      await axios.patch("/me", body);

      context.commit("modifyMe", body);
    }
  },
  getters: {
    loggedIn(state) {
      return Boolean(state.me);
    }
  },
  modules: {}
});
