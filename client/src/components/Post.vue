<template>
  <v-card max-width="344" class="mx-auto post">
    <v-list-item class="header">
      <!-- TODO: Take the user to the user page of the author -->
      <v-list-item-avatar class="avatar elevation-2" color="white">
        <!-- FIXME: The size should either be fetched on page load or be saved in
        the avatar data format -->
        <Avatar :size="5" :avatarData="post.author.avatar" />
      </v-list-item-avatar>
      <v-list-item-content>
        <v-list-item-title>{{ post.title }}</v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <ArtRenderer class="art" :data="post.data" />

    <v-card-actions class="actions">
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-share-variant</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>
<script>
import Avatar from "./Avatar.vue";
import ArtRenderer from "./ArtRenderer.vue";

export default {
  components: {
    Avatar,
    ArtRenderer
  },
  props: {
    post: {
      type: Object,
      required: true
    }
  }
};
</script>
<style lang="scss" scoped>
.post {
  .header {
    border-bottom: thin solid rgba(0, 0, 0, 0.12);
    .avatar {
      margin-left: -10px;
      margin-right: 12px;
    }
  }
  .art {
    width: 100%;
    height: 100%;
  }
  .actions {
    border-top: thin solid rgba(0, 0, 0, 0.12);
  }
}
</style>
