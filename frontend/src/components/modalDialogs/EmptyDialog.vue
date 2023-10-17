<template>
  <dialog ref="dialogRef" class="dialog-container">
    <slot></slot>
  </dialog>
</template>

<script>
export default {
  methods: {
    showModal() {
      this.$refs.dialogRef.showModal();
    },
    close() {
      this.$refs.dialogRef.close();
    }
  }
}
</script>

<style>
.dialog-container {
  position: absolute;
  top: 0;
  bottom: 14%;
  width: calc(200px + 32%);
  background-color: var(--back-2);
  border-radius: calc(0.35vh + 0.35vw + 7px);
  border-width: calc(0.05vh + 0.05vw + 4px);
    border-color: var(--front-2);
  padding-top: 0;
  padding-left: 1%;
  padding-right: 1%;
}
</style>
  