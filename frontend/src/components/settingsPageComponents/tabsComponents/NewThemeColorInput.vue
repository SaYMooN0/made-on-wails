<template>
    {{ label }}
    <input type= "color"/>
    <span :style="`background-color: ${color}`" class="color-div"></span>
</template>
<script>
export default {
  props: {
    label:String,
    color: String
  },
};
</script>
