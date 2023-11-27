<template>
  <p>
    {{ label }}
    {{ color }}
    <input type= "color" class="color-input" :value="color"/>
  </p>
 
</template>
<script>
export default {
  props: {
    label:String,
    color: String
  },
};
</script>
<style scoped>
.color-input {
	border: none;
	width: 321px;
	height: 32px;
}
input[type="color"]::-webkit-color-swatch-wrapper {
	padding: 0;
}
input[type="color"]::-webkit-color-swatch {
	border: none;
}
</style>