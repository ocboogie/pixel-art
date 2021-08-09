<script>
  import Navbar from "../lib/Navbar.svelte";
  import { redirect, Router, url } from "@roxi/routify";
  import { me } from "../store";
  import axios from "../axios";
  import { failure } from "../toasts";

  me.set(JSON.parse(localStorage.getItem("me")) || null);

  me.subscribe((me) => {
    localStorage.setItem("me", JSON.stringify(me));
  });

  axios.interceptors.response.use(
    (response) => response,
    (error) => {
      const { response } = error;

      if (error.response && error.response.status === 401) {
        me.set(null);
        $redirect("/login");
      }
      failure(error.response.data.message);
      return Promise.reject(error);
    }
  );
</script>

<main class="container m-auto text-center p-8">
  <slot />
</main>
<!-- The reason the navbar is below the main container is so the dropdown  -->
<!-- in the navbar won't be covered by elements in the main container -->
<Navbar />

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  @tailwind variants;

  #app {
    flex-direction: column-reverse;
    display: flex;
  }
</style>
