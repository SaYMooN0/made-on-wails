<template>
  <EmptyDialog ref="dialogRef">
    <div class="dialog-content">
      <p class="error-text">{{ errorText }}</p>
      <div class="btn-container">
        <button @click="closeDialog" class="close-button">Close</button>
      </div>

    </div>
  </EmptyDialog>
</template>

<script>
import EmptyDialog from './EmptyDialog.vue';

export default {
  props: {
    errorText: {
      type: String,
      default: ""
    }
  },
  methods: {
    showDialog() {
      this.$refs.dialogRef.showDialog();
    },
    closeDialog() {
      this.$refs.dialogRef.closeDialog();
    },
  },
  components: { EmptyDialog }
}
</script>
<style>
.dialog-container {
  border-color: var(--warning-bright);
}

.error-text {
  white-space: pre-line;
  font-family: 'Figtree';
  color: var(--front);
  font-size: calc(0.6vh + 0.55vw + 6px);
  font-weight: 200;
  width: 100%;
  margin-bottom: 0;
  text-align: justify;
}

.btn-container {
  margin-top: 0px;
  display: flex;
  justify-content: flex-end;
  width: 100%;
}
</style>