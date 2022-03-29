<script>
  import cn from "classnames";
  import ArtRenderer from "./ArtRenderer.svelte";
  import Avatar from "./Avatar.svelte";
  import { avatarSpec } from "../store";
  import { url } from "@roxi/routify";

  export let post;
  export let frameless = false;
</script>

<div
  class={cn(
    "rounded overflow-hidden bg-white shadow drop-shadow relative",
    { frameless },
    $$props.class
  )}
>
  {#if !frameless}
    <!-- TODO: Take the user to the user page of the author -->
    <a
      class="flex flex-row p-2 justify-start items-center"
      href={$url("/user/:id", { id: post.authorId })}
    >
      <Avatar
        class="h-10 w-10 rounded drop-shadow shadow"
        size={$avatarSpec.size}
        avatarData={post.author.avatar}
      />
      <div class="ml-2">{post.author.name}</div>
    </a>
    <hr />
  {/if}
  <div>
    <ArtRenderer class="w-full" data={post.art} />
  </div>
  {#if frameless}
    <div
      class="dimmer absolute left-0 right-0 top-0 bottom-0 pointer-events-none"
    />
  {:else}
    <hr />
  {/if}
  <div
    class={cn(
      "flex flex-row p-2 justify-end items-center",
      frameless && "top-0 right-0 absolute"
    )}
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="icon"
      viewBox="0 0 20 20"
      fill="currentColor"
    >
      <path
        fill-rule="evenodd"
        d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
        clip-rule="evenodd"
      />
    </svg>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="icon"
      fill="currentColor"
      viewBox="0 0 20 20"
    >
      <path
        d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z"
      />
    </svg>
  </div>
</div>

<style lang="postcss">
  .icon {
    @apply w-12 duration-150 rounded-full p-2.5 cursor-pointer;
  }

  :not(.frameless) .icon {
    @apply text-gray-500 bg-opacity-0 bg-black hover:bg-opacity-5;
  }

  .frameless .icon {
    @apply text-opacity-0 text-gray-100 bg-white bg-opacity-0 hover:bg-opacity-20;
  }

  .frameless:hover .icon {
    @apply text-opacity-100;
  }

  .frameless:hover .dimmer {
    @apply bg-black bg-opacity-20;
  }
</style>
