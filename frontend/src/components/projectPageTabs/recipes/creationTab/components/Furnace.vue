<template>
    <form @submit.prevent="handleSubmit" class="form-container">
        <DefLine labelText="input:">
            <InputWithSuggestions :value="initialInput" @updateValue="this.inputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="output:">
            <InputWithSuggestions :value="initialOutput" @updateValue="this.outputValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine labelText="additional:">
            <div class="radio-btns-container">
                <DefRadio name="additional" value="furnaceOnly" spanText="None" v-model="furnaceTypeValue" />
                <DefRadio name="additional" value="furnaceAndSmoker" spanText="Smoker" v-model="furnaceTypeValue" />
                <DefRadio name="additional" value="furnaceAndBlast" spanText="Blast Furnace" v-model="furnaceTypeValue" />
            </div>
        </DefLine>
        <DefSave :submitText="submitText" />
    </form>
    <InvalidInputDialog ref="errDialog" :errorText="`${errDialogText}`"></InvalidInputDialog>
</template>
  
<script>
import DefLine from './../../../../default/DefLine.vue';
import InputWithSuggestions from './../../../../default/InputWithSuggestions.vue';
import DefRadio from './../../../../default/DefRadio.vue';
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
            furnaceTypeValue: this.initialFurnaceType,
            errDialogText: '',
        };
    },
    computed: {
        isFurnaceOnlyChecked() { return this.initialFurnaceType === "None"; },
        isFurnaceAndSmokerChecked() { return this.initialFurnaceType === "Smoker"; },
        isFurnaceAndBlastChecked() { return this.initialFurnaceType === "Blast Furnace"; },
        submitText() { return this.isNew ? 'Save to file' : 'Save changes'; }
    },
    methods: {
        handleSubmit() {
            console.log("Input Value:", this.inputValue);
            console.log("Output Value:", this.outputValue);
            console.log("Furnace Type:", this.furnaceTypeValue);
        }
    },
    inject: ['newFurnaceRecipeSaved'],
    components: {
        DefLine,
        InputWithSuggestions,
        DefRadio,
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

.radio-btns-container {
    width: 60%;
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    justify-content: right;
    gap: 10px;
}

@media (max-width: 850px) {
    .radio-btns-container {
        grid-template-columns: 1fr;
    }
}
</style>
  