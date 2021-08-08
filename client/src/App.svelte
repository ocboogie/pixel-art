<script>
  import { SvelteToast, toast } from "@zerodevx/svelte-toast";
  import { Router } from "@roxi/routify";
  import { routes } from "../.routify/routes";
  import { me } from "./store";
  import axios from "./axios";
  import { failure } from "./toasts";

  axios.interceptors.response.use(
    (response) => response,
    (error) => {
      const { response } = error;

      if (error.response && error.response.status === 401) {
        me.set(null);
      }
      failure(error.response.data.message);
      return Promise.reject(error);
    }
  );
</script>

<SvelteToast />
<Router {routes} />
