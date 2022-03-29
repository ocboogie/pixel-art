<script>
  import cn from "classnames";

  export let data;
</script>

<img
  class={cn("art-renderer", $$props.class)}
  src={`data:image/gif;base64,${data}`}
/>

<style lang="postcss">
  .art-renderer {
    image-rendering: -moz-crisp-edges;
    image-rendering: -webkit-crisp-edges;
    image-rendering: pixelated;
    image-rendering: crisp-edges;
  }
</style>
