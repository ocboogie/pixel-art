<template>
  <v-card>
    <avatar
      ref="avatar"
      class="avatar"
      editable
      :avatarData="avatarData"
      :size="size"
      :color="selectedColor"
    />
    <div class="colors">
      <div
        v-for="color in palette"
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

export default {
  components: {
    Avatar
  },
  props: {
    avatarData: {
      type: String,
      default: null
    },
    palette: {
      type: Array,
      required: true
    },
    size: {
      type: Number,
      required: true
    }
  },
  data() {
    return {
      selectedColor: this.palette[
        Math.floor(Math.random() * this.palette.length)
      ]
    };
  },
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
