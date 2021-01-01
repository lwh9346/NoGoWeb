<template>
  <div>
    <stone
      v-for="id in ids"
      :key="id"
      :id="id"
      :color="board[id].color"
      @click="click"
      :disabled="disabled"
    ></stone>
  </div>
</template>

<script>
import Stone from "./Stone.vue";
export default {
  components: { Stone },
  data: function () {
    return {
      ids: [],
    };
  },
  props: {
    board: Array,
    lineNum: Number,
    disabled: Boolean
  },
  created: function () {
    for (let i = 0; i < 9; i++) {
      this.ids.push(this.lineNum * 9 + i);
    }
  },
  methods:{
    click:function(event){
      this.$emit("click",event)
    }
  }
};
</script>

<style>
</style>