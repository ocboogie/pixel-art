<script>
  import cn from "classnames";
  import Avatar from "../lib/Avatar.svelte";
  import AvatarEditorTool from "../lib/AvatarEditorTool.svelte";

  export let avatarData = null;
  export let palette;
  export let size;

  let avatar;
  let selectedColor = palette[Math.floor(Math.random() * palette.length)];
  function randomize() {
    avatar.randomize();
  }
  export function serialize() {
    return avatar.serialize();
  }
</script>

<div class="p-4 relative border-b">
  <Avatar
    editable
    {avatarData}
    {size}
    color={selectedColor}
    bind:this={avatar}
  />
</div>
<div class={cn("rounded mt-8 h-24 w-24 m-auto", $$props.class)}>
  <AvatarEditorTool
    class="tool"
    {palette}
    bind:selectedColor
    on:randomize={randomize}
  />
</div>
