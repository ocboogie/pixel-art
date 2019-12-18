<template>
  <v-card>
    <avatar ref="avatar" class="avatar" editable :color="selectedColor" />
    <div class="colors">
      <div
        v-for="color in colors"
        :key="color"
        class="color"
        :style="{ backgroundColor: color }"
        @click="changeColor(color)"
      />
    </div>
    <v-btn class="random" @click="randomize">Random</v-btn>
  </v-card>
</template>
<script>
import Avatar from "./Avatar.vue";

const colors = ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#e74c3c"];

export default {
  components: {
    Avatar
  },
  data: () => ({
    colors,
    selectedColor: colors[Math.floor(Math.random() * colors.length)]
  }),
  methods: {
    changeColor(color) {
      this.selectedColor = color;
    },
    randomize() {
      this.$refs.avatar.randomize();
    },
    intoData() {
      return this.$refs.avatar.intoData();
    }
  }
};
</script>
<style lang="scss" scoped>
.colors {
  display: flex;
  align-items: stretch;
  justify-content: space-around;
  margin-bottom: 10px;
  .color {
    cursor: pointer;
    width: 25px;
    height: 25px;
    box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.15);
    border-radius: 3px;
    &:last-child {
      margin: 0;
    }
  }
}
.random {
  width: 100%;
}
</style>
