<template>
  <div>
    <v-dialog fullscreen hide-overlay v-model="value" key="helper">
      <v-card>
        <v-toolbar color="primary">
          <v-btn icon @click="$emit('click')"><v-icon>mdi-close</v-icon></v-btn>
          <v-toolbar-title>帮助</v-toolbar-title>
        </v-toolbar>
        <v-spacer></v-spacer>
        <v-card-text>施工中</v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  props: {
    value: Boolean,
  },
};
</script>