<template>
  <label class="color-label">
    {{ label }}
    <div :style="`background-color: ${color}`" class="color-div"></div>
  </label>
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
.color-label {
  color: black;
  font-size: calc(0.5vh + 0.3vw + 11px);
  font-weight: 600;
  font-family: 'Figtree';
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 8%;
}
.color-div{
  height: calc(0.5vh + 0.32vw + 11px);
  aspect-ratio: 1/1;
  border-radius: 10%;
}
</style>
  