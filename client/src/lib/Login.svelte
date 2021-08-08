<script>
  import cn from "classnames";
  import { ButtonPrimary, Card, Input, Label } from "./base";
  import { me } from "../store";
  import axios from "../axios";

  let email;
  let password;

  async function login() {
    await axios.post("/auth/login", { email, password });
    me.set(await axios.get("/me"));
  }
</script>

<div class={cn(Card, "px-4 py-4", $$props.class)}>
  <form class="max-w-xs m-auto" on:submit|preventDefault={login}>
    <label class={cn(Label, "block text-left mb-4")}>
      <div>Email</div>
      <input class={cn(Input, "w-full")} type="email" bind:value={email} />
    </label>
    <label class={cn(Label, "block text-left mb-4")}>
      <div>Password</div>
      <input
        class={cn(Input, "w-full")}
        type="password"
        bind:value={password}
      />
    </label>
    <input type="submit" class={cn(ButtonPrimary, "w-full")} value="Login" />
  </form>
</div>
