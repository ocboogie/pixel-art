<script>
  import PostList from "../lib/PostList.svelte";
  import { me } from "../store";
  import axios from "../axios";
  import { onMount } from "svelte";

  let posts = null;
  onMount(async () => {
    posts = (
      await axios.get($me ? "/feed" : "/posts", {
        params: {
          author: "a",
          liked: "a",
          likes: "a",
        },
      })
    ).data;
  });
</script>

<div class="m-auto text-center mt-8 max-w-lg">
  {#if posts}
    <PostList {posts} />
  {/if}
</div>
