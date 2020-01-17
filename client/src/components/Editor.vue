<template>
  <div class="hello">
    <svg
      v-if="showGrid"
      class="grid"
      width="100%"
      height="100%"
      xmlns="http://www.w3.org/2000/svg"
    >
      <defs>
        <pattern
          id="smallGrid"
          :width="gridSizeX"
          :height="gridSizeY"
          patternUnits="userSpaceOnUse"
        >
          <path
            :d="`M ${gridSizeX} 0 L 0 0 0 ${gridSizeY}`"
            fill="none"
            stroke="gray"
            stroke-width="1"
          />
        </pattern>
      </defs>

      <rect
        width="100%"
        height="100%"
        stroke="#000000"
        fill="url(#smallGrid)"
      />
    </svg>
    <canvas
      class="editor-canvas"
      ref="canvas"
      tabindex="0"
      :width="size"
      :height="size"
      @mousedown="mouseDown"
      @mousemove="mouseMove"
      @keydown.ctrl.z.exact="undo"
      @keydown.ctrl.y="redo"
    />

    <div class="colors">
      <div
        class="color"
        :class="{ selected: selectedColor === index }"
        v-for="(color, index) in colors"
        :key="index"
        @click="selectColor(index)"
        :style="{ backgroundColor: color }"
      >
        <Chrome
          class="color-picker"
          disableAlpha
          v-if="index === editingColor"
          :value="colors[editingColor]"
          @input="updateColor"
        />
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { Chrome } from "vue-color";
import arrayBufferToBase64 from "../utils/arrayBufferToBase64";
import hexToRgb from "../utils/hexToRgb";

