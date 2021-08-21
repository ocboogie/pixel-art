<script>
  import { me } from "../store";
  import axios from "../axios";
  import { success } from "../toasts";
  import Loading from "../lib/Loading.svelte";
  import { goto } from "@roxi/routify";
  import { onMount } from "svelte";

  onMount(async () => {
    await axios.post("/auth/logout");
    me.set(null);
    $goto("/login");
    success("Logged out");
  });
</script>

<Loading>Logging out...</Loading>
