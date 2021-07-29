<script>
  import { onMount } from "svelte";
  import cn from "classnames";
  import decodeArtData from "../utils/decodeArtData";

  export let data;

  let canvas;
  let ctx;
  $: info = decodeArtData(data);

  $: if (data && ctx) {
    draw();
  }

  onMount(() => {
    ctx = canvas.getContext("2d");
  });
  function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    info.pixels.forEach((colorIndex, index) => {
      const x = index % info.width;
      const y = Math.floor(index / info.height);

      ctx.fillStyle = info.colors[colorIndex];
      ctx.fillRect(x, y, 1, 1);
    });
  }
</script>

<canvas
  class={cn("art-renderer", $$props.class)}
  bind:this={canvas}
  width={info.width}
  height={info.height}
/>

<style lang="postcss">
  .art-renderer {
    image-rendering: -moz-crisp-edges;
    image-rendering: -webkit-crisp-edges;
    image-rendering: pixelated;
    image-rendering: crisp-edges;
  }
</style>
