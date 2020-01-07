<template>
  <v-card>
    <v-toolbar color="primary" dark flat class="justify-center">
      <v-toolbar-title>Profile</v-toolbar-title>
    </v-toolbar>
    <v-card-text>
      <v-text-field label="Name" v-model="user.name" />
      <AvatarEditor
        ref="avatarEditor"
        v-if="avatarFormat"
        :avatarData="user.avatar"
        :palette="avatarFormat.palette"
        :size="avatarFormat.size"
      />
    </v-card-text>
    <v-card-actions>
      <!-- <v-btn :loading="loading" color="primary" x-large block @click="submit"> -->
      <v-btn color="primary" x-large block @click="submit">
        Save
      </v-btn>
    </v-card-actions>
  </v-card>
</template>
<script>
import AvatarEditor from "./AvatarEditor.vue";
// import { mapState } from "vuex";
import store from "../store";
import axios from "../plugins/axios";
import { mapActions } from "vuex";

export default {
  data: () => ({
    avatarFormat: null,
    user: store.state.me
  }),
  async created() {
    const { data } = await axios.get("/avatar/format");
    this.avatarFormat = data;
  },
  methods: {
    ...mapActions(["updateProfile"]),
    submit() {
      this.updateProfile({
        name: this.user.name,
        avatar: this.$refs.avatarEditor.intoData()
      });
    }
  },
  components: {
    AvatarEditor
  }
};
</script>