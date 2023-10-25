<template>
    <form @submit.prevent="handleSubmit" class="stonecutter-form">
      <p class="input-line">
        <label class="default-input-label"> input: </label>
        <input class="default-input" type="text" v-model="inputValue">
      </p>
    </form>
  </template>
  
  <script>
  export default {
    props: ['formArguments', 'actionId', 'path'],
    computed: {
      inputValue() {
        return this.formArguments && this.formArguments.input ? this.formArguments.input : '';
      },
    },
    methods: {
      handleSubmit(event) {
      }
    }
  }
  </script>
  