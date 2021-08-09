<script>
  import Profile from "../../lib/Profile.svelte";
  import Loading from "../../lib/Loading.svelte";
  import { onMount } from "svelte";
  import axios from "../../axios";

  export let userId;

  let user;

  onMount(async () => {
    user = (await axios.get(`/users/${userId}`)).data;
  });
</script>

{#if !user}
  <Loading>Loading user...</Loading>
{:else}
  <Profile {user} />
{/if}
