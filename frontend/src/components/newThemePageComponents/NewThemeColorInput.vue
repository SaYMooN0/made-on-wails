<template>
  <p>
    {{ label }}:
    {{ editableColor }}
    <input type="color" class="color-input" :value="editableColor" @input="updateColor($event.target.value)" />
    <svg class="paste-button" viewBox="0 0 512 512" @click="pasteColor">
      <path
        d="M104.6 48H64C28.7 48 0 76.7 0 112V384c0 35.3 28.7 64 64 64h96V400H64c-8.8 0-16-7.2-16-16V112c0-8.8 7.2-16 16-16H80c0 17.7 14.3 32 32 32h72.4C202 108.4 227.6 96 256 96h62c-7.1-27.6-32.2-48-62-48H215.4C211.6 20.9 188.2 0 160 0s-51.6 20.9-55.4 48zM144 56a16 16 0 1 1 32 0 16 16 0 1 1 -32 0zM448 464H256c-8.8 0-16-7.2-16-16V192c0-8.8 7.2-16 16-16l140.1 0L464 243.9V448c0 8.8-7.2 16-16 16zM256 512H448c35.3 0 64-28.7 64-64V243.9c0-12.7-5.1-24.9-14.1-33.9l-67.9-67.9c-9-9-21.2-14.1-33.9-14.1H256c-35.3 0-64 28.7-64 64V448c0 35.3 28.7 64 64 64z" />
    </svg>
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
    },
    pasteColor() {
      navigator.clipboard.readText().then((clipText) => {
        if (/^#[0-9A-Fa-f]{6}$/.test(clipText)) {
          this.updateColor(clipText);
        } else {
          this.showNotification('The color in the clipboard is not in a valid hex format. Use the hex representation of the color', true);
        }
      }).catch((err) => {
        this.showNotification('An error has occurred: '+err, true);
      });
    },
  },
  inject: ['showNotification'],
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

.paste-button {
  width: 32px;
  height: 32px;
}
</style>
