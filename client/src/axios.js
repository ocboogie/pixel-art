import axios from "axios";

const instance = axios.create({
  withCredentials: true,
  baseURL: import.meta.env.VITE_API_URL,
});

instance.interceptors.response.use(
  (response) => {
    if (response.data?.data) {
      response.data = response.data.data;
    }

    return response;
  },
  (error) => {
    if (error.response.data?.error) {
      error.response.data = error.response.data.error;
    }

    return Promise.reject(error);
  }
);

export default instance;
