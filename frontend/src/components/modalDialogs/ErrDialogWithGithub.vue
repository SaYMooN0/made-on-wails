<template>
    <EmptyDialog ref="dialogRef">
        <div class="dialog-content">
            <p class="error-text">{{ errorText }}</p>
            <div class="btn-container">
                <button @click="openGithub">Github</button>
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
            this.$refs.dialogRef.showModal();
        },
        closeDialog() {
            this.$refs.dialogRef.close();
        },
        openGithub() { window.open("https://github.com/SaYMooN0/Made/issues", '_blank'); },
    },
    components: { EmptyDialog }
}
</script>
<style>
.dialog-container {
    border-color: var(--warning-bright);
}

.error-text {
    font-family: 'Figtree';
    color: var(--front);
    font-size: calc(0.6vh + 0.55vw + 6px);
    font-weight: 200;
    width: 100%;
    margin-bottom: 0;
    text-align: justify;
}

.btn-container {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr auto auto;
    column-gap: calc(0.5vw + 7px);
    justify-items: end; 
}


</style>