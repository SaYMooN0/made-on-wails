<template>
    <EmptyDialog ref="dialogRef">
      <div class="dialog-content">
        <slot></slot>
        <button @click="closeDialog">Close</button>
      </div>
    </EmptyDialog>
  </template>
  
  <script>
  import EmptyDialog from './EmptyDialog.vue';
  
  export default {
    methods: {
      showDialog() {
        this.$refs.dialogRef.showModal();
      },
      closeDialog() {
        this.$refs.dialogRef.close();
      }
    },
    components: { EmptyDialog }
  }
  </script>
<style scoped>
.dialog-container {
    border-color: var(--warning-bright);
}

.warning-dialog-p {
    font-family: 'Figtree';
    color: var(--front);
    font-size: calc(0.6vh + 0.55vw + 6px);
    font-weight: 200;
    width: 100%;
    text-align: justify;
}

.warning-dialog:focus {
    outline: none;
}

.warning-dialog-btns-container {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr calc(36px + 1vh + 3.9vw) calc(36px + 1vh + 3.9vw);
    column-gap: calc(1vw + 5px);
}

.warning-dialog-btns-container button {
    width: 100%;
    height: calc(18px + 1.8vh + 0.3vw);
    background-color: var(--back-3);
    display: flex;
    justify-content: center;
    align-items: center;
    font-family: 'Figtree';
    font-size: calc(0.33vh + 0.48vw + 8px);
    font-weight: 200;
    color: var(--front);
    border: 1px solid transparent;
    border-radius: calc(1px + 0.04vw + 0.1vh)
}

.warning-dialog-btns-container button:hover {
    background-color: var(--bright-2);
    cursor: pointer;
}
</style>