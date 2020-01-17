<template>
  <canvas class="art-renderer" ref="canvas" :width="width" :height="height" />
</template>
<script>
import decodeArtData from "../utils/decodeArtData";

export default {
  props: {
    data: {
      type: String,
      required: true
    }
  },
  data: () => ({ width: 25, height: 25 }),
  watch: {
    data: {
      immediate: true,
      handler(data) {
        const info = decodeArtData(data);

        this.pixels = info.pixels;
        this.colors = info.colors;
        this.width = info.width;
        this.height = info.height;

        this.$nextTick(() => {
          this.draw();
        });
      }
    }
  },
  mounted() {
    this.ctx = this.$refs.canvas.getContext("2d");
  },
  methods: {
    draw() {
      this.ctx.clearRect(
        0,
        0,
        this.$refs.canvas.width,
        this.$refs.canvas.height
      );

      this.pixels.forEach((colorIndex, index) => {
        const x = index % this.width;
        const y = Math.floor(index / this.height);

        this.ctx.fillStyle = this.colors[colorIndex];
        this.ctx.fillRect(x, y, 1, 1);
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.art-renderer {
  image-rendering: -moz-crisp-edges;
  image-rendering: -webkit-crisp-edges;
  image-rendering: pixelated;
  image-rendering: crisp-edges;
}
</style>