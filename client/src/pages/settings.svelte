<script>
  import cn from "classnames";
  import AvatarSettings from "../lib/AvatarSettings.svelte";
  import UserSettings from "../lib/UserSettings.svelte";
  import { ButtonPrimary } from "../lib/base";

  const user = {
    id: "422cb2d3-7213-4734-aab2-5f31f0645925",
    name: "test",
    createdAt: "2019-11-30T00:31:22.414657Z",
    avatar: "1101101101110111101111011#e74c3c",
  };
</script>

<div class="flex flex-col justify-center max-w-2xl m-auto sm:flex-row">
  <AvatarSettings class="sm:mr-4 mb-8 flex-1" avatar={user.avatar} />
  <UserSettings class="flex-1 mb-8" {user} />
</div>
<button class={cn(ButtonPrimary, "w-32 block m-auto text-xl")} type="button"
  >Update</button
>
