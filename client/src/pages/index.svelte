<script>
  import PostList from "../lib/PostList.svelte";
  import { me } from "../store";
</script>

<div class="m-auto text-center mt-8 max-w-lg">
  <PostList source={$me ? "/feed" : "/posts"} />
</div>
