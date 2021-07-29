<script>
  import cn from "classnames";
  import { ButtonPrimary, Card, Input, Label } from "./base";
</script>

<div class={cn(Card, "px-4 py-4", $$props.class)}>
  <form class="max-w-xs m-auto">
    <label class={cn(Label, "block text-left mb-4")}>
      <div>Email</div>
      <input class={cn(Input, "w-full")} type="email" />
    </label>
    <label class={cn(Label, "block text-left mb-4")}>
      <div>Password</div>
      <input class={cn(Input, "w-full")} type="password" />
    </label>
    <input type="submit" class={cn(ButtonPrimary, "w-full")} value="Login" />
  </form>
</div>
