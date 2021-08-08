<script>
  import cn from "classnames";
  import { goto } from "@roxi/routify";
  import { random, serialize } from "../utils/avatar";
  import { ButtonPrimary, Card, Input, Label } from "./base";
  import { avatarSpec, me } from "../store";
  import axios from "../axios";

  let name;
  let email;
  let password;
  let confirmPassword;

  async function signUp() {
    const color =
      $avatarSpec.palette[
        Math.floor(Math.random() * $avatarSpec.palette.length)
      ];
    const avatar = serialize(random($avatarSpec.size), color);

    await axios.post("/auth/signUp", {
      name,
      email,
      password,
      avatar: avatar,
    });
    me.set(await axios.get("/me"));
    $goto("/");
  }
</script>

<div class={cn(Card, "px-4 py-4", $$props.class)}>
  <form class="max-w-xs m-auto" on:submit|preventDefault={signUp}>
    <label class={cn(Label, "block text-left mb-4")}>
      <div>Name</div>
      <input class={cn(Input, "w-full")} type="text" bind:value={name} />
    </label>
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
    <label class={cn(Label, "block text-left mb-4")}>
      <div>Confirm Password</div>
      <input
        class={cn(Input, "w-full")}
        type="password"
        bind:value={confirmPassword}
      />
    </label>
    <input type="submit" class={cn(ButtonPrimary, "w-full")} value="Sign Up" />
  </form>
</div>
