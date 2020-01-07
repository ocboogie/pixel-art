<template>
  <div class="home">
    <ul>
      <li v-for="post in posts" :key="post.id">{{ post.title }}</li>
    </ul>
  </div>
</template>

<script>
import { mapActions, mapState } from "vuex";

export default {
  name: "home",
  computed: mapState(["posts"]),
  created() {
    this.loadPosts();
  },
  methods: mapActions(["loadPosts"])
};
</script>
