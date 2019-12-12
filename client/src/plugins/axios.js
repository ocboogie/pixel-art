import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";
import store from "../store";

const instance = axios.create({
  withCredentials: true,
  baseURL: process.env.VUE_APP_API_URL
});

instance.interceptors.response.use(
  response => {
    if (response.data && response.data.data) {
      response.data = response.data.data;
    }
    return response;
  },
  error => {
    const { response } = error;
    if (response && response.status === 401) {
      store.dispatch("loggedOut");
    }
    return Promise.reject(error);
  }
);

Vue.use(VueAxios, instance);

export default instance;
