<template>
    <input type="text" class="default-input" @input="handleInput" />
</template>

<script>
export default {
  props: {
    suggestion: {
      type: String,
      default: 'none'
    }
  },
  methods: {
    handleInput() {
      if (this.suggestion === 'item') {
        console.log('item');
      } else if (this.suggestion === 'type') {
        console.log('type');
      }
    }
  }
}
</script>

<style scoped>
.default-input {
    color: cadetblue;
}
</style>
