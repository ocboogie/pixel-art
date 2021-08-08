<script>
  import { me } from "../store";
  import axios from "../axios";
  import { goto } from "@roxi/routify";
  import { onMount } from "svelte";

  onMount(async () => {
    await axios.post("/auth/logout");
    me.set(null);
    $goto("/login");
  });
</script>

<h1 class="text-3xl">Logging out...</h1>
