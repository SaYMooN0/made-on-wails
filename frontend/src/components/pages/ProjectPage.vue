<template>
    <h1>Project page</h1>
    <h1>Project page</h1>
    <h1>Project page</h1>
    <h1>Project page</h1>
    <h1>Project page</h1>
    <h1>Project page</h1>
    <h1>Project page</h1>

    <button @click="backToMain">Назад</button>
</template>
  
<script>
export default {
  methods: {
    backToMain() {
      this.$emit('goBack');
    }
  }
}
</script>