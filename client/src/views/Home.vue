<template>
  <div class="home">
    <ul>
      <li v-for="post in posts" :key="post.id">{{ post.title }}</li>
    </ul>
    <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import { mapActions, mapState } from "vuex";

export default {
  name: "home",
  computed: mapState(["posts"]),
  components: {
    HelloWorld
  },
  mounted() {
    this.loadPosts();
  },
  methods: mapActions(["loadPosts"])
};
</script>
