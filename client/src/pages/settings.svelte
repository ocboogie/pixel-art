<script>
  import cn from "classnames";
  import AvatarSettings from "../lib/AvatarSettings.svelte";
  import UserSettings from "../lib/UserSettings.svelte";
  import { ButtonPrimary } from "../lib/base";
  import { me } from "../store";
  import axios from "../axios";
  import { success } from "../toasts";

  $: newMe = { ...$me };
  let avatarSettings;

  function update() {
    const diff = {};
    let shouldUpdate = false;

    newMe.avatar = avatarSettings.serialize();

    for (let field in newMe) {
      if (newMe[field] !== $me[field]) {
        diff[field] = newMe[field];
        shouldUpdate = true;
      }
    }

    if (shouldUpdate) {
      axios.patch("/me", diff);

      success("Successful updated");

      $me = newMe;
    }
  }
</script>

<div class="flex flex-col justify-center max-w-2xl m-auto sm:flex-row">
  <AvatarSettings
    class="sm:mr-4 mb-8 flex-1"
    avatar={newMe.avatar}
    bind:this={avatarSettings}
  />
  <UserSettings class="flex-1 mb-8" user={newMe} />
</div>
<button
  class={cn(ButtonPrimary, "w-32 block m-auto text-xl")}
  type="button"
  on:click={update}>Update</button
>
