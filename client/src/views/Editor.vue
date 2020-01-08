<template>
  <div class="text-center">
    <v-text-field class="title-input" label="Title" v-model="title" />
    <Editor ref="editor" v-if="artFormat" :size="artFormat.size" />
    <v-btn :loading="posting" color="primary" x-large @click="post">Post</v-btn>
  </div>
</template>
<script>
import Editor from "../components/Editor.vue";
import axios from "../plugins/axios";

export default {
  components: {
    Editor
  },
  data: () => ({
    posting: false,
    title: "",
    artFormat: null
  }),
  async created() {
    const { data } = await axios.get("/art/format");
    this.artFormat = data;
  },
  methods: {
    async post() {
      this.posting = true;
      const data = this.$refs.editor.toBytes();
      await axios.post("/posts", {
        title: this.title,
        data
      });
      this.posting = false;
    }
  }
};
</script>
<style lang="scss" scoped>
.title-input {
  margin: auto;
  max-width: 300px;
  margin-top: 1rem;
}
</style>