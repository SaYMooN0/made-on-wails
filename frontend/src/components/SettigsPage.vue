<template>
    <h1>Settings</h1>
    <h1>Settings</h1>
    <h1>Settings</h1>
    <h1>Settings</h1>
    <h1>Settings</h1>
    <button @click="backToMain">Назад</button>
</template>
  
<script>
export default {
  methods: {
    backToMain() {
      this.$emit('goBack'); // когда пользователь нажимает кнопку "Назад", это событие будет передаваться в App.vue
    }
  }
}
</script>