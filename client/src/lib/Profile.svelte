<script>
  import cn from "classnames";
  import Avatar from "../lib/Avatar.svelte";
  import { avatarSpec } from "../store";
  import { ButtonPrimary, Card } from "./base";
  import Loading from "./Loading.svelte";
  import { onMount } from "svelte";
  import axios from "../axios";
  import { me } from "../store";
  import PostList from "./PostList.svelte";

  export let userId;

  let user = null;

  onMount(async () => {
    user = (
      await axios.get(`/users/${userId}`, {
        params: {
          // Just need to be non-empty
          isFollowing: "t",
          followers: "t",
          followingCount: "t",
        },
      })
    ).data;
  });
  function followUnfollow() {
    if (user.isFollowing) {
      axios.delete(`/follows/${user.id}`);
    } else {
      axios.put(`/follows/${user.id}`);
    }

    user.isFollowing = !user.isFollowing;
    user.followers += user.isFollowing ? 1 : -1;
    user = user;
  }
</script>

{#if user}
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
        <span class="text-gray-500">Followers </span>{user.followers}
      </div>
      <div>
        <span class="text-gray-500">Following </span>{user.followingCount}
      </div>
      {#if $me}
        {#if user.id != $me.id}
          <button
            class={cn(ButtonPrimary, "w-full p-1 mt-2")}
            on:click={followUnfollow}
          >
            {#if user.isFollowing}
              Unfollow
            {:else}
              Follow
            {/if}
          </button>
        {/if}
      {/if}
    </div>
  </div>
  <PostList gallery source={`/users/${userId}/posts`} />
{:else}
  <Loading>Loading...</Loading>
{/if}
