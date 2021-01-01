<template>
  <div style="margin-top: 50px;text-align:center">
    <board-line
      v-for="ln in lns"
      :key="ln"
      :lineNum="ln"
      :board="board"
      @click="clickStone"
      :disabled="disabled"
    ></board-line>
  </div>
</template>

<script>
import BoardLine from "./BoardLine.vue";

export default {
  components: { BoardLine },
  data: function () {
    return {
      board: [],
      lns: [0, 1, 2, 3, 4, 5, 6, 7, 8],
      disabled:false,
    };
  },
  created: function () {
    for (let i = 0; i < 81; i++) {
      this.board.push({color: "red" });
    }
  },
  methods: {
    clickStone: function (event) {
      this.board.splice(event,1,{color: "blue"})
    },
  },
};
</script>

<style>
</style>