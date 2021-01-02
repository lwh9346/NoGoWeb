<template>
  <v-btn fab outlined elevation="1" small @click="$emit('click',id)" :key="id" :color="color" :disabled="disabled">{{id}}</v-btn>
</template>

<script>
export default {
  props: ["id", "color","disabled"],
};
</script>

<style>
</style>