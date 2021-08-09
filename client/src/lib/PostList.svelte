<script>
  import cn from "classnames";
  import Post from "./Post.svelte";

  export let posts;
  export let gallery = false;
</script>

<div
  class={cn(
    gallery && "grid grid-cols-3 gap-3",
    gallery || "flex flex-col space-y-6"
  )}
>
  {#each posts as post}
    <Post {post} frameless={gallery} />
  {/each}
</div>
