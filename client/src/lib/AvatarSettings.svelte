<script>
  import cn from "classnames";
  import AvatarEditor from "./AvatarEditor.svelte";
  import { Card, CardTitle } from "./base";
  import { avatarSpec } from "../store";

  export let avatar;
</script>

<div class={cn(Card, "overflow-hidden", $$props.class)}>
  <div class={CardTitle}>Avatar</div>
  <AvatarEditor
    size={$avatarSpec.size}
    palette={$avatarSpec.palette}
    avatarData={avatar}
    class="mb-8"
  />
</div>
