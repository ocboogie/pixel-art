<script>
  import { createEventDispatcher } from "svelte";
  import cn from "classnames";
  import ColorWheel from "./ColorWheel.svelte";

  const dispatch = createEventDispatcher();

  export let selectedColor;
  export let palette;
</script>

<div class={cn("tool-button-container", $$props.class)}>
  <ColorWheel
    class="color-wheel"
    bind:selectedColor
    {palette}
    radius="30"
    strokeWidth="60"
    emptySpace="90"
    gap="5"
  />
  <div class="tool-button" on:click={() => dispatch("randomize")}>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-6 w-6 icon"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
      />
    </svg>
  </div>
</div>

<style lang="postcss">
  .tool-button-container {
    pointer-events: auto;
    display: inline-block;
    box-shadow: 0 7px 7px rgba(black, 0.15);
    position: relative;
  }
  .tool-button {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    top: 0;
    width: 100%;
    height: 100%;
    border-radius: 100%;
    background-color: rgb(59, 130, 246);
    transition: box-shadow 0.2s;
    cursor: pointer;
    padding: 0 10px;
  }
  :global(.color-wheel) {
    width: 100%;
    height: 100%;
    transition: transform 0.2s;
    transform: scale(1.5);
  }
  .tool-button-container:hover :global(.color-wheel) {
    transform: scale(2);
  }
  .icon {
    width: 100%;
    height: 100%;
    stroke: white;
    opacity: 0.75;
  }
  /* .tool-button :global(.icon) {
    width: 100%;
    height: 100%;
    fill: white;
    opacity: 0.5;
  } */
  /* .tool-button.selected,
  .tool-button.hover-select:hover {
    background-color: lighten($tool-button-color, 15%);
    .icon {
      opacity: 1;
    }
  } */
</style>
