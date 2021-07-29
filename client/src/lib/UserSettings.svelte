<script>
  import cn from "classnames";
  import { Label, Card, Input, CardTitle } from "./base";
</script>

<div class={cn(Card, "overflow-hidden", $$props.class)}>
  <div class={CardTitle}>User Settings</div>
  <form class="p-4 text-left">
    <label class={cn(Label, "flex mb-3 items-center")}>
      <span class="w-1/4">Name</span>
      <input class={cn(Input, "ml-3 w-3/4")} type="text" />
    </label>
    <label class={cn(Label, "flex mb-3 items-center")}>
      <span class="w-1/4">Email</span>
      <input class={cn(Input, "ml-3 w-3/4")} type="email" />
    </label>
  </form>
</div>
