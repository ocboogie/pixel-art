<script>
  import { url } from "@roxi/routify";
  import AvatarEditor from "../lib/AvatarEditor.svelte";
  import Avatar from "../lib/Avatar.svelte";
  import ColorWheel from "../lib/ColorWheel.svelte";
  import Post from "../lib/Post.svelte";

  const post = {
    id: "122accdf-019e-4992-8889-1b0a6b025887",
    author: {
      id: "422cb2d3-7213-4734-aab2-5f31f0645925",
      name: "test",
      createdAt: "2019-11-30T00:31:22.414657Z",
      avatar: "1101101101110111101111011#e74c3c",
    },
    title: "test",
    data: "",
    createdAt: "2019-11-30T00:34:04.207462Z",
  };

  console.log($url());
</script>

<!-- <Avatar size="3" avatarData="101010101" color="#1abc9c" editable /> -->
<div class="m-auto text-center mt-8">
  <!-- <ProfileSettings user={post.author} /> -->
  <!-- <ColorWheel
    value="0"
    colors={["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#e74c3c"]}
    radius="50"
    strokeWidth="20"
    emptySpace="90"
    gap="5"
  /> -->
  <Post {post} class="max-w-lg m-auto" />
  <!-- <AvatarEditor
    size="5"
    palette={["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#e74c3c"]}
  /> -->
</div>
