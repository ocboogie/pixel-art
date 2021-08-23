<script>
  import { onMount } from "svelte";
  import axios from "../axios";
  import Loading from "../lib/Loading.svelte";
  import Profile from "../lib/Profile.svelte";
  import { me } from "../store";

  onMount(async () => {
    if (!$me) {
      me.set((await axios.get("/me")).data);
    }
  });
</script>

<div class="max-w-lg m-auto">
  {#if $me}
    <Profile userId={$me.id} />
  {:else}
    <Loading>Loading...</Loading>
  {/if}
</div>
