<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    xmlns:xlink="http://www.w3.org/1999/xlink"
    preserveAspectRatio="xMinYMin"
    :viewBox="`0 0 ${size} ${size}`"
  >
    <rect
      v-for="[cell, pos] in filteredCells"
      :key="pos"
      class="cell"
      :x="pos % size"
      :y="Math.floor(pos / size)"
      width="1"
      height="1"
      :fill="cellColor"
      :fill-opacity="cell ? 1 : 0"
      v-on="
        // Enable events when editable
        editable
          ? {
              mousedown: event => cellClicked(event, pos),
              mouseover: event => cellHovered(event, pos)
            }
          : {}
      "
    />
  </svg>
</template>

<script>
import Vue from "vue";
export default {
  props: {
    avatarData: {
      type: String,
      default: null
    },
    editable: {
      type: Boolean,
      default: false
    },
    color: {
      type: String,
      default: null
    }
  },
  data: () => ({
    size: 5,
    cells: [],
    cellColor: null,
    // Draw states
    // null: not drawing
    // true: drawing
    // false: removing
    dragDrawState: null
  }),
  computed: {
    filteredCells() {
      return this.cells
        .map((cell, pos) => [cell, pos])
        .filter(([cell]) => cell || this.editable);
    }
  },
  watch: {
    color: {
      immediate: true,
      handler(color) {
        this.cellColor = color;
      }
    },
    avatarData: {
      immediate: true,
      handler(avatarData) {
        if (!avatarData) {
          this.randomize();
          return;
        }
        const [cellsString, color] = avatarData.split("#");
        this.size = Math.ceil(Math.sqrt(cellsString.length));
        this.cells = [...cellsString].map(cell => cell === "1");
        this.cellColor = `#${color}`;
      }
    }
  },
  mounted() {
    window.addEventListener("mouseup", this.resetDragDrawState);
  },
  methods: {
    randomize() {
      const sizeHalf = Math.ceil(this.size / 2);
      for (let y = 0; y < this.size; y += 1) {
        for (let x = 0; x < sizeHalf; x += 1) {
          const pos = y * this.size + x;
          const mirroredPos = y * this.size + (this.size - 1 - x);
          const cell = Math.random() >= 0.5;
          Vue.set(this.cells, pos, cell);
          if (mirroredPos !== pos) {
            Vue.set(this.cells, mirroredPos, cell);
          }
        }
      }
    },
    resetDragDrawState() {
      this.lastCellClickedState = null;
    },
    cellClicked(event, pos) {
      if (event.button === 0) {
        this.lastCellClickedState = !this.cells[pos];
        this.toggle(pos);
      }
    },
    cellHovered(event, pos) {
      if (this.lastCellClickedState !== null && event.buttons === 1) {
        Vue.set(this.cells, pos, this.lastCellClickedState);
      }
    },
    toggle(pos) {
      Vue.set(this.cells, pos, !this.cells[pos]);
    },
    intoData() {
      let data = this.cells.map(cell => (cell ? "1" : "0")).join("");
      data += this.cellColor;
      return data;
    }
  }
};
</script>
<style lang="scss" scoped>
svg {
  user-select: none;
  filter: drop-shadow(0px 3px 6px rgba(0, 0, 0, 0.16));
}
</style>
