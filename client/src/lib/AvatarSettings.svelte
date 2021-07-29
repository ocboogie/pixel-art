<script>
  import cn from "classnames";
  import AvatarEditor from "./AvatarEditor.svelte";
  import { ButtonPrimary, Card, CardTitle } from "./base";

  export let avatar;
</script>

<div class={cn(Card, "overflow-hidden", $$props.class)}>
  <div class={CardTitle}>Avatar</div>
  <AvatarEditor
    size={5}
    palette={["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#e74c3c"]}
    avatarData={avatar}
    class="mb-8"
  />
</div>
