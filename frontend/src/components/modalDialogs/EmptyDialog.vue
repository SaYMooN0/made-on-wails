<template>
  <div class="dialog-wrapper" v-if="isOpen">
    <dialog ref="dialogRef" class="dialog-container">
      <slot></slot>
    </dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isOpen: false
    }
  },
  mounted() {
    document.addEventListener('keydown', this.handleKeydown);
  },
  beforeDestroy() {
    document.removeEventListener('keydown', this.handleKeydown);
  },
  methods: {
    showDialog() {
      this.isOpen = true;
      this.$nextTick(() => {
        this.$refs.dialogRef.showModal();
      });
    },
    closeDialog() {
      this.isOpen = false;
      this.$refs.dialogRef.close();
    },
    handleKeydown(event) {
      if (event.key === 'Escape' && this.isOpen) {
        this.close();
      }
    }
  }
}
</script>

<style>
.dialog-wrapper {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.2);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;
}

.dialog-container {
  width: calc(180px + 23% + 2vh);
  background-color: var(--back-2);
  border-radius: calc(0.35vh + 0.35vw + 7px);
  border-width: calc(0.05vh + 0.05vw + 4px);
  border-color: var(--front-2);
  padding-left: calc(3px + 1%);
  padding-right: calc(3px + 1%);
  z-index: 11;
  padding: 1%
}

button {
  width: calc(40px + 1.1vh + 3.5vw);
  height: calc(20px + 1.8vh + 0.3vw);
  background-color: var(--back-3);
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: 'Figtree';
  font-size: calc(0.3vh + 0.41vw + 11px);
  font-weight: 200;
  color: var(--front);
  border: 1px solid transparent;
  border-radius: calc(2px + 0.04vw + 0.1vh);
  cursor: pointer;
}

button:hover {
  background-color: var(--bright-2);
}

.dialog-container:focus {
  outline: none;
}
</style>
  