<template>
    <input type="number" class="default-input-num" @input="handleInputEvent" />
  </template>
  
  <script>
  export default {
    methods: {
      handleInputEvent(event) {
        const inputValue = parseFloat(event.target.value);
        this.$emit('input', isNaN(inputValue) ? null : inputValue);
      }
    }
  }
  </script>
  
  
<style scoped>
.default-input-num {
    height: calc(1.6vh + 8px + 0.25vw);
    width: calc(2.6vw + 14px + 0.6vh);
    background-color: var(--back-2);
    border: 1px solid transparent;
    border-radius: calc(1px + 0.04vw + 0.1vh);
    font-family: 'Figtree';
    font-size: calc(0.88vh + 0.6vw + 7px);
    font-weight: 300;
    color: var(--front);
    text-align: center;
}

.default-input-num::-webkit-inner-spin-button,
.default-input-num::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

.default-input-num:focus {
    outline: none;
    background-color: var(--back-3);
}
</style>