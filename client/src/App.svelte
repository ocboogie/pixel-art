<script>
  import { SvelteToast } from "@zerodevx/svelte-toast";
  import { onMount } from "svelte";
  import { Router } from "@roxi/routify";
  import { routes } from "../.routify/routes";
  import { avatarSpec } from "./store";
  import axios from "./axios";

  onMount(async () => {
    avatarSpec.set((await axios.get("/avatar/spec")).data);
  });
</script>

<SvelteToast />
<Router {routes} />
