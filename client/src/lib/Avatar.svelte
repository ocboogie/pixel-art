<script>
  import { onMount } from "svelte";

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
      const [cellsString, color] = avatarData.split("#");
      size = Math.ceil(Math.sqrt(cellsString.length));
      cells = [...cellsString].map((cell) => cell === "1");
      cellColor = `#${color}`;
    }
  }

  onMount(() => {
    window.addEventListener("mouseup", resetDragDrawState);
    return () => {
      window.removeEventListener("mouseup", resetDragDrawState);
    };
  });

  export function randomize() {
    const sizeHalf = Math.ceil(size / 2);
    for (let y = 0; y < size; y += 1) {
      for (let x = 0; x < sizeHalf; x += 1) {
        const pos = y * size + x;
        const mirroredPos = y * size + (size - 1 - x);
        const cell = Math.random() >= 0.5;
        cells[pos] = cell;
        if (mirroredPos !== pos) {
          cells[mirroredPos] = cell;
        }
        cells = cells;
      }
    }
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

  export function intoData() {
    let data = cells.map((cell) => (cell ? "1" : "0")).join("");
    data += cellColor;
    return data;
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
