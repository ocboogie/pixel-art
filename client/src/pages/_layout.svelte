<script>
  import Navbar from "../lib/Navbar.svelte";
</script>

<main class="container m-auto text-center p-8">
  <slot />
</main>
<!-- The reason the navbar is below the main container is so the dropdown  -->
<!-- in the navbar won't be covered by elements in the main container -->
<Navbar loggedIn={false} />

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  @tailwind variants;

  #app {
    flex-direction: column-reverse;
    display: flex;
  }
</style>
