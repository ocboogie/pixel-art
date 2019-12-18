<template>
  <v-card max-width="344" class="mx-auto post">
    <v-list-item>
      <v-list-item-avatar
        class="avatar elevation-2"
        color="white"
      ></v-list-item-avatar>
      <v-list-item-content>
        <v-list-item-title>Our Changing Planetarstarstarst</v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <v-img src="https://art.pixilart.com/5c3a0d7c7f68636.png"></v-img>

    <v-card-actions>
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
export default {
  props: {
    post: {
      type: Object,
      required: true
    }
  }
};
</script>
<style lang="scss" scoped>
// The .post is to give this a higher CSS priority, making the margin-right
// override vuetify
.post .avatar {
  margin-left: -35px;
  margin-right: 12px;
}
</style>
