<template>
    <button @click="backToMain" class="go-back-button">Go back</button>
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
<style>
.go-back-button {
    position: absolute;
    right: calc(8px + 1vw);
    bottom: calc(6px + 2%);
    height: calc(2vh + 22px + 0.2vw);
    width: calc(4vw + 60px + 2vh);
    background-color: var(--bright-2);
    font-family: 'Figtree';
    font-size: calc(0.55vh + 0.55vw + 10px);
    font-weight: 300;
    color: var(--front);
    border: 1px solid transparent;
    border-radius: calc(3px + 0.09vw + 0.1vh);
    cursor: pointer;
    transition: 0.08s;
}

.go-back-button:hover {
    background-color: var(--bright-3);
}
</style>