<script>
  import { me } from "../store";
  import axios from "../axios";
  import { goto } from "@roxi/routify";
  import { onMount } from "svelte";
  import Loading from "../lib/Loading.svelte";

  onMount(async () => {
    await axios.post("/auth/logout");
    me.set(null);
    $goto("/login");
  });
</script>

<Loading>Logging out...</Loading>
