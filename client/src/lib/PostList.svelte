<script>
  import cn from "classnames";
  import Post from "./Post.svelte";

  export let posts;
  export let gallery = false;
</script>

<div
  class={cn(
    gallery && "grid grid-cols-3 gap-3",
    gallery || "flex flex-col space-y-6",
    $$props.class
  )}
>
  {#each posts as post}
    <Post {post} frameless={gallery} />
  {/each}
</div>
