<script>
  import { onMount } from "svelte";
  import { deserialize, random } from "../utils/avatar";

  export let size;
  export let avatarData = null;
  export let editable = false;
  export let color = null;

  let cells = [];
  // Draw states
  // null: not drawing
  // true: drawing
  // false: removing
  let dragDrawState = null;

  $: cellColor = color;
  $: filteredCells = cells
    .map((cell, pos) => [cell, pos])
    .filter(([cell]) => cell || editable);
  $: {
    if (!avatarData) {
      randomize();
    } else {
      [size, cells, cellColor] = deserialize(avatarData);
    }
  }

  onMount(() => {
    window.addEventListener("mouseup", resetDragDrawState);
    return () => {
      window.removeEventListener("mouseup", resetDragDrawState);
    };
  });

  export function randomize() {
    cells = random(size);
  }

  function resetDragDrawState() {
    dragDrawState = null;
  }

  function cellClicked(event, pos) {
    if (event.button === 0) {
      dragDrawState = !cells[pos];
      toggle(pos);
    }
  }

  function cellHovered(event, pos) {
    if (dragDrawState !== null && event.buttons === 1) {
      cells[pos] = dragDrawState;
      cells = cells;
    }
  }

  function toggle(pos) {
    cells[pos] = !cells[pos];
    cells = cells;
  }
</script>

<svg
  xmlns="http://www.w3.org/2000/svg"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  preserveAspectRatio="xMinYMin"
  viewBox={`0 0 ${size} ${size}`}
  class={$$props.class}
>
  {#each filteredCells as [cell, pos]}
    <rect
      name={pos}
      class="cell"
      x={pos % size}
      y={Math.floor(pos / size)}
      width="1"
      height="1"
      fill={cellColor}
      fill-opacity={cell ? 1 : 0}
      on:mousedown={editable ? (event) => cellClicked(event, pos) : null}
      on:mouseover={editable ? (event) => cellHovered(event, pos) : null}
    />
  {/each}
</svg>

<style lang="postcss">
  svg {
    user-select: none;
    filter: drop-shadow(0px 3px 6px rgba(0, 0, 0, 0.16));
  }
</style>
