<script>
  import cn from "classnames";
  import Avatar from "../lib/Avatar.svelte";
  import { avatarSpec } from "../store";
  import { ButtonPrimary, Card } from "./base";
  import Loading from "./Loading.svelte";
  import { onMount } from "svelte";
  import axios from "../axios";
  import PostList from "./PostList.svelte";

  export let user;

  let posts = null;

  onMount(async () => {
    posts = (await axios.get(`/users/${user.id}/posts`)).data;
  });
</script>

<div class={cn(Card, "w-full p-6 space-x-12 m-auto flex mb-6")}>
  <Avatar
    class="w-32 rounded shadow drop-shadow"
    size={$avatarSpec.size}
    avatarData={user.avatar}
  />
  <div class="text-left flex-grow">
    <h1 class="text-3xl mb-1">
      {user.name}
    </h1>
    <div>
      <span class="text-gray-500">Followers </span>123
    </div>
    <div>
      <span class="text-gray-500">Following </span>123
    </div>
    <button class={cn(ButtonPrimary, "w-full p-1 mt-2")}>Follow</button>
  </div>
</div>

{#if posts === null}
  <Loading>Loading Posts...</Loading>
{:else if posts.length === 0}
  <!-- FIXME: Don't use Loading -->
  <Loading>They don't have any posts :(</Loading>
{:else}
  <PostList gallery {posts} />
{/if}
