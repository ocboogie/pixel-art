<template>
  <div class="home">
    <Post v-for="post in posts" :key="post.id" :post="post" class="post" />
  </div>
</template>

<script>
import { mapActions, mapState } from "vuex";
import Post from "../components/Post.vue";

export default {
  name: "home",
  components: {
    Post
  },
  computed: mapState(["posts"]),
  created() {
    this.loadPosts();
  },
  methods: mapActions(["loadPosts"])
};
</script>
<style lang="scss" scoped>
.home {
  padding-top: 2rem;
  .post {
    margin-bottom: 1rem;
  }
}
</style>