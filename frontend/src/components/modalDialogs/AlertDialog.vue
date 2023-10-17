<template>
    <EmptyDialog ref="dialogRef">
      <div class="dialog-content">
        <button @click="closeDialog">Ok</button>
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
  