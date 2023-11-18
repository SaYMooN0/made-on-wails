<template>
    <div :style="{ color: color }">{{ label }}: {{ color }}</div>
  </template>
  
  <script>
  export default {
    props: {
      label: String,
      color: String
    }
  };
  </script>
  
  <style scoped>
  </style>
  