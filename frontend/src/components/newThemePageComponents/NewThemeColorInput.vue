<template>
  <p>
    {{ label }}: 
    {{ editableColor }}
    <input 
      type="color" 
      class="color-input" 
      :value="editableColor" 
      @input="updateColor($event.target.value)"
    />
  </p>
</template>

<script>
export default {
  props: {
    label: String,
    color: String
  },
  data() {
    return {
      editableColor: this.color
    };
  },
  methods: {
    updateColor(newColor) {
      this.editableColor = newColor;
      this.$emit('color-updated', newColor);
    }
  }
};
</script>

<style scoped>
.color-input {
  border: none;
  width: 32px;
  height: 32px;
}
input[type="color"]::-webkit-color-swatch-wrapper {
  padding: 0;
}
input[type="color"]::-webkit-color-swatch {
  border: none;
}
</style>
