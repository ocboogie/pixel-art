<script>
  import cn from "classnames";
  import IntersectionObserver from "svelte-intersection-observer";
  import axios from "../axios";
  import Post from "./Post.svelte";

  export let source;
  export let gallery = false;
  export let limit = 12;

  let posts = null;
  let retrievedAll = false;

  let loaderElement;

  async function loadPosts() {
    const params = gallery
      ? { liked: "a", limit: limit }
      : {
          liked: "a",
          author: "a",
          likes: "a",
          limit: limit,
        };

    if (posts) {
      params.after = posts[posts.length - 1].createdAt;
    } else {
      posts = [];
    }

    const postsRetrieved = (await axios.get(source, { params })).data;

    if (postsRetrieved.length === 0 || postsRetrieved.length < limit) {
      retrievedAll = true;
    }

    posts = posts.concat(postsRetrieved);
  }
</script>

<div
  class={cn(
    gallery && "grid grid-cols-3 gap-3",
    gallery || "flex flex-col space-y-6",
    $$props.class
  )}
>
  {#if posts}
    {#each posts as post}
      <Post {post} frameless={gallery} />
    {/each}
  {/if}
  {#if !retrievedAll}
    <IntersectionObserver on:intersect={loadPosts} element={loaderElement}>
      <div bind:this={loaderElement} class="row-auto col-span-3">
        Loading...
      </div>
    </IntersectionObserver>
  {/if}
</div>