export default {
  components: {
    Chrome
  },
  props: {
    size: {
      type: Number,
      required: true
    }
  },
  data: () => ({
    showGrid: false,
    colors: [
      "#ffffff",
      "#000000",
      "#b440a3",
      "#ff91ab",
      "#79c220",
      "#f1e899",
      "#2a1a54",
      "#20798b"
    ],
    selectedColor: 1,
    editingColor: null
  }),
  computed: {
    gridSizeX() {
      const rect = this.$refs.canvas.getBoundingClientRect();

      return rect.width / this.size;
    },
    gridSizeY() {
      const rect = this.$refs.canvas.getBoundingClientRect();

      return rect.height / this.size;
    }
  },
  created() {
    this.isDrawing = false;
    this.pixels = new Uint8Array(this.size * this.size);
    this.dragPixels = new Set();
    this.undoStack = [];
    this.redoStack = [];
    this.engulfedClick = false;
  },
  mounted() {
    this.ctx = this.$refs.canvas.getContext("2d");
    this.showGrid = true;
    this.draw();

    document.addEventListener("mouseup", this.mouseUp);
    document.addEventListener("mousedown", this.docClick);
    document.addEventListener("keyup", this.pickColorShortcut);
  },
  beforeDestroy() {
    document.removeEventListener("mouseup", this.mouseUp);
    document.removeEventListener("mousedown", this.docClick);
    document.removeEventListener("keyup", this.pickColorShortcut);
  },
  watch: {
    editingColor() {
      this.engulfedClick = false;
    }
  },
  methods: {
    updateColor(color) {
      Vue.set(this.colors, this.selectedColor, color.hex);
      this.draw();
    },
    selectColor(colorIndex) {
      if (this.selectedColor === colorIndex) {
        this.editingColor = colorIndex;
        return;
      }
      this.selectedColor = colorIndex;
    },
    draw() {
      this.ctx.clearRect(
        0,
        0,
        this.$refs.canvas.width,
        this.$refs.canvas.height
      );

      this.pixels.forEach((colorIndex, index) => {
        const x = index % this.size;
        const y = Math.floor(index / this.size);

        this.ctx.fillStyle = this.colors[colorIndex];
        this.ctx.fillRect(x, y, 1, 1);
      });

      this.ctx.fillStyle = this.colors[this.selectedColor];
      this.dragPixels.forEach(index => {
        const x = index % this.size;
        const y = Math.floor(index / this.size);
        this.ctx.fillRect(x, y, 1, 1);
      });
    },
    dragComplete() {
      this.undoStack.push({});
      this.dragPixels.forEach(index => {
        this.undoStack[this.undoStack.length - 1][index] = this.pixels[index];
        this.pixels[index] = this.selectedColor;
      });
      this.dragPixels.clear();
      this.redoStack = [];
    },
    drag(x, y) {
      const index = x + this.size * y;

      if (this.pixels[index] === this.selectedColor) {
        return;
      }
      this.dragPixels.add(index);
      this.draw();
    },
    undo() {
      const undo = this.undoStack.pop();
      if (!undo) {
        return;
      }
      this.redoStack.push({});
      for (let index in undo) {
        this.redoStack[this.redoStack.length - 1][index] = this.pixels[index];
        this.pixels[index] = undo[index];
      }

      this.draw();
    },
    redo() {
      const redo = this.redoStack.pop();
      if (!redo) {
        return;
      }
      this.undoStack.push({});
      for (let index in redo) {
        this.undoStack[this.undoStack.length - 1][index] = this.pixels[index];
        this.pixels[index] = redo[index];
      }

      this.draw();
    },
    toBytes() {
      // One 16uint
      const widthBytes = 2;
      // One 16uint
      const heightBytes = 2;
      // Two 16uint
      const sizeBytes = widthBytes + heightBytes;
      // One 8uint
      const colorsBytes = 1;
      // Number of colors for each is 3 bytes
      const colorTableBytes = this.colors.length * 3;
      // Each pixel is 1 byte
      const pixelsBytes = this.size * this.size;

      const length = sizeBytes + colorsBytes + colorTableBytes + pixelsBytes;

      const bytes = new ArrayBuffer(length);
      const dataview = new DataView(bytes, 0, length);

      dataview.setUint16(0, this.size);
      dataview.setUint16(widthBytes, this.size);
      dataview.setUint8(sizeBytes, this.colors.length);
      this.colors.forEach((color, index) => {
        const { r, g, b } = hexToRgb(color);
        dataview.setUint8(sizeBytes + colorsBytes + index * 3, r);
        dataview.setUint8(sizeBytes + colorsBytes + index * 3 + 1, g);
        dataview.setUint8(sizeBytes + colorsBytes + index * 3 + 2, b);
      });

      this.pixels.forEach((colorIndex, index) => {
        dataview.setUint8(
          sizeBytes + colorsBytes + colorTableBytes + index,
          colorIndex
        );
      });

      return arrayBufferToBase64(bytes);
    },

    getMousePos(canvas, evt) {
      const rect = canvas.getBoundingClientRect();

      const widthRatio = this.size / rect.width;
      const heightRatio = this.size / rect.height;

      // The "Math.max(..., 0)"" fixes a weird bug where if the user drags the
      // mouse near the left edge of the canvas while holding click it would
      // place a pixel one row up and on the far right of the canvas.
      // This would happen when "evt.clientX - rect.left" or the vertical
      // counterpart evaluated to a negative number.
      return [
        Math.floor(Math.max(evt.clientX - rect.left, 0) * widthRatio),
        Math.floor(Math.max(evt.clientY - rect.top, 0) * heightRatio)
      ];
    },
    mouseDown(event) {
      this.isDrawing = true;
      const [x, y] = this.getMousePos(this.$refs.canvas, event);
      this.drag(x, y);
    },
    mouseMove(event) {
      if (!this.isDrawing) {
        return;
      }
      const [x, y] = this.getMousePos(this.$refs.canvas, event);
      this.drag(x, y);
    },
    mouseUp() {
      this.isDrawing = false;
      this.dragComplete();
    },
    docClick(event) {
      // This is a bit hacky but it seem to be the best method to fix that
      // when we click on a color it would set editingColor then this handler
      // would run reseting back to null
      // if (!this.engulfedClick) {
      //   this.engulfedClick = true;
      //   return;
      // }
      if (
        this.editingColor === null ||
        event.path.some(
          el => el.classList && el.classList.contains("color-picker")
        )
      ) {
        return;
      }
      this.editingColor = null;
    },
    pickColorShortcut(event) {
      if (event.keyCode < 49 || event.keyCode >= 49 + this.colors.length) {
        return;
      }
      this.selectedColor = event.keyCode - 49;
    }
  }
};
</script>
<style lang="scss" scoped>
.grid {
  position: absolute;
  pointer-events: none;
}
.grid,
.editor-canvas {
  width: 512px;
  height: 512px;
}
.editor-canvas {
  image-rendering: -moz-crisp-edges;
  image-rendering: -webkit-crisp-edges;
  image-rendering: pixelated;
  image-rendering: crisp-edges;
}
.color {
  width: 35px;
  height: 35px;
  border: solid 3px black;
  margin: 1rem;
  display: inline-block;
  position: relative;
  .color-picker {
    position: absolute;
    z-index: 1;
    left: 50%; /* position the left edge of the element at the middle of the parent */
    top: 100%;

    transform: translate(-50%, 3px);
  }
  &.selected {
    /* box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22); */
    /* .edit {
      transform: translateY(55px);
      opacity: 1;
    } */
    border-color: #f00;
  }

  /* .edit {
    opacity: 0;
    z-index: -1;
    position: absolute;
    pointer-events: auto;
    display: inline-block;
    width: 35px;
    height: 35px;
    border-radius: 100%;
    box-shadow: 0 7px 7px rgba(black, 0.15);
    /* background-color: #871f78; 
    transition: transform 0.2s, opacity 0.2s;
    /* transition: box-shadow 0.2s; 
    cursor: pointer;
    padding: 5px 5px;
    position: relative;
    .icon {
      width: 100%;
      height: 100%;
      fill: black;
      opacity: 0.5;
    }
  } */
}
</style>
