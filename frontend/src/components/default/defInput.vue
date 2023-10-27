<template>
  <input type="text" class="default-input" @input="handleInputEvent" />
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
    handleInputEvent() {
      if (this.suggestion === 'item') {
        console.log('item');
      } else if (this.suggestion === 'type') {
        console.log('type');
      }
      this.$emit('input', this.inputValue);
    }
  }
}
</script>

<style scoped>
.default-input {
  height: calc(1.6vh + 8px + 0.25vw);
  width: calc(22vw + 70px + 1vh);
  background-color: var(--back-2);
  border: 1px solid transparent;
  border-radius: calc(2px + 0.05vw + 0.1vh);
  font-family: 'Figtree';
  font-size: calc(0.88vh + 0.6vw + 7px);
  font-weight: 300;
  color: var(--front);
  padding-left: calc(1px + 0.2vw + 0.2vh);
  padding-right: calc(1px + 0.2vw + 0.2vh);
}

.default-input:focus {
  outline: none;
  background-color: var(--back-3);
}
</style>
