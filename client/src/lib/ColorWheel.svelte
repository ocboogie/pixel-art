<script>
  import { createEventDispatcher } from "svelte";

  export let selectedColor;
  export let palette;
  export let radius;
  export let strokeWidth;
  export let gap;
  export let emptySpace;

  let hovering = null;

  $: circumference = 2 * Math.PI * radius;
  $: circumferenceWithEmptySpace =
    circumference - circumference / (360 / emptySpace);
  $: circumferenceWithEmptySpaceInDegrees = 360 - emptySpace;
  $: strokeDashOffset =
    circumference - circumferenceWithEmptySpace / palette.length + gap;
  let globalOffset;
  $: {
    const gapOffsetRadians = ((gap / 2) * palette.length) / circumference;
    const gapOffset = gapOffsetRadians * (180 / Math.PI);
    const emptySpaceOffset = emptySpace / 2;

    globalOffset = 90 + emptySpaceOffset + gapOffset;
  }

  function circleTransformValue(index, selected = false, hovered = false) {
    const degreesForEachColor =
      circumferenceWithEmptySpaceInDegrees / palette.length;
    const angleOffset = degreesForEachColor * index + globalOffset;
    let scale = 1;
    if (selected) {
      scale = 1.3;
    } else if (hovered) {
      scale = 1.1;
    }
    return `rotate(${angleOffset}) scale(${scale})`;
  }
</script>

<svg height="160" width="160" viewBox="0 0 160 160" class={$$props.class}>
  {#each palette as color, index}
    <circle
      class="color"
      name={index.toString()}
      cx="80"
      cy="80"
      r={radius}
      stroke={color}
      stroke-width={strokeWidth}
      stroke-dasharray={circumference}
      stroke-dashoffset={strokeDashOffset}
      transform={circleTransformValue(
        index,
        color === selectedColor,
        index === hovering
      )}
      on:click={() => (selectedColor = color)}
      on:mouseenter={() => (hovering = index)}
      on:mouseleave={() => (hovering = null)}
    />
  {/each}
</svg>

<style lang="postcss">
  .color {
    pointer-events: stroke;
    fill: transparent;
    transform-origin: center center;
    transition: transform 0.2s;
    cursor: pointer;
  }
</style>
