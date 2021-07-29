<script>
  import { url } from "@roxi/routify";
  import cn from "classnames";

  export let loggedIn;
</script>

<nav class="drop-shadow shadow bg-blue-500 flex justify-between items-center">
  <div class="left p-4">
    <a href={$url("/")} class="text-white text-xl">Pixel Art</a>
  </div>
  <div class="right px-2">
    {#if loggedIn}
      <div class="profile-container">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="w-12 text-white hover:bg-opacity-10 bg-opacity-0 bg-white duration-150 rounded-full p-2.5"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
            clip-rule="evenodd"
          />
        </svg>
        <div
          class="profile-menu absolute right-0 w-44 bg-white py-2 drop-shadow-md shadow-md rounded z-30"
        >
          <a
            class="px-2 py-3 hover:bg-gray-200 duration-75 block cursor-pointer"
            href="/profile"
          >
            Profile
          </a>
          <a
            class="px-2 py-3 hover:bg-gray-200 duration-75 block cursor-pointer"
            href="/settings"
          >
            Settings
          </a>
          <a
            class="px-2 py-3 hover:bg-gray-200 duration-75 block cursor-pointer"
            href="/logout">Logout</a
          >
        </div>
      </div>
    {:else}
      <a
        class="px-3 py-2 text-white rounded hover:bg-white hover:bg-opacity-20 duration-75 block cursor-pointer"
        href="/login"
      >
        Login
      </a>
    {/if}
  </div>
</nav>

<style lang="postcss">
  .profile-container .profile-menu {
    display: none;
  }
  .profile-container:hover .profile-menu {
    display: block;
  }
</style>
