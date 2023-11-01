<template>
    <form @submit.prevent="handleSubmit" class="form-container">
        <DefLine labelText="input:">
            <InputWithSuggestions :value="initialInput" @updateValue="this.inputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="output:">
            <InputWithSuggestions :value="initialOutput" @updateValue="this.outputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="additional:" >
            ${radiosHtml}
        </DefLine>
        <DefSave :submitText="submitText" />
    </form>
    <InvalidInputDialog ref="errDialog" :errorText="`${errDialogText}`"></InvalidInputDialog>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import DefInputNum from './../../../../default/DefInputNum.vue';
import DefSave from './../../../../default/DefSave.vue';
import InvalidInputDialog from './../../../../modalDialogs/InvalidInputDialog.vue';
export default {
    props: {
        initialInput: {
            type: String,
            default: ''
        },
        initialOutput: {
            type: String,
            default: ''
        },
        initialFurnaceType: {
            type: String,
            default: "None"
        },
        isNew: {
            type: Boolean,
            default: true
        },
        filePath: {
            type: String,
            default: ''
        },
        actionId: {
            type: Number,
            default: 0
        },
    },
    data() {
        return {
            inputValue: this.initialInput,
            outputValue: this.initialOutput,
            initialFurnaceTypeValue: this.initialFurnaceType,
            errDialogText: '',
        };
    },
    computed: {
        submitText() {
            return this.isNew ? 'Save to file' : 'Save changes';
        }
    },
    methods: {
        handleSubmit() {
        }
    },
    inject: ['newFurnaceRecipeSaved'],
    components: {
        DefLine,
        InputWithSuggestions,
        DefInputNum,
        DefSave,
        InvalidInputDialog
    }
}
</script>
  
<style scoped>
.form-container {
    padding-top: calc(2vw + 10px + 20vh);
    padding-left: calc(4vw + 20px + 2vh);
    width: calc(100% - 4vw - 20px - 2vh);
    display: block;
}
</style>
  