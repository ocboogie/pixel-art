<script>
  import cn from "classnames";
  import ArtRenderer from "./ArtRenderer.svelte";
  import Avatar from "./Avatar.svelte";

  export let post;
</script>

<div class={cn("rounded bg-white shadow drop-shadow", $$props.class)}>
  <!-- TODO: Take the user to the user page of the author -->
  <div class="flex flex-row p-2 justify-start items-center">
    <!-- FIXME: The size should either be fetched on page load or be saved in
        the avatar data format -->
    <Avatar class="h-10 w-10 rounded-full drop-shadow shadow" size="5" />
    <div class="ml-2">{post.author.name}</div>
  </div>
  <hr />
  <div>
    <ArtRenderer class="w-full" data="AAMAAwP/AAAA/wAAAP8AAQIAAQIAAQI=" />
  </div>
  <hr />
  <div class="flex flex-row p-2 justify-end items-center">
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="w-12 text-gray-500 hover:bg-opacity-5 bg-opacity-0 bg-black duration-150 rounded-full p-2.5 cursor-pointer"
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
      class="w-12 text-gray-500 hover:bg-opacity-5 bg-opacity-0 bg-black duration-150 rounded-full p-2.5 cursor-pointer"
      fill="currentColor"
      viewBox="0 0 20 20"
    >
      <path
        d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z"
      />
    </svg>
  </div>
</div>
