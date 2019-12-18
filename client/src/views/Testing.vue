<template>
  <div>
    <Post :post="post" />
  </div>
</template>
<script>
import Post from "../components/Post.vue";

export default {
  components: {
    Post
  },
  data: () => ({
    post: {
      id: "122accdf-019e-4992-8889-1b0a6b025887",
      author: {
        id: "422cb2d3-7213-4734-aab2-5f31f0645925",
        name: "test",
        createdAt: "2019-11-30T00:31:22.414657Z"
      },
      title: "test",
      data: "",
      createdAt: "2019-11-30T00:34:04.207462Z"
    }
  })
};
</script>
